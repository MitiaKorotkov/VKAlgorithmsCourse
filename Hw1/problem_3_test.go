package Hw1

import "testing"

func TestTask3Case1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 3
	expected := []int{4, 5, 6, 1, 2, 3}
	Rotate(arr, k)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask3Case2(t *testing.T) {
	arr := []int{1}
	k := 4
	expected := []int{1}
	Rotate(arr, k)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask3Case3(t *testing.T) {
	arr := []int{}
	k := 1
	expected := []int{}
	Rotate(arr, k)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask3Case4(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	k := 10
	expected := []int{3, 4, 1, 2}
	Rotate(arr, k)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}

func TestTask3Case5(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	k := 0
	expected := []int{1, 2, 3, 4}
	Rotate(arr, k)

	if Compare(expected, arr) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, arr)
	}
}
