package glyph

import (
	"os"
	"testing"
)

func TestLoadPreset(t *testing.T) {
	pr, err := LoadPreset("test_preset")
	if err == nil {
		t.Errorf("unexpected success Load: %v", err)
	}
	if err := pr.Save(); err != nil {
		t.Errorf("unexpected error Save: %v", err)
	}
	pr, err = LoadPreset("test_preset")
	if err != nil {
		t.Errorf("unexpected error Load: %v", err)
	}

	os.RemoveAll(pr.path)

}
