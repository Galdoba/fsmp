package subtitle

import (
	"fmt"
	"strings"

	"github.com/Galdoba/fsmp/internal/glyph"
	"github.com/Galdoba/utils/slicetricks"
	"github.com/fatih/color"
)

type SubtitleEvaluationRules struct {
	widthLimit         int
	linesLimit         int
	minDuration        float64
	maxDuration        float64
	maxGlyphsPerSecond int
}

var DefaultEvaluator = SubtitleEvaluationRules{
	widthLimit:         50,
	linesLimit:         2,
	minDuration:        1.1,
	maxDuration:        7.0,
	maxGlyphsPerSecond: 22,
}

type SubtitleReport struct {
	Render string
	Index  int
	Errs   []error
}

func (ev SubtitleEvaluationRules) Evaluate(st Subtitle) SubtitleReport {
	rep := SubtitleReport{}
	rep.Render, rep.Errs = render(st)
	rep.Errs = collectErrors(rep.Errs,
		assertIndexValue(st),
		assertTimeCodes(st),
	)
	rep.Errs = collectErrors(rep.Errs, assertText(st)...)
	rep.Index = st.Index
	return rep
}

func render(s Subtitle) (string, []error) {
	text := removeFormatTags(s.Text)
	lines := strings.Split(text, "\n")
	errs := []error{}
	preset, err := glyph.LoadPreset()
	if err != nil {
		return "", append(errs, err)
	}
	render := ""
	for ln, line := range lines {
		lineWidth := 0
		for _, gl := range strings.Split(line, "") {
			lineWidth++
			switch preset.GetType(gl) {
			default:
				panic(fmt.Sprintf("undescribed glypg type '%v'", preset.GetType(gl)))
			case glyph.Cyrillic, glyph.Space:
				render += color.CyanString(gl)
			case glyph.Latin:
				render += color.YellowString(gl)
			case glyph.Number, glyph.Punctuation:
				render += color.GreenString(gl)
			case glyph.Undefined:
				render += color.RedString(gl)
				if err := preset.AddUnknownGlyph(gl); err != nil {
					errs = append(errs, fmt.Errorf("failed to add glyph to known types: %v", err))
				}
				errs = append(errs, fmt.Errorf("subtitle %v: line %v: position %v: unknown glyphs met: '%v'", s.Index, ln, lineWidth, gl))
			case glyph.Dangerous:
				render += color.RedString(gl)
				errs = append(errs, fmt.Errorf("subtitle %v: line %v: position %v: dangerous glyphs met: '%v'", s.Index, ln, lineWidth, gl))
			}

		}
		render += "\n"
	}

	if len(errs) > 0 {
		return fmt.Sprintf("%v\n%v --> %v\n%v", s.Index, timestamp(s.StartSeconds), timestamp(s.EndSeconds), render), errs
	}
	return fmt.Sprintf("%v\n%v --> %v\n%v", s.Index, timestamp(s.StartSeconds), timestamp(s.EndSeconds), render), nil
}

func assertIndexValue(s Subtitle) error {
	if s.Index < 1 {
		return fmt.Errorf("index value is less than 1")
	}
	return nil
}

func assertTimeCodes(s Subtitle) error {
	if s.StartSeconds >= s.EndSeconds {
		return fmt.Errorf("title start is not less than end (%v --> %v)", s.StartSeconds, s.EndSeconds)
	}
	if s.StartSeconds < 0 {
		return fmt.Errorf("title start is negative (%v)", s.StartSeconds)
	}
	if s.EndSeconds < 0 {
		return fmt.Errorf("title end is negative (%v)", s.StartSeconds)
	}
	return nil
}

func assertText(s Subtitle) []error {
	errs := []error{}
	if s.Text == "" {
		errs = append(errs, fmt.Errorf("no text present"))
		return errs
	}
	text := removeFormatTags(s.Text)
	lines := strings.Split(text, "\n")
	render := ""
	titleGlyphs := 0
	for l, line := range lines {
		lineGlyths := 0
		for _, gl := range strings.Split(line, "") {
			render += gl
			lineGlyths++
			titleGlyphs++
		}
		if lineGlyths > DefaultEvaluator.widthLimit {
			errs = append(errs, fmt.Errorf("index %v: line %v: width is %v (>%v)", s.Index, l, lineGlyths, DefaultEvaluator.widthLimit))
		}
		render += "\n"
	}
	dur := duration(s)

	if dur > DefaultEvaluator.maxDuration {
		errs = slicetricks.Prepend(errs, fmt.Errorf("index %v: duration = %0.3f (>%v)", s.Index, dur, DefaultEvaluator.maxDuration))
	}
	if dur < DefaultEvaluator.minDuration {
		errs = slicetricks.Prepend(errs, fmt.Errorf("index %v: duration = %0.3f (<%v)", s.Index, dur, DefaultEvaluator.minDuration))
	}
	if ll := len(lines); ll > DefaultEvaluator.linesLimit {
		errs = slicetricks.Prepend(errs, fmt.Errorf("index %v: text lines = %d (>%v)", s.Index, ll, DefaultEvaluator.linesLimit))
	}
	if cps := charactersPerDuration(duration(s), titleGlyphs); cps > DefaultEvaluator.maxGlyphsPerSecond {
		errs = slicetricks.Prepend(errs, fmt.Errorf("index %v: charracters per second = %v (>%v)", s.Index, cps, DefaultEvaluator.maxGlyphsPerSecond))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func duration(s Subtitle) float64 {
	return s.EndSeconds - s.StartSeconds
}

func glyphCount(s Subtitle) int {
	text := s.Text
	lines := strings.Split(text, "\n")
	gc := 0
	for _, line := range lines {
		gc += len(strings.Split(line, ""))
	}
	return gc
}

func charactersPerDuration(duration float64, glyphCount int) int {
	return int(float64(glyphCount) / duration)
}

func collectErrors(errsOut []error, errsIn ...error) []error {
	if errsOut == nil {
		errsOut = []error{}
	}
	for _, err := range errsIn {
		if err != nil {
			errsOut = append(errsOut, err)
		}
	}
	if len(errsOut) > 0 {
		return errsOut
	}
	return nil
}
