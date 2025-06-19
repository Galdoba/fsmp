package block

import (
	"fmt"

	"github.com/Galdoba/fsmp/pkg/subrip/subtitle"
)

type BlockChecks struct {
	IDFormat          bool // Целое число >0
	TimestampFormat   bool // Строгий формат "00:00:00,000"
	StartBeforeEnd    bool // Start < End
	DurationRange     bool // MinDuration ≤ (End-Start) ≤ MaxDuration
	LineCount         bool // 1 ≤ строк ≤ MaxLines (обычно 2)
	LineLength        bool // Длина ≤ MaxLineLength (обычно 42)
	CharacterValidity bool // Разрешенные символы (кириллица/латиница/пунктуация)
	MixedAlphabet     bool // Отсутствие слов типа "Ананаc"
	ForbiddenWords    bool // Отсутствие запрещенных терминов
	Punctuation       bool // Правильная пунктуация (нет " ,")
	Capitalization    bool // Нет ВЕРХНЕГО РЕГИСТРА
	ReadingSpeed      bool // Символов/сек ≤ MaxCPS (обычно 20)
	TagValidity       bool // Корректность {\an8}, <font color>
	SoundCues         bool // *хлопок* вместо (хлопок)
}

// CheckID - check if subtitle ID is positive interger.
func CheckID(s subtitle.Subtitle) error {
	if s.Index > 0 {
		return fmt.Errorf("ID must be >0")
	}
	return nil
}

// CheckDuration - check if subtitle start before end.
// Elso check timestamp formatting of each title timepoint.
func CheckTimestamp(s subtitle.Subtitle) error {
	if s.EndSeconds-s.StartSeconds <= 0 {
		return fmt.Errorf("title start must be less than end: (%v !< %v)", s.StartSeconds, s.EndSeconds)
	}

	return nil
}
