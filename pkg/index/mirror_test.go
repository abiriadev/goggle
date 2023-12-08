package index

import (
	"strings"
	"testing"
)

func TestModuleIndexParser(t *testing.T) {
	raw := "{}\n{}\n{}"

	v, err := ParseModuleIndex(strings.NewReader(raw))

	if err != nil {
		t.Fatal(err)
	}

	if len(v) != 3 {
		t.Fatal()
	}
}
