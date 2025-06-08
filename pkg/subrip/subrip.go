package subrip

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/fsmp/pkg/subrip/subtitle"
)

// SubRip represents srt file or subtitle stream in media file.
type SubRip struct {
	//Filepath to original file.
	Source string

	//ffmpeg stream key (example ':s:0' for srt file)
	Key string

	//slice of titles
	Subtitles []*subtitle.Subtitle

	maxLinesPerSubtitle        int
	maxLineWidth               int
	videoDurationSeconds       float64
	minimumSubtitleDuration    float64
	maximumSubtitleDuration    float64
	minimumGapBetweenSubtitles float64
	maximumGapBetweenSubtitles float64
	maximumCharactersPerSecond float64
	strictCharacterValidation  bool
	autoCorrectMixedAlphabet   bool
	evaluationErrors           []error
}

// New - create new SubRip structure
func New(path string, options ...SubRipOption) (*SubRip, error) {
	sr := SubRip{
		Source:               path,
		Key:                  "",
		Subtitles:            []*subtitle.Subtitle{},
		maxLinesPerSubtitle:  0,
		maxLineWidth:         0,
		videoDurationSeconds: 0,
	}
	for _, modify := range options {
		modify(&sr)
	}
	bt, err := readFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't create SubRip: %v", err)
	}
	sr.Subtitles = subtitle.Parse(bt)
	return &sr, nil
}

func (sr *SubRip) Evaluate() error {
	return nil
}

func readFile(path string) ([]byte, error) {
	if !strings.HasSuffix(path, ".srt") {
		return nil, fmt.Errorf("path is not srt file")
	}
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	return bt, nil
}
