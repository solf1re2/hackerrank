package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
}

func theGreatXOR() {
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

func boatTrips() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := "Yes"
	var tourGroups int
	var boatCapacity int
	var boatCount int
	var maxCap int
	for scanner.Scan() {
		input := scanner.Text()
		inputArray := strings.Split(input, " ")
		if len(inputArray) == 3 {
			tourGroups, _ = strconv.Atoi(inputArray[0])
			boatCapacity, _ = strconv.Atoi(inputArray[1])
			boatCount, _ = strconv.Atoi(inputArray[2])
			maxCap = boatCount * boatCapacity
		} else if len(inputArray) == tourGroups {
			for i := 0; i < tourGroups; i++ {
				passengers, _ := strconv.Atoi(inputArray[i])
				if maxCap < passengers {
					ans = "No"
				}
			}

		}
	}
	fmt.Println(ans)
}
