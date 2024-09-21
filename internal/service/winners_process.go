package service

import (
	"fmt"
	"sort"
	"strings"

	"alpaca.com/hungarian-lottery-checker/internal/domain"
	"alpaca.com/hungarian-lottery-checker/internal/util"
)

var lotteryPickNumbers [5]int

var winnersReport [4]int // following the order according the matchers numbers (2 > 3 > 4 > 5)

func ScanLotteryPickNumbers() {
	_, err := fmt.Scan(&lotteryPickNumbers[0], &lotteryPickNumbers[1], &lotteryPickNumbers[2], &lotteryPickNumbers[3], &lotteryPickNumbers[4])
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func WinnersProcess() {
	tuplesSize := []int{2, 3, 4, 5}
	ocorrenciesMap := *domain.GetCombinationsOcurrrenciesInstance()
	computed := *domain.GetWinnersComputedInstance()
	resetWinnersReport()

	for _, size := range tuplesSize {
		possibleCombinations := util.Combinations(lotteryPickNumbers[:], size)
		for _, comb := range possibleCombinations {
			sort.Ints(comb)
			strArr := util.ConvertIntToStringArray(comb)
			combinationKey := strings.Join(strArr, "")

			if ocorrenciesMap[combinationKey] != nil {
				for num_player := range ocorrenciesMap[combinationKey] {

					if computed[int32(num_player)] == nil {
						computed[int32(num_player)] = make([]bool, 4)
					}

					if !computed[int32(num_player)][size-2] {
						computed[int32(num_player)][size-2] = true
						winnersReport[size-2]++
					}
				}
			}
		}
	}
}

func DisplayWinnersReport() {
	fmt.Println("| Numbers matching | Winners")
	for pos, num := range winnersReport {
		fmt.Printf("|         %d        | %d \n", pos+2, num)
	}
}

func resetWinnersReport() {
	for i := 0; i < len(winnersReport); i++ {
		winnersReport[i] = 0
	}
}
