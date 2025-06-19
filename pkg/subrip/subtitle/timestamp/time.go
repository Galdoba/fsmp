package timestamp

import (
	"fmt"
	"regexp"
	"strconv"
)

func ToString(time float64) string {
	milliseconds := int64(seconds * 1000)
	hh := 0
	for milliseconds > 3600000 {
		milliseconds -= 3600000
		hh++
	}
	mm := 0
	for milliseconds > 60000 {
		milliseconds -= 60000
		mm++
	}
	ss := 0
	for milliseconds >= 1000 {
		milliseconds -= 1000
		ss++
	}
	ms := int(milliseconds)
	return fmt.Sprintf("%v:%v:%v,%v", hhmmssToString(hh), hhmmssToString(mm), hhmmssToString(ss), millisecondsToString(ms))
}

var timeRegex = regexp.MustCompile(`^(\d{2}):(\d{2}):(\d{2}),(\d{3})\s*-->\s*(\d{2}):(\d{2}):(\d{2}),(\d{3})$`)

// parseTimestamps - Parse start and end time from srt line. It return error in no timestamps found.
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

func hhmmssToString(num int) string {
	s := fmt.Sprintf("%v", num)
	for len(s) < 2 {
		s = "0" + s
	}
	return s
}

func millisecondsToString(num int) string {
	s := fmt.Sprintf("%v", num)
	for len(s) < 3 {
		s = "0" + s
	}
	return s
}
