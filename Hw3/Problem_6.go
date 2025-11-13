package Hw3

import (
	"fmt"
	"sort"
)

/*
 * Массив анаграмм
 *
 * Дан массив строк strs. Сгруппируйте анаграммы вместе.
 * Вы можете вернуть ответ в любом порядке.
 *
 * Анаграмма — это слово или фраза, полученная путем
 * перестановки букв другого слова или фразы, обычно
 * с использованием всех исходных букв ровно один раз.
 */

func GroupAnagrams(words []string) [][]string {
	unique_words := make(map[string][]string)

	for _, word := range words {
		sortedWord := []rune(word)
		sort.Slice(sortedWord, func(i, j int) bool { return sortedWord[i] < sortedWord[j] })
		unique_words[string(sortedWord)] = append(unique_words[string(sortedWord)], word)
	}

	res := make([][]string, 0)
	for _, anagrams := range unique_words {
		res = append(res, anagrams)
	}

	return res
}

func Problem6() {
	var n int
	fmt.Print("Enter count of words: ")
	_, _ = fmt.Scan(&n)

	fmt.Print("Enter words: ")
	words := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&words[i])
	}

	fmt.Print("Grouped anagrams: ", GroupAnagrams(words))
}
