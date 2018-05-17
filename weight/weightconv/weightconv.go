package weightconv

import (
	"fmt"
)

type Inch float64
type Mili float64

func (c Inch) String() string { return fmt.Sprintf("%g英尺", c) }
func (m Mili) String() string { return fmt.Sprintf("%g米", m) }
