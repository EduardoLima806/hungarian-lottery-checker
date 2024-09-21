package domain

import (
	"sync"
)

var lock = &sync.Mutex{}

type CombinationsOcurrrencies map[string][]int32
type WinnersComputed map[int32][]bool

var ocorrenciesInstance CombinationsOcurrrencies // Singleton instance

func GetCombinationsOcurrrenciesInstance() *CombinationsOcurrrencies {
	if ocorrenciesInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if ocorrenciesInstance == nil {
			ocorrenciesInstance = CombinationsOcurrrencies{}
		}
	}
	return &ocorrenciesInstance
}

func GetWinnersComputedInstance() *WinnersComputed {
	return &WinnersComputed{}
}

func (p CombinationsOcurrrencies) IncrementOccorrence(combination string, size int, num_player int32) {

	ocorrenciesInstance[combination] = append(ocorrenciesInstance[combination], num_player)
}
