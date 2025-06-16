package subtitle

import (
	"fmt"
	"strings"
	"time"
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
	rep := DefaultEvaluator.Evaluate(*st)
	text += rep.Report()
	fmt.Println(text)
	if len(rep.Errs) > 0 {
		time.Sleep(time.Second * 3)
	}
}

func (rep SubtitleReport) Report() string {
	s := rep.Render
	if len(rep.Errs) > 0 {
		s += "\n=======================\n"
		for _, msg := range rep.Errs {
			s += fmt.Sprintf("  - %v\n", msg)
		}
	}

	return s
}

func removeFormatTags(title string) string {
	title = strings.ReplaceAll(title, "<i>", "")
	title = strings.ReplaceAll(title, "</i>", "")
	return title
}
