package subrip

import (
	"fmt"

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

	originalEncoding string
	originalLanguage string
	sourceConfidence int
	runesByTypes     map[int]int

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
	rr := readSRT(path)
	if rr.err != nil {
		return nil, fmt.Errorf("can't create SubRip: %v", rr.err)
	}
	sr.originalEncoding = rr.encoding.Charset
	sr.originalLanguage = rr.encoding.Language
	sr.sourceConfidence = rr.encoding.Confidence
	sr.runesByTypes = rr.characterTypes
	sr.Subtitles = subtitle.Parse(rr.bt)
	return &sr, nil
}

func (sr *SubRip) Evaluate() error {
	return nil
}

// func readFileAsUTF8(path string) ([]byte, error) {
// 	if !strings.HasSuffix(path, ".srt") {
// 		return nil, fmt.Errorf("path is not srt file")
// 	}
// 	bt, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("read error: %v", err)
// 	}
// 	encoding, err := detectEncoding(bt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch encoding.Charset {
// 	case "windows-1251":
// 		utf8bt, err := decodeWindows1251Toutf8(bt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		bt = utf8bt
// 	case "UTF-8-BOM":
// 		bt = removeBOM(bt)
// 	case "UTF-8":
// 	default:
// 		return nil, fmt.Errorf("unsupported charset: %v", encoding.Charset)
// 	}
// 	if !utf8.Valid(bt) {
// 		return nil, fmt.Errorf("not a valid UTF-8 encoding")
// 	}
// 	return bt, nil
// }
