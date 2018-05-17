package weightconv

//IToM function fo feet to meter
func IToM(i Inch) Mili { return Mili(i * 0.3048) }

//MToI function for meter to feet
func MToI(m Mili) Inch { return Inch(m * 3.2808399) }
