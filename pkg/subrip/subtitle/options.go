package subtitle

// WithIndex - set index.
func WithIndex(i int) SubtitleOption {
	return func(s *Subtitle) {
		s.Index = i
	}
}

// WithStart - set Start point.
func WithStart(start float64) SubtitleOption {
	return func(s *Subtitle) {
		s.StartSeconds = start
	}
}

// Withend - set end point.
func WithEnd(end float64) SubtitleOption {
	return func(s *Subtitle) {
		s.EndSeconds = end
	}
}

// WithText- set subtile text.
func WithText(t string) SubtitleOption {
	return func(s *Subtitle) {
		s.Text = t
	}
}
