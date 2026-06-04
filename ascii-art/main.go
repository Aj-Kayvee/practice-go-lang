// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Println("Invalid Input")
// 		return
// 	}

// 	text := os.Args[1]

// 	if len(text) == 0 {
// 		return
// 	}

// 	inputText := strings.Split(text, "\\n")

// 	file, err := os.Open("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error Opening File")
// 		return
// 	}

// 	var banner []string

// 	newFile := bufio.NewScanner(file)
// 	for newFile.Scan() {
// 		line := newFile.Text()
// 		banner = append(banner, line)
// 	}

// 	for index, val := range inputText {
// 		if val == "" {
// 			fmt.Println()
// 			continue
// 		}

// 		for row := 0; row < 8; row++ {
// 			for col := 0; col < len(inputText[index]); col++ {
// 				start := int(inputText[index][col]-32)*9 + 1
// 				result := banner[start+row]
// 				fmt.Print(result)
// 			}
// 			fmt.Println()
// 		}
// 	}

// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid Input")
	}

	text := os.Args[1]

	inputText := strings.Split(text, "\\n")

	if len(text) == 0 {
		return
	}

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	var banner []string

	inputFile := bufio.NewScanner(file)
	for inputFile.Scan() {
		banner = append(banner, inputFile.Text())
	}

	for index, value := range inputText {
		if value == "" {
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			for col := 0; col < len(inputText[index]); col++ {
				start := int(inputText[index][col]-32) * 9+1
				result := banner[start + row]
				fmt.Print(result)
			}
			fmt.Println()
		}
	}
}