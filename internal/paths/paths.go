package paths

import (
	"log"
	"strings"

	"github.com/Galdoba/fsmp/internal/declare"
	"github.com/Galdoba/utils/pathfinder"
)

var appName = declare.APPNAME
var ConfigDir = ""
var PresetDir = ""

func init() {
	plot_config, err := pathfinder.NewPlot(
		pathfinder.WithRoot(pathfinder.CONFIG),
		pathfinder.WithTarget(appName),
	)
	if err != nil {
		log.Fatal(err)
	}
	ConfigDir = pathfinder.Project(plot_config)
	if !strings.HasPrefix(pathfinder.Status(plot_config), "path status: ok") {
		if err := pathfinder.Pave(plot_config); err != nil {
			log.Fatal(err)
		}
	}

	plot_glyph, err := pathfinder.NewPlot(
		pathfinder.WithRoot(pathfinder.CONFIG),
		pathfinder.WithTarget(appName),
		pathfinder.WithLowerLayers("runtime", "glyph_presets"),
	)
	if err != nil {
		log.Fatal(err)
	}
	PresetDir = pathfinder.Project(plot_glyph)
	if !strings.HasPrefix(pathfinder.Status(plot_glyph), "path status: ok") {
		if err := pathfinder.Pave(plot_glyph); err != nil {
			log.Fatal(err)
		}
	}
}

func ConfigFile(keys ...string) string {
	name := "default"
	for _, k := range keys {
		name = k
	}
	return ConfigDir + name + ".toml"
}
func GlyphDataPresetFile(keys ...string) string {
	name := "default"
	for _, k := range keys {
		name = k
	}
	return PresetDir + name + ".toml"
}
