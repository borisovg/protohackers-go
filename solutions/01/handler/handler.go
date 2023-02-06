package handler

import (
	"errors"
	"math/big"
	"solutions/01/protocol"
)

func GetNumber(line string) (*big.Int, error) {
	int := new(big.Int)

	req, err := protocol.ParseRequest(line)
	if err != nil {
		return int, err
	}

	float := new(big.Float)
	err = float.UnmarshalText(req.Number)
	if err != nil {
		return int, errors.New("bad number")
	}

	float.Int(int)
	return int, err
}
