package retroport

import (
	"testing"
)

type eg struct {
	b []byte
	s string
}

func TestUpdate(t *testing.T) {
	c := &SNES{}

	data := []eg{
		eg{[]byte{128, 128, 0, 0}, ""},
		eg{[]byte{000, 128, 0, 0}, "◀"},
		eg{[]byte{255, 128, 0, 0}, "▶"},
		eg{[]byte{128, 000, 0, 0}, "▲"},
		eg{[]byte{128, 255, 0, 0}, "▼"},
		eg{[]byte{128, 128, 0, 4}, "L"},
		eg{[]byte{128, 128, 0, 8}, "R"},
		eg{[]byte{128, 128, 0, 12}, "LR"},
		eg{[]byte{128, 128, 3, 3}, "XYAB"},
		eg{[]byte{128, 128, 12, 0}, "_S"},
	}

	for _, eg := range data {
		c.update(eg.b)
		s := c.Buttons()
		if s != eg.s {
			t.Errorf("bytes %v resulted in state %#v, expected %#v", eg.b, s, eg.s)
		}
	}
}
