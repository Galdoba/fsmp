package subtitle

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galdoba/fsmp/internal/glyph"
	"github.com/fatih/color"
)

// Adjust - move subtitle to n seconds on the timeline.
func (st *Subtitle) Adjust(n float64) {
	st.StartSeconds += n
	st.EndSeconds += n
}

// Scale - recalculate subtitle's timecodes by factor of n.
func (st *Subtitle) Scale(n float64) {
	st.StartSeconds = st.StartSeconds * n
	st.EndSeconds = st.EndSeconds * n
}

// Put - set subtitle's timecodes directly.
func (st *Subtitle) Put(start, end float64) {
	st.StartSeconds = start
	st.EndSeconds += end
}

func (st *Subtitle) Print() {
	text := ""
	text += fmt.Sprintf("%v\n", st.Index)
	text += fmt.Sprintf("%v --> %v\n", timestamp(st.StartSeconds), timestamp(st.EndSeconds))
	rep := st.Evaluate()
	text += rep.Report()
	fmt.Println(text)
	if len(rep.messages) > 0 {
		time.Sleep(time.Second * 3)
	}
}

type SubtitleReport struct {
	render     string
	lines      int
	glyphs     int
	messages   []string
	LineWidths []int
	stop       bool
}

func (st *Subtitle) Evaluate() SubtitleReport {
	rep := SubtitleReport{}
	text := removeFormatTags(st.Text)
	lines := strings.Split(text, "\n")
	rep.lines = len(lines)
	if rep.lines > 2 {
		rep.messages = append(rep.messages, fmt.Sprintf("lines in title: %v (limit=2)", rep.lines))
	}
	if rep.lines < 1 {
		rep.messages = append(rep.messages, fmt.Sprintf("lines in title: %v (should be atleast 1)", rep.lines))
	}
	preset, err := glyph.LoadPreset()
	if err != nil {
		panic("no preset: " + fmt.Sprintf("%v", err))
	}

	for lileNumber, line := range lines {
		lineWidth := 0
		for _, gl := range strings.Split(line, "") {
			switch preset.GetType(gl) {
			case glyph.Cyrillic, glyph.Space:
				rep.render += color.CyanString(gl)
			case glyph.Latin:
				rep.render += color.YellowString(gl)

			case glyph.Number, glyph.Punctuation:
				rep.render += color.GreenString(gl)
			case glyph.Undefined:
				rep.render += color.RedString(gl)
				preset.AddUnknownGlyph(gl)
				rep.messages = append(rep.messages, fmt.Sprintf("unknown glyph added: '%v'", gl))
			case glyph.Dangerous:
				rep.render += color.RedString(gl)
				rep.messages = append(rep.messages, fmt.Sprintf("dangerous glyph detected: '%v'", gl))
			}
			rep.glyphs++
			lineWidth++
		}
		if lineWidth > 40 {
			rep.messages = append(rep.messages, fmt.Sprintf("line %v contains %v glyphs (should be <= 40)", lileNumber, lineWidth))
		}
		rep.LineWidths = append(rep.LineWidths, lineWidth)
		rep.render += "\n"
	}

	return rep
}

func (rep SubtitleReport) Report() string {
	s := rep.render
	s += "\n=======================\n"
	s += fmt.Sprintf("  lines : %v\n", rep.lines)
	s += fmt.Sprintf("  glyphs: %v\n", rep.glyphs)
	for _, msg := range rep.messages {
		s += fmt.Sprintf("  - %v\n", msg)
	}
	return s
}

func removeFormatTags(title string) string {
	title = strings.ReplaceAll(title, "<i>", "")
	title = strings.ReplaceAll(title, "</i>", "")
	return title
}
