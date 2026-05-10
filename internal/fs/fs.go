package fs

import "embed"

//go:embed static
var StaticFS embed.FS

const (
	IconNameAction = "action.svg"
	IconNameAdd = "add.svg"
	IconNameBookmark = "bookmark.svg"
	IconNameBranch = "branch.svg"
	IconNameCheck = "check.svg"
	IconNameDanger = "danger.svg"
	IconNameDB = "db.svg"
	IconNameExplore = "explore.svg"
	IconNameExport = "export.svg"
	IconNameFormat = "format.svg"
	IconNameGroup = "group.svg"
	IconNameHistory = "history.svg"
	IconNameKey = "key.svg"
	IconNameLink = "link.svg"
	IconNameRun = "run.svg"
	IconNameSave = "save.svg"
	IconNameSearch = "search.svg"
	IconNameSettings = "settings.svg"
	IconNameTable = "table.svg"
	IconNameThred = "thred.svg"
	IconNameVisible = "visible.svg"
)
