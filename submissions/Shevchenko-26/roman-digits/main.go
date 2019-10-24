package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const digitLength, digitWidth = 7, 5

func main() {
	if len(os.Args) != 2 && len(os.Args) != 4 {
		log.Fatal("Need only one input number")
	}
	symbolFlag := flag.Bool("symbol", false, "set custom symbol for the art")
	flag.Parse()

	var symbol, number string
	if *symbolFlag {
		number, symbol = os.Args[3], os.Args[2]
	} else {
		number = os.Args[1]
	}
	if _, err := strconv.ParseUint(number, 10, 64); err != nil {
		log.Fatalf("Input argument is integer non-convertible:  %s\n", err)
	}

	digs := prepareDigits(symbol)

	asteriskBorder := strings.Repeat("*", len(number)*(digitWidth+1)-1)

	numberArr := strings.Split(number, "")
	var numbers []int
	for _, num := range numberArr {
		n, _ := strconv.Atoi(num)
		numbers = append(numbers, n)
	}

	fmt.Println(asteriskBorder)
	for i := 0; i < digitLength; i++ {
		for _, num := range numbers {
			fmt.Printf("%s ", digs[num][i])
		}
		fmt.Println()
	}
	fmt.Println(asteriskBorder)

}

func prepareDigits(symbol string) [10][]string {
	artNumbers := strings.Split(digits, "\n\n")
	numbers := [10][]string{}
	for i, digit := range artNumbers {
		if symbol != "" {
			digit = strings.ReplaceAll(digit, strconv.Itoa(i), symbol)
		}
		digitRows := strings.Split(strings.Trim(digit, "\n"), "\n")
		for _, digitString := range digitRows {
			digitString = fmt.Sprintf("%-5s", digitString)
			numbers[i] = append(numbers[i], digitString)
		}
	}
	return numbers
}
