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
		fmt.Println(err)
	}
	defer file.Close()

	total_score := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		target, numbers := strings.Split(scanner.Text(), ": ")[0], strings.Split(scanner.Text(), ": ")[1]
		arr_nums := []int{}
		int_target, _ := strconv.Atoi(target)
		for _, num := range strings.Split(numbers, " ") {
			fmt.Printf("Number: %s\n", num)
			int_num, _ := strconv.Atoi(num)
			arr_nums = append(arr_nums, int_num)
		}
		if isPartOfExpression(int_target, arr_nums) {
			total_score += int_target
		}
		fmt.Printf("Target: %s, Number: %s\n", target, numbers)
	}
	fmt.Println(total_score)
}

func isPartOfExpression(target int, arr []int) bool {
	//If there is only one element in the array, return true if the target is equal to the target
	fmt.Printf("Target: %d, Array: %v\n", target, arr)

	if len(arr) == 1 {
		return target == arr[0]
	}
	//If the target is divisible by the last element in the array and the rest of the array is a part of the expression, return true
	if target%arr[len(arr)-1] == 0 && isPartOfExpression(target/arr[len(arr)-1], arr[:len(arr)-1]) {
		return true
	}
	//If the target is greater than the last element in the array and the rest of the array is a part of the expression, return true
	if target > arr[len(arr)-1] && isPartOfExpression(target-arr[len(arr)-1], arr[:len(arr)-1]) {
		return true
	}
	string_target := strconv.Itoa(target)
	fmt.Printf("String target: %s\n", string_target)
	string_last := strconv.Itoa(arr[len(arr)-1])
	fmt.Printf("String last: %s\n", string_last)
	if len(string_target) > len(string_last) && strings.HasSuffix(string_target, string_last) && isPartOfExpression(getTargetWithOutLast(string_target, string_last), arr[:len(arr)-1]) {
		return true
	}
	return false
}

func getTargetWithOutLast(target string, last string) int {
	target_wo_last, _ := strconv.Atoi(target[:len(target)-len(last)])
	return target_wo_last
}
