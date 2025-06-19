package subrip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Galdoba/fsmp/pkg/text/charset"
)

func TestNew(t *testing.T) {
	path := `/home/galdoba/workbench/fsmp/srt/`
	fi, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	lineWidths := make(map[int]int)
	max := len(fi)
	srEval := SubRipEvaluator{}
	for i, f := range fi {
		if i < -1 {
			continue
		}

		fmt.Printf("%v/%v\r", i, max)
		// sr, err := New(path + f.Name())
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
		rep := srEval.Evaluate(path + f.Name())
		fmt.Println(f.Name())
		text := rep.Report()
		fmt.Println(text)
		fmt.Println("\n\n====================")
		if text != "is valid SubRip" {
			// time.Sleep(time.Second * 5)
		}
		// fmt.Print("\033[H\033[2J")
		//fmt.Println(rep)

		// for _, st := range sr.Subtitles {

		//
		// 	// rep := subtitle.DefaultEvaluator.Evaluate(*st)

		// 	// print := false
		// 	// for _, lw := range rep.LineWidths {
		// 	// 	if lw > 45 {
		// 	// 		print = true
		// 	// 	}
		// 	// 	if lw == 1 {
		// 	// 		print = true
		// 	// 	}
		// 	// 	lineWidths[lw]++
		// 	// }
		// 	// if print {
		// 	// 	fmt.Printf("file: (%v) %v\n", i, f.Name())
		// 	// 	st.Print()
		// 	// 	time.Sleep(time.Second)
		// 	// }
		// 	// st.Print()
		// }
		//fmt.Print("\033[H\033[2J")

		//	if sr.originalLanguage != "" || sr.sourceConfidence != 100 {

		//	}

	}
	for l := 0; l < 500; l++ {
		if val, ok := lineWidths[l]; ok {
			fmt.Printf("width %v : %v\n", l, val)
		}
	}

}

func (sr *SubRip) Report() string {
	s := ""
	s += fmt.Sprintf("file=%v; titles=%v; language=%v; encoding=%v; confidence=%v\n", filepath.Base(sr.Source), len(sr.Subtitles), sr.originalLanguage, sr.originalEncoding, sr.sourceConfidence)
	s += fmt.Sprintf("charaters: Cyrillic=%v; Latin=%v; Numbers=%v; Punctuation=%v; Spaces=%v; UNDEFINED=%v", sr.runesByTypes[charset.Cyrillic], sr.runesByTypes[charset.Latin], sr.runesByTypes[charset.Numbers], sr.runesByTypes[charset.Punctuation], sr.runesByTypes[charset.Space], sr.runesByTypes[charset.Undefined])
	return s
}

/*
width 1 : 42
width 2 : 455
width 3 : 4687
width 4 : 7115
width 5 : 12428
width 6 : 14387
width 7 : 14598
width 8 : 17635
width 9 : 18323
width 10 : 18435
width 11 : 18377
width 12 : 20472
width 13 : 22493
width 14 : 25464
width 15 : 27500
width 16 : 28324
width 17 : 30779
width 18 : 31528
width 19 : 33074
width 20 : 34167
width 21 : 34564
width 22 : 35858
width 23 : 35835
width 24 : 35399
width 25 : 35293
width 26 : 34567
width 27 : 32861
width 28 : 32064
width 29 : 30409
width 30 : 29354
width 31 : 27895
width 32 : 26608
width 33 : 24412
width 34 : 22754
width 35 : 20896
width 36 : 19539
width 37 : 18004
width 38 : 15658
width 39 : 1611
width 40 : 1450
width 41 : 1248
width 42 : 832
width 43 : 242
width 44 : 190
width 45 : 164
width 46 : 170
width 47 : 145
width 48 : 127
width 49 : 145
width 50 : 99
width 51 : 95
width 52 : 78
width 53 : 79
width 54 : 82
width 55 : 64
width 56 : 52
width 57 : 40
width 58 : 45
width 59 : 35
width 60 : 36
width 61 : 39
width 62 : 31
width 63 : 33
width 64 : 22
width 65 : 26
width 66 : 38
width 67 : 27
width 68 : 20
width 69 : 21
width 70 : 12
width 71 : 13
width 72 : 18
width 73 : 18
width 74 : 11
width 75 : 11
width 76 : 9
width 77 : 9
width 78 : 14
width 79 : 7
width 80 : 3
width 81 : 3
width 82 : 3
width 83 : 4
width 84 : 3
width 85 : 5
width 86 : 1
width 87 : 2
width 88 : 1
width 89 : 3
width 92 : 5
width 93 : 1
width 94 : 1
width 95 : 1
width 98 : 1
width 102 : 2



file: FILM_Vozvrashenie_tigra_TIGER_3_SONGS_blanked.srt
53
01:18:44,618 --> 01:18:44,619


=======================
*/
