package th

import "image/color"

var Palette = struct {
        Background   color.RGBA
        Foreground   color.RGBA
        Border       color.RGBA
        Accent       color.RGBA
        Danger       color.RGBA
        Success      color.RGBA
				Indicator		color.RGBA
				Disabled		color.RGBA
				SecondaryText	color.RGBA
  }{
        Background: color.RGBA{10,  10,  10,  255},
        Foreground: color.RGBA{240, 240, 240, 255},
        Border:     color.RGBA{40,  40,  45,  255},
        Accent:     color.RGBA{99,  102, 241, 255},
        Danger:     color.RGBA{220, 38,  38,  255},
        Success:    color.RGBA{34,  197, 94,  255},
				Indicator: 	color.RGBA{97, 175, 229, 255},
				Disabled: 	color.RGBA{74, 74, 84, 255},
				SecondaryText: color.RGBA{66, 66, 75, 255},
  }
