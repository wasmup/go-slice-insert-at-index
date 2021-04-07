package main

import "fmt"

func main() {
	// My method (@wasmup):
	{
		// objective put "C" at index 2
		const index = 2
		// index:            0    1    2    3    4
		a := []string{"A", "B", "D", "E", "F"}

		// 1+2) add an element to the end of the slice and copy
		a = append(a[:index+1], a[index:]...) // Step 1+2

		// 3) set "C" at index 2
		a[index] = "C" // Step 3
		fmt.Println(a) // [A B C D E F]
	}

	// @icza: https://stackoverflow.com/a/46130603/8208215
	{
		// objective put "C" at index 2
		const index = 2
		// index:            0    1    2    3    4
		a := []string{"A", "B", "D", "E", "F"}

		// 1) add the last element to the end of the slice
		last := len(a) - 1
		a = append(a, a[last]) // Step 1

		// 2) copy rest not last
		copy(a[index+1:], a[index:last]) // Step 2

		// 3) set "C" at index 2
		a[index] = "C" // Step 3
		fmt.Println(a) // [A B C D E F]
	}

	// The Book: https://www.practical-go-lessons.com/chap-21-slices
	{
		// objective put "C" at index 2
		const index = 2
		// index:            0    1    2    3    4
		letters := []string{"A", "B", "D", "E", "F"}

		// 1) add an empty element to the end of the slice
		letters = append(letters, "")

		// 2) copy all letters[i:] to letters[i+1:]
		copy(letters[index+1:], letters[index:])

		// 3) set "C" at index 2
		letters[index] = "C"
		fmt.Println(letters) // [A B C D E F]
	}

}
