package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type KeyShortcutUtils struct {}

func NewKeyShortcutUtils() *KeyShortcutUtils {
	return &KeyShortcutUtils{}
}

func (s *KeyShortcutUtils) MapDefaultKeyBindings(w fyne.Window) {
	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		w.Hide()
	})
} 
