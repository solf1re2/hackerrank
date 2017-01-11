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
			var x int64
			input := scanner.Text()
			// x, err = strconv.Atoi(input)
			x, _ = strconv.ParseInt(input, 10, 64)
			for j := 0; int64(j) < x; j++ {
				// check XOR of j and x is > x; If true increment answer.
				bitwiseCalc := int64(j) ^ x
				if bitwiseCalc > x {
					answer++
				}
			}
		}

		fmt.Println(answer)
	}
}
