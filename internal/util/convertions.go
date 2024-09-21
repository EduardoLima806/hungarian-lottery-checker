package util

import (
	"fmt"
	"strconv"
)

func ConvertToIntArray(stringArray []string) []int {

	intArray := make([]int, len(stringArray))

	// Loop through each string and convert to int
	for i, str := range stringArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting:", str, err)
		}
		intArray[i] = num
	}

	return intArray
}

func ConvertIntToStringArray(intArr []int) []string {
	strArr := make([]string, len(intArr))

	// Step 3: Convert each integer to string
	for i, num := range intArr {
		strArr[i] = strconv.Itoa(num) // Convert int to string
	}

	return strArr
}
