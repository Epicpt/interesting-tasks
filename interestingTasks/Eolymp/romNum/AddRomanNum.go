package main

import (
	"fmt"
	"sort"
	"strings"
)

func latNumberInRomNumb(inputNumber int) string {
	var romNumberSubber int
	var romNumbersMap = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := make([]int, len(romNumbersMap))
	i := 0
	for k := range romNumbersMap {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i int, j int) bool {
		return keys[i] > keys[j]
	})
	var answerSlice []string
	for _, value := range keys {
		for {
			romNumberSubber = inputNumber - value
			if romNumberSubber >= 0 {
				answerSlice = append(answerSlice, romNumbersMap[value])
				inputNumber = romNumberSubber
			} else {
				break
			}
		}

	}
	finalAnswer := strings.Join(answerSlice, "")
	return finalAnswer
}

func RomNumberInLatNumb(list []string) []int {

	var latNumb []int
	var romNumbersMap = map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	for i := 0; i < 2; i++ {
		var latNumbSum int
		check := 0
		for _, rimNumb := range list[i] {

			if val, ok := romNumbersMap[string(rimNumb)]; ok {

				if val > check {
					latNumbSum += val - check*2
					check = val
				} else {
					latNumbSum += val
					check = val
				}

			}
		}
		latNumb = append(latNumb, latNumbSum)
	}

	return latNumb
}

func main() {
	var a string
	fmt.Scanf("%s", &a)
	sliceString := strings.Split(a, "+")
	sliceNumbers := RomNumberInLatNumb(sliceString)
	inputNumber := sliceNumbers[0] + sliceNumbers[1]
	fmt.Println(latNumberInRomNumb(inputNumber))
}
