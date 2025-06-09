package subrip

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/Galdoba/fsmp/pkg/subrip/subtitle"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
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
	bt, err := readFileAsUTF8(path)
	if err != nil {
		return nil, fmt.Errorf("can't create SubRip: %v", err)
	}
	sr.Subtitles = subtitle.Parse(bt)
	return &sr, nil
}

func (sr *SubRip) Evaluate() error {
	return nil
}

func readFileAsUTF8(path string) ([]byte, error) {
	if !strings.HasSuffix(path, ".srt") {
		return nil, fmt.Errorf("path is not srt file")
	}
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	encoding, err := detectEncoding(bt)
	if err != nil {
		return nil, err
	}
	switch encoding.Charset {
	case "windows-1251":
		utf8bt, err := decodeWindows1251Toutf8(bt)
		if err != nil {
			return nil, err
		}
		bt = utf8bt
	case "UTF-8-BOM":
		bt = removeBOM(bt)
	case "UTF-8":
	default:
		return nil, fmt.Errorf("unsupported charset: %v", encoding.Charset)
	}
	if !utf8.Valid(bt) {
		return nil, fmt.Errorf("not a valid UTF-8 encoding")
	}
	return bt, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func detectEncoding(bt []byte) (*chardet.Result, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(bt)
	if err != nil {
		return nil, fmt.Errorf("encoding detection failed: %v", err)
	}
	if bytes.HasPrefix(bt, []byte{0xEF, 0xBB, 0xBF}) {
		result.Charset = result.Charset + "-BOM"
	}

	return result, nil
}

func removeBOM(data []byte) []byte {
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return data[3:]
	}
	return data
}

func decodeWindows1251Toutf8(bt []byte) ([]byte, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(bt)
	if err != nil {
		return nil, fmt.Errorf("windows-1251 decoding failed: %v", err)
	}
	return out, nil
}
