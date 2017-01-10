package w28

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TheGreatXOR() {
	scanner := bufio.NewScanner(os.Stdin)
	// get the number of queries q
	var q int
	if scanner.Scan() {
		q, _ = strconv.Atoi(scanner.Text())
	}
	// loop through all the queries q
	for i := 0; i < q; i++ {

		var answer int
		if scanner.Scan() {
			input := scanner.Text()
			x, _ := strconv.Atoi(input)
			for j := 0; j < x; j++ {
				// check XOR of j and x is > x; If true increment answer.
				bitwiseCalc := j ^ x
				if bitwiseCalc > x {
					answer++
				}
			}
		}

		fmt.Println(answer)
	}
}
