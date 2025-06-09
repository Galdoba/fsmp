package subrip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
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
		if sr.originalLanguage != "" || sr.sourceConfidence != 100 {
			fmt.Println(sr.Report())
		}

	}

}

func (sr *SubRip) Report() string {
	s := ""
	s += fmt.Sprintf("file=%v; titles=%v; language=%v; encoding=%v; confidence=%v", filepath.Base(sr.Source), len(sr.Subtitles), sr.originalLanguage, sr.originalEncoding, sr.sourceConfidence)
	return s
}
