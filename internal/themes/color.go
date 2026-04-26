package themes

import "image/color"

var Palette = struct {
        Background   color.RGBA
        Foreground   color.RGBA
        Border       color.RGBA
        Accent       color.RGBA
        Danger       color.RGBA
        Success      color.RGBA
  }{
        Background: color.RGBA{10,  10,  10,  255},
        Foreground: color.RGBA{240, 240, 240, 255},
        Border:     color.RGBA{40,  40,  45,  255},
        Accent:     color.RGBA{99,  102, 241, 255},
        Danger:     color.RGBA{220, 38,  38,  255},
        Success:    color.RGBA{34,  197, 94,  255},
  }
