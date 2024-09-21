package main

import (
	"fmt"
	"os"
	"strconv"

	"alpaca.com/hungarian-lottery-checker/internal/service"
)

func main() {

	args := os.Args

	path := args[1]

	totalLines := 0

	if len(args) >= 3 {
		lines, _ := strconv.Atoi(args[2])
		totalLines = lines
	}

	service.PreprocessFile(path, totalLines)
	fmt.Println("\nREADY")

	for {
		service.ScanLotteryPickNumbers()
		service.WinnersProcess()
		service.DisplayWinnersReport()
	}
}
