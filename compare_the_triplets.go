package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	b, _ := reader.ReadString('\n')
	fmt.Println(compare(strToSlice(a), strToSlice(b)))
}

// convert string to slice
func strToSlice(s string) []int {
	strSlice := strings.Split(strings.TrimSuffix(s, "\n"), " ")
	retSlice := []int{}
	for _, val := range strSlice {
		digit, _ := strconv.Atoi(val)
		retSlice = append(retSlice, digit)
	}
	return retSlice
}

func compare(sl1, sl2 []int) string {
	resA, resB := 0, 0
	for i, _ := range sl1 {
		if sl1[i] > sl2[i] {
			resA++
		} else if sl1[i] < sl2[i] {
			resB++
		}
	}
	return strconv.Itoa(resA) + " " + strconv.Itoa(resB)
}
