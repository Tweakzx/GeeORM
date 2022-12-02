package log

import (
	"testing"
)

func TestSetLevl(t *testing.T) {
	SetLevel(0)
	SetLevel(1)
	SetLevel(2)
}
