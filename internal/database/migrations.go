package database

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/Fremenkiel/gophant/v2/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func ApplyMigrations() {
	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Could not recover runtime information")
		return
	}

	fullFuncName := runtime.FuncForPC(pc).Name()
	pkgPath := filepath.Dir(fullFuncName) 
	
	db := CurrentDB()

	am := AutoMigrator{
		db: db,
		pkgPath: pkgPath,
	}

	am.migrate(
		models.Connection{},
		models.Database{},
		models.Group{},
		)
}

type AutoMigrator struct {
	db		*sql.DB
	pkgPath		string
}

func (m *AutoMigrator) migrate(s ...any) {
	var q []string
	for _, obj := range s {
		v := reflect.ValueOf(obj)
		vkind := v.Kind()
		if vkind == reflect.Pointer {
			v = v.Elem()
		}

			t := reflect.TypeOf(obj)
		n := fmt.Sprintf("%ss", toSnake(t.Name()))

		var r sql.NullString
		err := m.db.QueryRow("SELECT sql FROM sqlite_master WHERE type='table' AND name = ?;", n).Scan(&r)
		if err == nil && r.Valid && r.String != "" {
			q = append(q, m.alterTable(v, t, n, r.String, vkind))
		} else {
			q = append(q, m.createTable(v, t, n, vkind))
		}
	}
	qs := fmt.Sprintf(`PRAGMA foreign_keys=off;
BEGIN TRANSACTION;
		%s
COMMIT;
PRAGMA foreign_keys=on;`, strings.Join(q, "\n"))
	_, err := m.db.Exec(qs)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *AutoMigrator) createTable(v reflect.Value, t reflect.Type, name string, vkind reflect.Kind) string {
	cols, fkeys := m.generateColumns(v, t, vkind)

	for _, obj := range fkeys {
		cols = append(cols, obj)
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", name, strings.Join(cols, ", "))
}

func (m *AutoMigrator) alterTable(v reflect.Value, t reflect.Type, name, sql string, vkind reflect.Kind) string {
	spre := fmt.Sprintf("CREATE TABLE %s ", name)
	sprevq := fmt.Sprintf("CREATE TABLE \"%s\" ", name)
	if !strings.Contains(sql, spre) && !strings.Contains(sql, sprevq) {
		log.Fatal("Not a valid sql string")
	}
	sql = strings.Replace(sql, spre, "", 1)
	sql = strings.Replace(sql, sprevq, "", 1)
	cols, fkeys := m.generateColumns(v, t, vkind)
	
	for _, obj := range fkeys {
		cols = append(cols, obj)
	}

	cs := fmt.Sprintf("(%s)", strings.Join(cols, ", "))

	if cs == sql {
		return ""
	}

	var cl []string
	for obj := range strings.SplitSeq(strings.Replace(sql, "(", "", 1), ", ") {
		if sobj := strings.Split(obj, " "); len(sobj) != 0 && strings.Contains(cs, sobj[0]) {
			cl = append(cl, sobj[0])
		}
	}

	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS temp_%s %s;
INSERT INTO temp_%s (%s)
SELECT %s
FROM %s;
DROP TABLE %s;
ALTER TABLE temp_%s RENAME TO %s;`, name, cs, name, strings.Join(cl, ", "), strings.Join(cl, ", "), name, name, name, name)

	return q
}

func (m *AutoMigrator) generateColumns(v reflect.Value, t reflect.Type, vkind reflect.Kind) ([]string, []string) {
	var cols []string
	var fkeys	[]string
	for i := range v.NumField() {
		f := t.Field(i)
		ft := f.Type
		fkind := ft.Kind()
		if fkind == reflect.Pointer {
			ft = ft.Elem()
		}
		n := toSnake(f.Name)
		if ft.Kind() == reflect.Struct && strings.Contains(ft.PkgPath(), m.pkgPath) {
			sub := reflect.New(ft).Elem()
			subf := sub.FieldByName("ID")
			t, err := m.dbType(subf.Type())
			if err != nil {
				log.Fatal(err)
			}

			kid := fmt.Sprintf("%s_id", n)
			ks := []string{kid, t}
			if fkind != reflect.Pointer {
				ks = append(ks, "NOT NULL")
			}

			tn := fmt.Sprintf("%ss", toSnake(sub.Type().Name()))
			fks := fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", kid, tn, kid)
			q := fmt.Sprintf("%s, %s", strings.Join(ks, " "), fks)
			fkeys = append(fkeys, q)
			continue
		}

		t, err := m.dbType(f.Type)
		if err != nil {
			log.Fatal(err)
		}
		cs := []string{n, t}

		if n == "id" {
			cs = append(cs, "PRIMARY KEY")
		}

		if tag := f.Tag.Get("SQLITE"); tag != "" {
			ts := strings.SplitSeq(tag, ";")
			for obj := range ts {
				cs = append(cs, obj)
			}
		}

		if vkind != reflect.Pointer {
			cs = append(cs, "NOT NULL")
		}

		cols = append(cols, strings.Join(cs, " "))
	}
	return cols, fkeys
}

func toSnake(camel string) string {
    var b strings.Builder
    diff := 'a' - 'A'
    l := len(camel)
    for i, v := range camel {
        if v >= 'a' {
            b.WriteRune(v)
            continue
        }
        if (i != 0 || i == l-1) && (
            (i > 0 && rune(camel[i-1]) >= 'a') ||
                (i < l-1 && rune(camel[i+1]) >= 'a')) {
            b.WriteRune('_')
        }
        b.WriteRune(v + diff)
    }
    return b.String()
}

func (m *AutoMigrator) dbType(t reflect.Type) (string, error) {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Uint:
		return "INTEGER", nil
	case reflect.Uint8:
		return "INTEGER", nil
	case reflect.Uint16:
		return "INTEGER", nil
	case reflect.Uint32:
		return "INTEGER", nil
	case reflect.Uint64:
		return "INTEGER", nil
	case reflect.Int:
		return "INTEGER", nil
	case reflect.Int16:
		return "INTEGER", nil
	case reflect.Int32:
		return "INTEGER", nil
	case reflect.Int64:
		return "INTEGER", nil
	case reflect.Bool:
		return "BOOLEAN", nil
	case reflect.Float32:
		return "REAL", nil
	case reflect.Float64:
		return "REAL", nil
	case reflect.String:
		return "TEXT", nil
	case reflect.Complex64:
		return "BLOB", nil
	case reflect.Complex128:
		return "BLOB", nil
	}
	
	n := t.Name()
	switch n {
	case "UUID":
		return "TEXT", nil
	case "Time":
		return "DATE", nil
	}

	return "", fmt.Errorf("Type %s is not implemented", n)
}
