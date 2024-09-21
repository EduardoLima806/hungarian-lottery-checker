package util

func Combinations(arr []int, k int) [][]int {
	var result [][]int
	var comb []int
	generateCombinations(arr, k, 0, comb, &result)
	return result
}

func generateCombinations(arr []int, k, start int, comb []int, result *[][]int) {
	if len(comb) == k {
		combCopy := make([]int, k)
		copy(combCopy, comb)
		*result = append(*result, combCopy)
		return
	}

	for i := start; i < len(arr); i++ {
		comb = append(comb, arr[i])
		generateCombinations(arr, k, i+1, comb, result)
		comb = comb[:len(comb)-1]
	}
}
