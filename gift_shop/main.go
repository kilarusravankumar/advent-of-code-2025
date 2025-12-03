package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	invalidIDTotal := 0
	var inputText string
	for scanner.Scan() {
		inputText += scanner.Text()
	}
	ranges := strings.Split(inputText, ",")
	for i := range ranges {
		lowHighArr := strings.Split(ranges[i], "-")
		low, err := strconv.Atoi(lowHighArr[0])
		if err != nil {
			panic(err)
		}
		high, err := strconv.Atoi(lowHighArr[1])
		if err != nil {
			panic(err)
		}

		fmt.Printf("looking for invalidIDs between low %d : high %d \n", low, high)
		invalidIDTotal += invalidIDsInRange(low, high)
	}
	fmt.Printf("invalidIDTotal is -----------------------> %d \n", invalidIDTotal)
}

func invalidIDsInRange(low, high int) int {
	invalidIDSum := 0
	for i := low; i < high+1; i++ {
		if isInvalidID(i) {
			fmt.Printf("is it invalid; yes it is  --> %d \n", i)
			invalidIDSum += i
		}
	}
	return invalidIDSum
}

func isInvalidID(id int) bool {
	idStr := strconv.Itoa(id)

	if len(idStr) <= 1 {
		return false
	}

	doubled := idStr + idStr

	return strings.Contains(doubled[1:len(doubled)-1], idStr)
}
