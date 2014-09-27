package retroport

import (
	"bytes"
	"fmt"
	"io"
)

type SNES struct {
	r      io.Reader
	Up     bool
	Down   bool
	Left   bool
	Right  bool
	X      bool
	Y      bool
	A      bool
	B      bool
	L      bool
	R      bool
	Select bool
	Start  bool
}

// MakeSNES returns a pointer to a SNES controller for the given io.reader.
func MakeSNES(reader io.Reader) *SNES {
	return &SNES{r: reader}
}

// Buttons returns a string containing a representation of the buttons which are
// currently pressed. It's just informational, don't rely on the format.
func (c *SNES) Buttons() string {
	var s bytes.Buffer

	if c.Up {
		s.WriteRune('▲')
	}
	if c.Down {
		s.WriteRune('▼')
	}
	if c.Left {
		s.WriteRune('◀')
	}
	if c.Right {
		s.WriteRune('▶')
	}
	if c.X {
		s.WriteRune('X')
	}
	if c.Y {
		s.WriteRune('Y')
	}
	if c.A {
		s.WriteRune('A')
	}
	if c.B {
		s.WriteRune('B')
	}
	if c.L {
		s.WriteRune('L')
	}
	if c.R {
		s.WriteRune('R')
	}
	if c.Select {
		s.WriteRune('_')
	}
	if c.Start {
		s.WriteRune('S')
	}

	return s.String()
}

func (c *SNES) String() string {
	return fmt.Sprintf("&SNES{%s}", c.Buttons())
}

// Run loops forever, keeping the state of the controller up to date. This
// should be called in a goroutine.
func (c *SNES) Run() {
	buf := make([]byte, 4, 4)

	for {
		c.r.Read(buf)
		c.update(buf)
	}
}

// Any returns true if any button is currently pressed.
func (c *SNES) Any() bool {
	return c.Up || c.Down || c.Left || c.Right || c.X || c.Y || c.A || c.B || c.L || c.R || c.Select || c.Start
}

// update overwrites the state of the controller by parsing the specified byte
// array, as read from the device.
func (c *SNES) update(b []byte) {

	//
	// Packets contain four bytes. The first two are D-pad axis (horizontal and
	// vertical), but only have three possible values (0, 128, 255), because it's
	// just digital. The following two bytes are bitmasks of the other buttons.
	//

	c.Left = b[0] < 128
	c.Right = b[0] > 128
	c.Up = b[1] < 128
	c.Down = b[1] > 128
	c.Y = b[2]&1 != 0
	c.B = b[2]&2 != 0
	c.Select = b[2]&4 != 0
	c.Start = b[2]&8 != 0
	c.X = b[3]&1 != 0
	c.A = b[3]&2 != 0
	c.L = b[3]&4 != 0
	c.R = b[3]&8 != 0
}
