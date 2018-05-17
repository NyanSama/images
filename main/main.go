package main

import (
	"fmt"
)

// func main() {
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		x, err := strconv.ParseUint(input.Text(), 0, 64)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "pop: %v\n", err)
// 			os.Exit(1)
// 		}
// 		fmt.Println(popcount.PopCount(x))
// 	}
// }

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Printf("%v\n", pc)
}
