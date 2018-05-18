package weightconv

import (
	"fmt"
)

//Inch : type for feet
type Inch float64

//Mili : type for meter
type Mili float64

func (c Inch) String() string { return fmt.Sprintf("%g英尺", c) }
func (m Mili) String() string { return fmt.Sprintf("%g米", m) }
