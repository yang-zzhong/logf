package logf

import (
	"testing"
)

func TestLog(t *testing.T) {

	SetFilenamePrefix("test-")
	SetFormat("2006-01")
	SetPath("./")

	Printf("hello world: %s\n", "1111111")
}
