package gen

import (
	"os"
	"testing"
)

func TestHomePath(t *testing.T) {
	h, err := os.UserHomeDir()
	t.Log(h, err)
}
