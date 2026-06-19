package themes

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/thirdmartini/gogui/pkg/ux/canvas/color"
)

var ThemePath = path.Join("assets", "dark")

func SetTheme(path string) {
	ThemePath = path
}

func LoadColors(palette color.Palette) error {
	data, err := ioutil.ReadFile(path.Join(ThemePath, "colors.json"))
	if err == nil {
		tc := new(ThemeColors)
		err = json.Unmarshal(data, tc)

		ColorBackground = tc.toColor(palette, "background")
		ColorForeground = tc.toColor(palette, "foreground")
		ColorTextPrimary = tc.toColor(palette, "text:primary")
		ColorTextMuted = tc.toColor(palette, "text:muted")
		ColorBorder = tc.toColor(palette, "border")

		// append in all graph colors
		ColorGraphAxis = append(ColorGraphAxis, tc.toColor(palette, "graph:axis:1"))
		ColorGraphAxis = append(ColorGraphAxis, tc.toColor(palette, "graph:axis:2"))
		ColorGraphAxis = append(ColorGraphAxis, tc.toColor(palette, "graph:axis:3"))
		ColorGraphAxis = append(ColorGraphAxis, tc.toColor(palette, "graph:axis:4"))
		ColorGraphAxis = append(ColorGraphAxis, tc.toColor(palette, "graph:axis:5"))
		ColorGraphTicks = tc.toColor(palette, "graph:ticks")
		ColorMenuBackground = tc.toColor(palette, "menu:background")
		return nil
	}
	return err
	/*
		if theme == "" {
			ColorBackground = palette.NewRGB8(255, 255, 255)
			ColorForeground = palette.NewRGB8(0, 0, 0)
			ColorTextPrimary = palette.NewRGB8(0, 0, 0)
			ColorTextMuted = palette.NewRGB8(100, 100, 100)
			ColorBorder = palette.NewRGB8(222, 222, 222)

			// append in all graph colors
			ColorGraphAxis = append(ColorGraphAxis, palette.NewRGB8(66, 120, 245))
			ColorGraphTicks = palette.NewRGB8(128, 128, 128)

			ColorMenuBackground = palette.NewRGB8(255, 250, 199)
		} else {
			ColorBackground = palette.NewRGB8(0, 0, 0)
			ColorForeground = palette.NewRGB8(255, 255, 255)
			ColorTextPrimary = palette.NewRGB8(255, 255, 255)
			ColorTextMuted = palette.NewRGB8(200, 200, 200)
			ColorBorder = palette.NewRGB8(222, 222, 222)

			// append in all graph colors
			ColorGraphAxis = append(ColorGraphAxis, palette.NewRGB8(66, 120, 245))
			ColorGraphTicks = palette.NewRGB8(128, 128, 128)

			ColorMenuBackground = palette.NewRGB8(32, 32, 0)
		}*/
}
