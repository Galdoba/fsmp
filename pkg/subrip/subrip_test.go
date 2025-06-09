package subrip

import (
	"fmt"
	"os"
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
		_, err := New(path + f.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}

	}

}
