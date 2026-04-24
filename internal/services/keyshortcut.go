package services

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type KeyShortcutService struct {}

func NewKeyShortcutService() *KeyShortcutService {
	return &KeyShortcutService{}
}

func (s *KeyShortcutService) MapDefaultKeyBindings(w fyne.Window) {
	closeKey := desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierShortcutDefault }
	w.Canvas().AddShortcut(&closeKey, func(shortcut fyne.Shortcut) {
		w.Hide()
	})
} 
