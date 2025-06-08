package subrip

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	sr, err := New(`C:\Workspace\BUFFER_IN\srt\Storona_zashity_s01e06_PRT240819074341_SER_04869_18.RUS.srt`)
	fmt.Println(err)
	for _, st := range sr.Subtitles {
		st.Print()
	}

}
