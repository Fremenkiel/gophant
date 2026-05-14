package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func MapDefaultKeyBindings(w fyne.Window) {
	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		w.Close()
	})
} 

func MapMainKeyBindings(w fyne.Window) {
	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		w.Close()
	})
} 
