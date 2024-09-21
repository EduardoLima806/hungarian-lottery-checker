package service

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"alpaca.com/hungarian-lottery-checker/internal/domain"
	"alpaca.com/hungarian-lottery-checker/internal/util"
)

type Players []domain.InputPlayer

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func PreprocessFile(path string, totalLines int) {
	// Open file (get only the reference without load into memory yet)
	file, err := os.Open(path)
	check_error(err)
	defer file.Close()

	r := bufio.NewReader(file)

	combinations := domain.GetCombinationsOcurrrenciesInstance()

	fmt.Println("Pre-processing file...")

	num_line := 1
	var rProgress float64

	for {
		line, _, err := r.ReadLine()

		generateAndIncrementCombinations(line, int32(num_line), combinations)

		if totalLines != 0 {
			rProgress = (float64(num_line) / float64(totalLines)) * 100

			fmt.Printf("\rProgress: %.2f%%", math.Round(rProgress*100)/100)
		}

		num_line++

		if err != nil {
			break
		}
	}
}

func generateAndIncrementCombinations(line []byte, num_line int32, combinations *domain.CombinationsOcurrrencies) {

	if len(line) > 0 {
		numbersStr := strings.Split(string(line), " ")

		if len(numbersStr) == 5 {
			numbers := util.ConvertToIntArray(numbersStr)
			tuplesSize := []int{2, 3, 4, 5}

			for _, size := range tuplesSize {
				possiblesComb := util.Combinations(numbers, size)
				for _, comb := range possiblesComb {
					sort.Ints(comb)
					strArr := util.ConvertIntToStringArray(comb)
					combinationKey := strings.Join(strArr, "")
					combinations.IncrementOccorrence(combinationKey, size, num_line)
				}
			}
		}
	}
}
