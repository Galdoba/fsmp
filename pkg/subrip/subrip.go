package subrip

import (
	"github.com/Galdoba/fsmp/pkg/subrip/subtitle"
)

// SubRip represents srt file or subtitle stream in media file.
type SubRip struct {
	//Filepath to original file.
	Source string

	//ffmpeg stream key (example ':s:0' for srt file)
	Key string

	//slice of titles
	Subtitles []subtitle.Subtitle

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
func New(path string, options ...SubRipOption) error {
	sr := SubRip{
		Source:               path,
		Key:                  "",
		Subtitles:            []subtitle.Subtitle{},
		maxLinesPerSubtitle:  0,
		maxLineWidth:         0,
		videoDurationSeconds: 0,
	}
	for _, modify := range options {
		modify(&sr)
	}
	return nil
}

func (sr *SubRip) Evaluate() error {
	return nil
}
