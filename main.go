package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func search(arr []int, value int, low int, high int) (int, error) {
	if low > high {
		return 0, errors.New("low is higher than high")
	}

	mid := low + (high-low)/2

	if value == arr[mid] {
		return mid, nil
	}

	if value > arr[mid] {
		return search(arr, value, mid+1, high)
	}

	return search(arr, value, low, mid-1)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	print("Please enter search value: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	atoi, _ := strconv.Atoi(scanner.Text())

	res, err := search(arr, atoi, 0, len(arr)-1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
