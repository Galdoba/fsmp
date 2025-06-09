package glyph

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Info struct {
	path                string
	GlyphByType         map[string]string `toml:"glyphs by type"`
	GlyphReplacementMap map[string]string `toml:"glyphs replacement map"`
}

func LoadPreset(preset ...string) (*Info, error) {
	presetToLoad := "default"
	for _, p := range preset {
		presetToLoad = p
	}
	info, err := newInfo(presetToLoad)
	if err != nil {
		return nil, err
	}
	if len(info.GlyphByType) == 0 {
		info.GlyphByType = defaultTypes()
		info.GlyphReplacementMap["glyph"] = "glyph"
	}
	bt, _ := toml.Marshal(info)
	fmt.Println(string(bt))
	return info, nil
}

func defaultTypes() map[string]string {
	typesMap := make(map[string]string)
	typesMap["А"] = "Cyrillic"
	typesMap["Б"] = "Cyrillic"
	typesMap["В"] = "Cyrillic"
	typesMap["Г"] = "Cyrillic"
	typesMap["Д"] = "Cyrillic"
	typesMap["Е"] = "Cyrillic"
	typesMap["Ё"] = "Cyrillic"
	typesMap["Ж"] = "Cyrillic"
	typesMap["З"] = "Cyrillic"
	typesMap["И"] = "Cyrillic"
	typesMap["Й"] = "Cyrillic"
	typesMap["К"] = "Cyrillic"
	typesMap["Л"] = "Cyrillic"
	typesMap["М"] = "Cyrillic"
	typesMap["Н"] = "Cyrillic"
	typesMap["О"] = "Cyrillic"
	typesMap["П"] = "Cyrillic"
	typesMap["Р"] = "Cyrillic"
	typesMap["С"] = "Cyrillic"
	typesMap["Т"] = "Cyrillic"
	typesMap["У"] = "Cyrillic"
	typesMap["Ф"] = "Cyrillic"
	typesMap["Х"] = "Cyrillic"
	typesMap["Ц"] = "Cyrillic"
	typesMap["Ч"] = "Cyrillic"
	typesMap["Ш"] = "Cyrillic"
	typesMap["Щ"] = "Cyrillic"
	typesMap["Ь"] = "Cyrillic"
	typesMap["Ы"] = "Cyrillic"
	typesMap["Ъ"] = "Cyrillic"
	typesMap["Э"] = "Cyrillic"
	typesMap["Ю"] = "Cyrillic"
	typesMap["Я"] = "Cyrillic"
	typesMap["а"] = "Cyrillic"
	typesMap["б"] = "Cyrillic"
	typesMap["в"] = "Cyrillic"
	typesMap["г"] = "Cyrillic"
	typesMap["д"] = "Cyrillic"
	typesMap["е"] = "Cyrillic"
	typesMap["ё"] = " "
	typesMap["ж"] = "Cyrillic"
	typesMap["з"] = "Cyrillic"
	typesMap["и"] = "Cyrillic"
	typesMap["й"] = "Cyrillic"
	typesMap["к"] = "Cyrillic"
	typesMap["л"] = "Cyrillic"
	typesMap["м"] = "Cyrillic"
	typesMap["н"] = "Cyrillic"
	typesMap["о"] = "Cyrillic"
	typesMap["п"] = "Cyrillic"
	typesMap["р"] = "Cyrillic"
	typesMap["с"] = "Cyrillic"
	typesMap["т"] = "Cyrillic"
	typesMap["у"] = "Cyrillic"
	typesMap["ф"] = "Cyrillic"
	typesMap["х"] = "Cyrillic"
	typesMap["ц"] = "Cyrillic"
	typesMap["ч"] = "Cyrillic"
	typesMap["ш"] = "Cyrillic"
	typesMap["щ"] = "Cyrillic"
	typesMap["ь"] = "Cyrillic"
	typesMap["ы"] = "Cyrillic"
	typesMap["ъ"] = "Cyrillic"
	typesMap["э"] = "Cyrillic"
	typesMap["ю"] = "Cyrillic"
	typesMap["я"] = "Cyrillic"
	typesMap["A"] = "Latin"
	typesMap["B"] = "Latin"
	typesMap["C"] = "Latin"
	typesMap["D"] = "Latin"
	typesMap["E"] = "Latin"
	typesMap["F"] = "Latin"
	typesMap["G"] = "Latin"
	typesMap["H"] = "Latin"
	typesMap["I"] = "Latin"
	typesMap["J"] = "Latin"
	typesMap["K"] = "Latin"
	typesMap["L"] = "Latin"
	typesMap["M"] = "Latin"
	typesMap["N"] = "Latin"
	typesMap["O"] = "Latin"
	typesMap["P"] = "Latin"
	typesMap["Q"] = "Latin"
	typesMap["R"] = "Latin"
	typesMap["S"] = "Latin"
	typesMap["T"] = "Latin"
	typesMap["U"] = "Latin"
	typesMap["V"] = "Latin"
	typesMap["W"] = "Latin"
	typesMap["X"] = "Latin"
	typesMap["Y"] = "Latin"
	typesMap["Z"] = "Latin"
	typesMap["a"] = "Latin"
	typesMap["b"] = "Latin"
	typesMap["c"] = "Latin"
	typesMap["d"] = "Latin"
	typesMap["e"] = "Latin"
	typesMap["f"] = "Latin"
	typesMap["g"] = "Latin"
	typesMap["h"] = "Latin"
	typesMap["i"] = "Latin"
	typesMap["j"] = "Latin"
	typesMap["k"] = "Latin"
	typesMap["l"] = "Latin"
	typesMap["m"] = "Latin"
	typesMap["n"] = "Latin"
	typesMap["o"] = "Latin"
	typesMap["p"] = "Latin"
	typesMap["q"] = "Latin"
	typesMap["r"] = "Latin"
	typesMap["s"] = "Latin"
	typesMap["t"] = "Latin"
	typesMap["u"] = "Latin"
	typesMap["v"] = "Latin"
	typesMap["w"] = "Latin"
	typesMap["x"] = "Latin"
	typesMap["y"] = "Latin"
	typesMap["z"] = "Latin"
	typesMap["0"] = "Numbers"
	typesMap["1"] = "Numbers"
	typesMap["2"] = "Numbers"
	typesMap["3"] = "Numbers"
	typesMap["4"] = "Numbers"
	typesMap["5"] = "Numbers"
	typesMap["6"] = "Numbers"
	typesMap["7"] = "Numbers"
	typesMap["8"] = "Numbers"
	typesMap["9"] = "Numbers"
	typesMap["%"] = "Numbers"
	typesMap["№"] = "Numbers"
	typesMap[","] = "Punctuation"
	typesMap["."] = "Punctuation"
	typesMap["!"] = "Punctuation"
	typesMap["?"] = "Punctuation"
	typesMap[":"] = "Punctuation"
	typesMap[";"] = "Punctuation"
	typesMap["<"] = "Punctuation"
	typesMap[">"] = "Punctuation"
	typesMap["-"] = "Punctuation"
	typesMap["+"] = "Punctuation"
	typesMap["="] = "Punctuation"
	typesMap["*"] = "Punctuation"
	typesMap["/"] = "Punctuation"
	typesMap[`\`] = "Punctuation"
	typesMap["["] = "Punctuation"
	typesMap["]"] = "Punctuation"
	typesMap["{"] = "Punctuation"
	typesMap["}"] = "Punctuation"
	typesMap["'"] = "Punctuation"
	typesMap["@"] = "Punctuation"
	typesMap["#"] = "Punctuation"
	typesMap[" "] = "Space"
	//typesMap["\n"] = "Space"
	//typesMap["\t"] = "Space"
	return typesMap
}

func newInfo(preset string) (*Info, error) {
	inf := Info{}
	path, err := dbPath(preset)
	if err != nil {
		return nil, fmt.Errorf("failed to find glyph info path: %v", err)
	}
	inf.path = path
	inf.GlyphByType = make(map[string]string)
	inf.GlyphReplacementMap = make(map[string]string)
	return &inf, nil
}

func dbPath(preset string) (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to find glyph info path: %v", err)
	}
	path += ".config/fsmp/glyph_presets/"
	switch preset {
	case "", "default":
		path += "default.toml"
	default:
		path += preset + ".toml"
	}
	return filepath.ToSlash(path), nil
}
