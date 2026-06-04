package main

import (
	"fmt"

	// "slices"
	// "slices"
	"strings"
	// "strconv"
	// "strings"
)

func main() {
	PrintChars("Hello")

	c := CountChars("golang")
	fmt.Println(c)

	fmt.Println(ToUpperCase("hello"))
	fmt.Println(ToLowerCase("GOLANG"))
	fmt.Println(Reverse("hello"))

	fmt.Println(IsPalindrome("hello"))
	fmt.Println(SwapCases("GoLang!"))
	fmt.Println(CountVowels("vowel"))
	fmt.Println(RemoveSpaces("go is fun"))
	fmt.Println(FirstIndex("golang", 'a'))
	fmt.Println(CountWords("go is very fast"))
	fmt.Println(Join("go", "lang"))
	fmt.Println(RelaceChars("banana", 'a', 'o'))
	Duplicates("Programming")
	fmt.Println(IsAnagram("silent", "listen"))
	fmt.Println(IsAnagram("hello", "world"))
}

func PrintChars(chars string) {
	ch := "Hello"

	for _, r := range ch {
		fmt.Println(string(r))
	}
}

func CountChars(ch string) int {

	// count := 0
	// for i := 0; i < len(ch); i++ {
	// 	count ++
	// }
	return len(ch)
}

func ToUpperCase(str string) string {

	words := strings.ToUpper(str)
	return words
}

func ToLowerCase(str string) string {
	return strings.ToLower(str)
}

func Reverse(str string) string {

	// result := ""

	var newStr strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		// result += string(str[i])
		newStr.WriteByte(str[i])
	}
	return newStr.String()
}

func IsPalindrome(str string) bool {

	// word := ""
	var isPal strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		isPal.WriteByte(str[i])
	}
	if str == isPal.String() {
		return true
	}
	return false
}

func SwapCases(str string) string {
	word := ""

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			word += string(char - 32)
		} else if char >= 'A' && char <= 'Z' {
			word += string(char + 32)
		} else {
			word += string(char)
		}

	}

	return word
}

func CountVowels(str string) int {
	count := 0

	for _, ch := range str {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}
	return count
}

func RemoveSpaces(s string) string {
	newStr := strings.Split(s, " ")

	return strings.Join(newStr, "")
}

func FirstIndex(s string, ch rune) int {
	for index, newCh := range s {
		if newCh == ch {
			return index
		}

	}
	return 0
}

func CountWords(str string) int {

	countIndex := strings.Fields(str)
	count := 0

	for range countIndex {
		count++
	}
	return count
}

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range a {
		counts[char]++
	}
	for _, char := range b {
		counts[char]--

		if counts[char] < 0 {
			return false
		}
	}
	return true
}

func Join(a, b string) string {

	return a + b
}

func RelaceChars(str string, old, new rune) string {
	var newChar []string

	for _, char := range str {
		if char == old {
			char = new
		}
		newChar = append(newChar, string(char))
	}

	return strings.Join(newChar, "")
	// var newChar strings.Builder

	// for _, char := range str {
	// 	if char == old {
	// 		newChar.WriteRune(new)
	// 	} else {
	// 		if char != old {
	// 			newChar.WriteRune(char)
	// 		}
	// 	}
	// }
	// return newChar.String()
}

func Duplicates(str string) {
	counts := make(map[rune]int)

	for _, char := range str {
		counts[char]++
	}

	for char, count := range counts {
		if count > 1 {
			fmt.Println(string(char))
		}
	}
}
