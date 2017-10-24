package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	numSlice, _ := reader.ReadString('\n')
	textSlice := strings.Split(text, "\n")
	fmt.Println(arraySum(textSlice[0], numSlice))
}

func arraySum(size string, arr string) int {
	counter := 0
	noSepSlice := strings.Split(arr, " ")
	length, _:= strconv.Atoi(size)
	for i := 0; i < length; i++ {
		digit, _ := strconv.Atoi(noSepSlice[i])
		counter += digit
	}
	return counter
}