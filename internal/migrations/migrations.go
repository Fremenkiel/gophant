package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/Fremenkiel/gophant/v2/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func ApplyMigrations() {
	db, err := sql.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	am := AutoMigrator{db: db}

	test := models.Connection{}
	am.migrate(test)

}

type AutoMigrator struct {
	db		*sql.DB
}

func (m *AutoMigrator) migrate(s any) {
	v := reflect.ValueOf(s)
	vkind := v.Kind()
	if vkind == reflect.Pointer {
		v = v.Elem()
	}

	t := reflect.TypeOf(s)
	n := fmt.Sprintf("%ss", toSnake(t.Name()))

	var r sql.NullString
	err := m.db.QueryRow("SELECT sql FROM sqlite_master WHERE type='table' AND name = ?;", n).Scan(&r)
	if err == nil && r.Valid && r.String != "" {
		m.alterTable(v, t, n, r.String, vkind)
	} else {
		m.createTable(v, t, n, vkind)
	}
}

func (m *AutoMigrator) createTable(v reflect.Value, t reflect.Type, name string, vkind reflect.Kind) {
	cols := generateColumns(v.NumField(), t, vkind)

	_, err := m.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", name, strings.Join(cols, ", ")))
	if err != nil {
		log.Fatal(err)
	}
}

func (m *AutoMigrator) alterTable(v reflect.Value, t reflect.Type, name, sql string, vkind reflect.Kind) {
	spre := fmt.Sprintf("CREATE TABLE %s ", name)
	sprevq := fmt.Sprintf("CREATE TABLE \"%s\" ", name)
	if !strings.Contains(sql, spre) && !strings.Contains(sql, sprevq) {
		log.Fatal("Not a valid sql string")
	}
	sql = strings.Replace(sql, spre, "", 1)
	sql = strings.Replace(sql, sprevq, "", 1)
	cols := generateColumns(v.NumField(), t, vkind)
	cs := fmt.Sprintf("(%s)", strings.Join(cols, ", "))

	if cs == sql {
		return
	}

	var cl []string
	for obj := range strings.SplitSeq(strings.Replace(sql, "(", "", 1), ", ") {
		if sobj := strings.Split(obj, " "); len(sobj) != 0 {
			cl = append(cl, sobj[0])
		}
	}

	q := fmt.Sprintf(`PRAGMA foreign_keys=off;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS temp_%s %s;
INSERT INTO temp_%s (%s)
SELECT %s
FROM %s;
DROP TABLE %s;
ALTER TABLE temp_%s RENAME TO %s; 
COMMIT;
PRAGMA foreign_keys=on;`, name, cs, name, strings.Join(cl, ", "), strings.Join(cl, ", "), name, name, name, name)

	_, err := m.db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
}

func generateColumns(length int, t reflect.Type, vkind reflect.Kind) []string {
	var cols []string
	for i := range length {
		ft := t.Field(i)
		n := toSnake(ft.Name)
		t, err := dbType(ft.Type)
		if err != nil {
			log.Fatal(err)
		}
		cs := []string{n, t}

		if n == "id" {
			cs = append(cs, "PRIMARY KEY")
		}

		if tag := ft.Tag.Get("SQLITE"); tag != "" {
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
	return cols
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

func dbType(t reflect.Type) (string, error) {
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
