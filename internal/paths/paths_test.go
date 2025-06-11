package paths_test

import (
	"fmt"
	"testing"

	"github.com/Galdoba/fsmp/internal/paths"
)

func TestHomeDir(t *testing.T) {
	fmt.Println(paths.ConfigFile())
}
