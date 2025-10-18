package Hw2

import "fmt"

/**
 * Является ли слово палиндромом?
 *
 * Напишите функцию, которая принимает
 * на вход строку и возвращает true, если
 * она является палиндромом*.
 *
 * В противном случае верните false.
 * *слово или текст, одинаково читающиеся
 * в обоих направлениях
 */

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func Problem6() {
	var s string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&s)

	if isPalindrome(s) {
		fmt.Println("String IS palindrome")
	} else {
		fmt.Println("String IS NOT palindrome")
	}
}
