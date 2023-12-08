package index

import (
	"strings"
	"testing"
	"time"
)

func TestModuleIndexParser(t *testing.T) {
	raw := "{}\n{}\n{\"Timestamp\":\"2019-04-10T19:08:52.997264Z\"}"

	v, err := ParseModuleIndex(strings.NewReader(raw))

	if err != nil {
		t.Fatal(err)
	}

	if len(v) != 3 {
		t.Fail()
	}

	if v[2].Timestamp != time.Date(2019, 04, 10, 19, 8, 52, 997264000, time.UTC) {
		t.Fatal(v[2].Timestamp)
	}
}
