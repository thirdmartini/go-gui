package themes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/thirdmartini/gogui/pkg/ux"
	"github.com/thirdmartini/gogui/pkg/ux/canvas/fonts"
)

const (
	FontLarge      = "default:large"
	FontHeader     = "default:header"
	FontDialog     = "default:dialog"
	FontGraphLabel = "default:graph-label"
)

var defaultFonts = make(map[string]struct {
	Font string
	Size float64
})

func LoadFonts() error {
	data, err := ioutil.ReadFile(path.Join(ThemePath, "fonts.json"))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &defaultFonts); err != nil {
		return err
	}

	for k, v := range defaultFonts {
		font := path.Join(ThemePath, v.Font)
		if err := ux.LoadFont(k, font, v.Size); err != nil {
			return err
		}
	}

	return nil
}

func Font(name string) *fonts.Font {
	if font := ux.Font(name); font != nil {
		return font
	}
	panic(fmt.Sprintf("trying to load font name(%s) that was not preloaded", name))
}
