package subrip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Galdoba/fsmp/pkg/text/charset"
)

func TestNew(t *testing.T) {
	path := `\\192.168.31.4\buffer\IN\CASES\srt\`
	fi, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	max := len(fi)
	for i, f := range fi {
		fmt.Printf("%v/%v\r", i, max)
		sr, err := New(path + f.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		//	if sr.originalLanguage != "" || sr.sourceConfidence != 100 {
		fmt.Println(sr.Report())
		//	}

	}

}

func (sr *SubRip) Report() string {
	s := ""
	s += fmt.Sprintf("file=%v; titles=%v; language=%v; encoding=%v; confidence=%v\n", filepath.Base(sr.Source), len(sr.Subtitles), sr.originalLanguage, sr.originalEncoding, sr.sourceConfidence)
	s += fmt.Sprintf("charaters: Cyrillic=%v; Latin=%v; Numbers=%v; Punctuation=%v; Spaces=%v; UNDEFINED=%v", sr.runesByTypes[charset.Cyrillic], sr.runesByTypes[charset.Latin], sr.runesByTypes[charset.Numbers], sr.runesByTypes[charset.Punctuation], sr.runesByTypes[charset.Space], sr.runesByTypes[charset.Undefined])
	return s
}
