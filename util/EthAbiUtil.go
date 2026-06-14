package util

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func BuildLocalParsedABI(minifiedJSON string) (*abi.ABI, error) {
	parsedABI, err := abi.JSON(strings.NewReader(minifiedJSON))
	if err != nil {
		return nil, err
	}
	return &parsedABI, nil
}

func FilterIndexedFields(inputs abi.Arguments) abi.Arguments {
	var indexedFields abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexedFields = append(indexedFields, arg)
		}
	}
	return indexedFields
}
