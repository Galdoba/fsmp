package subrip

type SubRipOption func(*SubRip)

// WithVideoDuration - set reference video duration.
// If value = 0 no duration check will be made.
func WithVideoDuration(seconds float64) SubRipOption {
	return func(sr *SubRip) {
		sr.videoDurationSeconds = seconds
	}
}

// WithMaxLines - set maximum lines per subtitle.
// If value = 0, no limit is applied.
func WithMaxLines(max int) SubRipOption {
	return func(sr *SubRip) {
		sr.maxLinesPerSubtitle = max
	}
}

// WithLineWidth - set maximum line width per subtitle.
// If value = 0, no limit is applied.
func WithLineWidth(max int) SubRipOption {
	return func(sr *SubRip) {
		sr.maxLinesPerSubtitle = max
	}
}
