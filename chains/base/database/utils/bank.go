package utils

import "github.com/forbole/bdjuno/v3/chains/base/types"

const (
	MaxPostgreSQLParams = 65535
)

func SplitAccounts(accounts []types.Account, paramsNumber int) [][]types.Account {
	maxBalancesPerSlice := MaxPostgreSQLParams / paramsNumber
	slices := make([][]types.Account, len(accounts)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accounts {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}