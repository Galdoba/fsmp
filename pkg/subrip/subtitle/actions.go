package subtitle

import "fmt"

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
	text += fmt.Sprintf("%v\n", st.Text)
	fmt.Println(text)
}
