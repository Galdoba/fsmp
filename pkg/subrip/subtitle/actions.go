package subtitle

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
