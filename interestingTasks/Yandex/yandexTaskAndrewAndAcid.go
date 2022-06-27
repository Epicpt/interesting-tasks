package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func doliv(st []int) int {
	var max int
	for _, storage := range st {
		if storage > max {
			max = storage
		}
		if st[0] > storage {
			return -1
		}
	}
	if st[0] == max {
		return -1
	} else {
		return max - st[0]
	}
}

func createIntSlice(nums []string) []int {
	var res []int
	for _, v := range nums {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, i)
	}
	return res
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	err := scanner.Scan()
	if err != true {
		log.Fatal()
	}
	scanner.Scan()

	parts := strings.Split(scanner.Text(), " ")
	storages := createIntSlice(parts)

	value := doliv(storages)
	fmt.Println(value)

}
