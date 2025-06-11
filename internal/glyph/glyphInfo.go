package glyph

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Galdoba/fsmp/internal/paths"
)

const (
	Cyrillic    = "Cyrillic"
	Latin       = "Latin"
	Number      = "Number"
	Punctuation = "Punctuation"
	Space       = "Space"
	Undefined   = " "
)

func init() {
	preset, err := LoadPreset()
	if err == nil {
		return
	}
	if err := preset.Save(); err != nil {
		log.Fatal(err)
	}
}

type Preset struct {
	path                string
	GlyphByType         map[string]string `toml:"glyphs by type"`
	GlyphReplacementMap map[string]string `toml:"glyphs replacement map"`
}

var presetDefault = Preset{
	path:                paths.GlyphDataPresetFile(),
	GlyphByType:         defaultTypes(),
	GlyphReplacementMap: defaultReplacements(),
}

func LoadPreset(keys ...string) (*Preset, error) {
	presetToLoad := "default"
	for _, p := range keys {
		presetToLoad = p
	}
	preset := newInfo(presetToLoad)
	expectedPath := preset.path
	if err := preset.Load(); err != nil {
		preset = &presetDefault
		preset.path = expectedPath
		return preset, err
	}
	return preset, nil
}

func newInfo(preset string) *Preset {
	inf := Preset{}
	path := paths.GlyphDataPresetFile(preset)
	inf.path = path
	inf.GlyphByType = make(map[string]string)
	inf.GlyphReplacementMap = make(map[string]string)
	return &inf
}

func (pr *Preset) Save() error {
	bt, err := toml.Marshal(pr)
	if err != nil {
		return fmt.Errorf("failed to marshal: %v", err)
	}
	allBytes := header()
	allBytes = append(allBytes, bt...)
	return os.WriteFile(pr.path, allBytes, 0777)
}

func (pr *Preset) Load() error {
	path := pr.path
	bt, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	err = toml.Unmarshal(bt, pr)
	if err != nil {
		return fmt.Errorf("failed to unmarshal toml")
	}
	pr.path = path
	return nil
}

func header() []byte {
	s := fmt.Sprintf("# last updated at %v\n", time.Now().Format(time.DateTime))
	s += fmt.Sprintf("# file is using toml format. see https://github.com/toml-lang/toml\n\n")
	return []byte(s)
}
