package weightconv

func IToM(i Inch) Mili { return Mili(i * 0.3048) }
func MToI(m Mili) Inch { return Inch(m * 3.2808399) }
