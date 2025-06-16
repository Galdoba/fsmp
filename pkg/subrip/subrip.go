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

type FileReport struct {
	sources        map[string]*SubRip
	renders        []string
	FilewiseErrors []error
	SubtitleErrors map[int][]error
}

func reportError(err error) FileReport {
	return FileReport{FilewiseErrors: append([]error{}, fmt.Errorf("nothing to report"))}
}

type SubRipEvaluator struct {
}

func (se *SubRipEvaluator) Evaluate(paths ...string) FileReport {
	if len(paths) < 1 {
		return reportError(fmt.Errorf("nothing to report"))
	}
	srep := FileReport{}
	srep.sources = make(map[string]*SubRip)
	srep.SubtitleErrors = make(map[int][]error)
	for _, path := range paths {
		//fmt.Println(path)
		sr, err := New(path)
		if err != nil {
			return reportError(fmt.Errorf("%v is not a subtitle", path))
		}
		if err := assertTimeGaps(*sr); err != nil {
			srep.FilewiseErrors = append(srep.FilewiseErrors, err)
		}
		if err := assertIndexes(*sr); err != nil {
			srep.FilewiseErrors = append(srep.FilewiseErrors, err)
		}
		srep.sources[path] = sr
		for _, title := range sr.Subtitles {
			rep := subtitle.DefaultEvaluator.Evaluate(*title)
			if len(rep.Errs) > 0 {
				srep.SubtitleErrors[rep.Index] = append([]error{fmt.Errorf(rep.Render)}, rep.Errs...)
			}
		}
	}
	//index positions

	return srep
}

func (frep *FileReport) Report() string {
	if len(frep.FilewiseErrors) == 0 && len(frep.SubtitleErrors) == 0 {
		return fmt.Sprintf("is valid SubRip")
	}
	s := ""
	if len(frep.FilewiseErrors) > 0 {
		s += "file errors:\n"
		for _, err := range frep.FilewiseErrors {
			s += fmt.Sprintf("  %v\n", err)
		}
	}
	if len(frep.SubtitleErrors) >= 0 {
		s += "by index errors:\n"
		for i := 0; i < 10000; i++ {
			if errs, ok := frep.SubtitleErrors[i]; ok {
				for _, err := range errs {
					s += fmt.Sprintf("\n%v\n", err)
				}
			}
		}

	}
	return s
}

func assertIndexes(sr SubRip) error {
	unmatchedIndexes := []int{}
	for i, s := range sr.Subtitles {
		switch s.Index == i+1 {
		case true:
		case false:
			unmatchedIndexes = append(unmatchedIndexes, s.Index)
		}
	}
	if len(unmatchedIndexes) > 0 {
		return fmt.Errorf("file contains %v invalid indexes:\n%v", len(unmatchedIndexes), unmatchedIndexes)
	}
	return nil
}

func assertTimeGaps(sr SubRip) error {
	gaps := []string{}
	lastEnd := 0.0
	for i, s := range sr.Subtitles {
		gap := s.StartSeconds - lastEnd
		if gap >= 35*60 {
			switch i {
			case 0:
				gaps = append(gaps, fmt.Sprintf("index [start=>%v]: gap %0.3f seconds", s.Index, gap))
			default:
				gaps = append(gaps, fmt.Sprintf("index [%v=>%v]: gap %0.3f seconds", sr.Subtitles[i-1].Index, s.Index, gap))
			}
		}
		lastEnd = s.EndSeconds
	}
	if len(gaps) > 0 {
		s := "gaps:"
		for _, gap := range gaps {
			s += "\n" + gap
		}
		return fmt.Errorf(s)
	}
	return nil
}
