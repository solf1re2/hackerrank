package w28

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BoatTrips() {
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
