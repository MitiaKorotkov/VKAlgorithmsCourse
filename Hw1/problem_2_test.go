package Hw1

import "testing"

func TestTask2Case1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	ReverseArray(arr)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask2Case2(t *testing.T) {
	arr := []int{}
	expected := []int{}
	ReverseArray(arr)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask2Case3(t *testing.T) {
	arr := []int{1}
	expected := []int{1}
	ReverseArray(arr)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}
