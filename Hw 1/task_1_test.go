package main

import "testing"

func Compare(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestTask1Case1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	actual := TwoSum(nums, target)

	if Compare(expected, actual) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, actual)
	}
}

func TestTask1Case2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	expected := []int{0, 3}
	actual := TwoSum(nums, target)

	if Compare(expected, actual) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, actual)
	}
}

func TestTask1Case3(t *testing.T) {
	nums := []int{1, 1, 3, 3}
	target := 6
	expected := []int{2, 3}
	actual := TwoSum(nums, target)

	if Compare(expected, actual) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, actual)
	}
}

func TestTask1Case4(t *testing.T) {
	nums := []int{1, 1}
	target := 2
	expected := []int{0, 1}
	actual := TwoSum(nums, target)

	if Compare(expected, actual) == false {
		t.Errorf("Result incorrect. Expected: %v, actual: %v", expected, actual)
	}
}
