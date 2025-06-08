package subtitle

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

func Parse(bt []byte) []*Subtitle {
	text := string(bt)
	lines := strings.Split(text, "\n")
	inBlock := false
	// subtitleText := ""
	title := &Subtitle{}
	titles := []*Subtitle{}
	for _, line := range lines {
		switch inBlock {
		case false:
			index, err := strconv.Atoi(line)
			if err != nil {
				continue
			}
			inBlock = true
			title = New()
			title.Index = index
		case true:

		}
	}
	return titles
}

var timeRegex = regexp.MustCompile(`^(\d{2}):(\d{2}):(\d{2}),(\d{3})\s*-->\s*(\d{2}):(\d{2}):(\d{2}),(\d{3})$`)

func parseTimestamps(line string) (float64, float64, error) {
	sub := timeRegex.FindStringSubmatch(line)
	start, end := float64(-1.0), float64(-1.0)
	switch len(sub) {
	case 9:
		val1, _ := strconv.Atoi(sub[1])
		val2, _ := strconv.Atoi(sub[2])
		val3, _ := strconv.Atoi(sub[3])
		val4, _ := strconv.Atoi(sub[4])
		val5, _ := strconv.Atoi(sub[5])
		val6, _ := strconv.Atoi(sub[6])
		val7, _ := strconv.Atoi(sub[7])
		val8, _ := strconv.Atoi(sub[8])
		start = float64((val1*3600)+(val2*60)+val3) + (float64(val4) * 0.001)
		end = float64((val5*3600)+(val6*60)+val7) + (float64(val8) * 0.001)
	default:
		return -1, -1, fmt.Errorf("no timestamps found")
	}
	return start, end, nil
}
