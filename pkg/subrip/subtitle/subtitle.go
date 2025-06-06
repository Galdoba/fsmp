package subtitle

const (
	index_unset = -1
	time_unset  = -3.14
)

// Subtitle - Basis of SubRip format. This struct is a single 'message' to appear on video player during playback.
type Subtitle struct {
	// Index is unique subtitle id. It starts from 1.
	// Next subtitle must have index = n+1. Previous - n-1.
	Index int `json:"index"`

	// StartSeconds and EndSeconds are Subtitle's timestamps.
	// EndSeconds MUST be greater than StartSeconds
	StartSeconds float64 `json:"start"`
	EndSeconds   float64 `json:"end"`

	// Text - is a message body to appear on screen.
	Text string `json:"text"`
}

// New - creates new subtitle struct
func New(options ...SubtitleOption) *Subtitle {
	st := Subtitle{}
	st.Index = index_unset
	st.StartSeconds = time_unset
	st.EndSeconds = time_unset
	for _, modify := range options {
		modify(&st)
	}
	return &st
}
