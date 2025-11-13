package Hw3

import "fmt"

/*
 * Найти разницу между двух строк
 *
 * На вход функции подается две строки: a и b. Строка b образована
 * из строки a путем перемешивания и добавления одной буквы.
 * Необходимо вернуть эту букву
 */

func FindExtraLetter(a string, b string) rune {
	bCounter := make(map[rune]int)

	for _, ch := range b {
		bCounter[ch]++
	}

	for _, ch := range a {
		bCounter[ch]--
	}

	for ch, count := range bCounter {
		if count != 0 {
			return ch
		}
	}

	return ' '
}

func Problem4() {
	var a string
	var b string

	fmt.Print("Enter string a: ")
	_, _ = fmt.Scan(&a)
	fmt.Print("Enter string b: ")
	_, _ = fmt.Scan(&b)

	fmt.Print("Extra letter is: ", string(FindExtraLetter(a, b)))
}
