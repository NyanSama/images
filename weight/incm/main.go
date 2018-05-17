package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/weight/weightconv"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "incm: %v\n", err)
				os.Exit(1)
			}
			inch := weightconv.Inch(t)
			mili := weightconv.Mili(t)

			fmt.Printf("%s = %s, %s = %s\n",
				inch, weightconv.IToM(inch), mili, weightconv.MToI(mili))
		}
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "incm: %v\n", err)
			os.Exit(1)
		}
		inch := weightconv.Inch(t)
		mili := weightconv.Mili(t)

		fmt.Printf("%s = %s, %s = %s\n",
			inch, weightconv.IToM(inch), mili, weightconv.MToI(mili))
	}
}
