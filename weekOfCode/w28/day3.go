package w28

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func LuckyNumberEight() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	var nDigitIntegerString string
	var combinationsArray []int
	var answerCount int
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		log.Printf("n = %v", n)
	}
	if scanner.Scan() {
		nDigitIntegerString = scanner.Text()
		log.Printf("nDigitIntegerString = %v", nDigitIntegerString)
	}
	// There are n different lengths for combinations.
	// for each combination length, shift it along the nDigitIntegerString for n-x => x+1 shifts for that length?
	for i := 0; i < n; i++ {
		combinationLength := n - i // 3,2,1 etc
		log.Printf("combinationLength = %v", combinationLength)
		for j := 0; j <= i+1; j++ {
			log.Printf("n-combinationLength = %v", n-combinationLength)
			subString := nDigitIntegerString[j : j+combinationLength]
			log.Printf("At i=%v and j=%v, substring=%s", i, j, subString)
			subStringValue, _ := strconv.Atoi(subString)
			if subStringValue > 0 {
				combinationsArray = append(combinationsArray, subStringValue)
				log.Println(combinationsArray)
			}
		}
	}
	for i := 0; i < len(combinationsArray); i++ {
		combinationValue := combinationsArray[i]
		floatComboValue := float64(combinationValue)
		if math.Mod(floatComboValue, float64(8)) == float64(0) {
			answerCount++
		}
	}
	fmt.Println(answerCount)
}
