package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Galdoba/fsmp/internal/paths"
)

func init() {
	preset, err := LoadConfig()
	if err == nil {
		return
	}
	if err := preset.Save(); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	path                 string
	Force_Encoding       string //UTF-8/Windows-1251/None
	Carriage_Return      string //Remove/Force/Ignore
	DefaultGlyphMappings string //key to glyphMapsFile (default none)
	AutoReplaceGlyphs    string //Always/No/Auto/Ask
	RuntimeLog           string //file to log session data
	LogLevel             string
}

/*
Read
Report
Correct
Edit
*/

var configDefault = Config{
	path: paths.ConfigFile(),
}

func LoadConfig(keys ...string) (*Config, error) {
	presetToLoad := "default"
	for _, p := range keys {
		presetToLoad = p
	}
	cfg := newConfig(presetToLoad)
	expectedPath := cfg.path
	if err := cfg.Load(); err != nil {
		cfg = &configDefault
		cfg.path = expectedPath
		return cfg, err
	}
	return cfg, nil
}

func newConfig(preset string) *Config {
	cfg := Config{}
	path := paths.ConfigFile(preset)
	cfg.path = path
	return &cfg
}

func (cfg *Config) Save() error {
	bt, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal: %v", err)
	}
	allBytes := header()
	allBytes = append(allBytes, bt...)
	return os.WriteFile(cfg.path, allBytes, 0777)
}

func (cfg *Config) Load() error {
	path := cfg.path
	bt, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	err = toml.Unmarshal(bt, cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal toml")
	}
	cfg.path = path
	return nil
}

func header() []byte {
	s := fmt.Sprintf("# file is using toml format. see https://github.com/toml-lang/toml\n\n")
	return []byte(s)
}
