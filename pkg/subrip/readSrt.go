package subrip

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/Galdoba/fsmp/pkg/text/charset"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
)

type readReport struct {
	bt             []byte
	characterTypes map[int]int
	encoding       chardet.Result
	err            error
}

// readSRT - Reads file as UTF-8.
// If encoding is not UTF-8 attempt to decode to UTF-8.
// Remove BOM.
// Report error on reading/decoding process
func readSRT(path string) readReport {
	if !strings.HasSuffix(path, ".srt") {
		return readError("path is not srt file")
	}
	bt, err := os.ReadFile(path)
	if err != nil {
		return readError("read error: %v", err)
	}
	encoding, err := detectEncoding(bt)
	if err != nil {
		return readError("%v", err)
	}
	switch encoding.Charset {
	case "windows-1251":
		utf8bt, err := decodeWindows1251Toutf8(bt)
		if err != nil {
			return readError("%v", err)
		}
		bt = utf8bt
	case "UTF-8-BOM":
		bt = removeBOM(bt)
	case "UTF-8":
	default:
		return readError("unsupported charset: %v", encoding.Charset)
	}
	bt = removeCR(bt)
	if !utf8.Valid(bt) {
		return readError("not a valid UTF-8 encoding")
	}
	return readReport{bt: bt, encoding: *encoding, characterTypes: countCharaterTypes(bt), err: nil}
}

func readError(format string, args ...interface{}) readReport {
	return readReport{err: fmt.Errorf(format, args...)}
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

func removeCR(bt []byte) []byte {
	return bytes.ReplaceAll(bt, []byte("\r"), []byte(""))
}

func countCharaterTypes(bt []byte) map[int]int {
	text := string(bt)
	characterTypes := make(map[int]int)
	for _, r := range text {
		characterTypes[charset.CharacterType[r]]++
	}
	return characterTypes
}
