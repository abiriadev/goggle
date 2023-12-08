package index

import (
	"strings"
	"testing"
)

func TestModuleIndexParser(t *testing.T) {
	raw := "{}\n{}"

	_, err := ParseModuleIndex(strings.NewReader(raw))

	if err != nil {
		t.Fatal(err)
	}
}
