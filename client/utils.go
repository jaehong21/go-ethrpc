package client

import (
	"math/big"
)

func ParseHexToBigInt(hex string, decimals int) *big.Int {
	num := new(big.Int)
	num.SetString(hex[2:], 16)
	if decimals == 0 {
		return num
	}

	decimal, pow := big.NewInt(10), big.NewInt(int64(decimals))
	decimal.Exp(decimal, pow, nil)

	result := new(big.Int)
	result.Div(num, decimal)
	return result
}

func ParseHexToBigFloat(hex string, decimals int) *big.Float {
	num := ParseHexToBigInt(hex, 0)
	float := ParseIntToFloat(num)
	if decimals == 0 {
		return float
	}

	decimal := new(big.Float)
	decimal.SetInt(Exp(10, decimals))

	result := new(big.Float)
	result.Quo(float, decimal)
	return result
}

func ParseIntToFloat(num *big.Int) *big.Float {
	float := new(big.Float)
	float.SetInt(num)
	return float
}

func Exp(base int, exp int) *big.Int {
	num, pow := big.NewInt(int64(base)), big.NewInt(int64(exp))
	num.Exp(num, pow, nil)
	return num
}
