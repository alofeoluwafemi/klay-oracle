package node

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

var Reducers = make(map[string]func(interface{}, ...interface{}) (interface{}, error))

func LoadReducers() {
	Reducers["PARSE"] = PARSE
	Reducers["MUL"] = MUL
	Reducers["ADD"] = ADD
	Reducers["SUB"] = SUB
	Reducers["DIV"] = DIV
}

func PARSE(jsonInterface interface{}, args ...interface{}) (interface{}, error) {

	var result map[string]interface{}

	jsonString := jsonInterface.(string)

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		log.Fatalf("Unable to decode server response due to %v", err)
	}

	for i := 0; i < len(args); i++ {
		index := args[i].(string)
		fmt.Println("Key: ", index)

		if i == (len(args) - 1) {
			return result[index], nil
		}

		jsonByte, err := json.Marshal(result[index])

		err = json.Unmarshal(jsonByte, &result)
		if err != nil {
			log.Fatalf("Unable to decode server response due to %v", err)
		}

	}

	return nil, fmt.Errorf("%v", "Decode error")
}

func MUL(value interface{}, args ...interface{}) (interface{}, error) {
	stringVal := value.(string)

	total, _ := strconv.ParseFloat(stringVal, 64)

	bigTotal := big.NewFloat(total)

	for i := 0; i < len(args); i++ {
		num := args[i].(float64)
		bigNum := big.NewFloat(num)

		bigTotal = bigTotal.Mul(bigTotal, bigNum)
	}

	bigTotalStr := fmt.Sprintf("%.0f", bigTotal)
	return bigTotalStr, nil
}

func SUB(value interface{}, args ...interface{}) (interface{}, error) {
	stringVal := value.(string)

	total, _ := strconv.ParseFloat(stringVal, 64)

	bigTotal := big.NewFloat(total)

	for i := 0; i < len(args); i++ {
		num := args[i].(float64)
		bigNum := big.NewFloat(num)

		bigTotal = bigTotal.Sub(bigTotal, bigNum)
	}

	bigTotalStr := fmt.Sprintf("%.0f", bigTotal)
	return bigTotalStr, nil
}

func ADD(value interface{}, args ...interface{}) (interface{}, error) {
	stringVal := value.(string)

	total, _ := strconv.ParseFloat(stringVal, 64)

	bigTotal := big.NewFloat(total)

	for i := 0; i < len(args); i++ {
		num := args[i].(float64)
		bigNum := big.NewFloat(num)

		bigTotal = bigTotal.Add(bigTotal, bigNum)
	}

	bigTotalStr := fmt.Sprintf("%.0f", bigTotal)
	return bigTotalStr, nil
}

func DIV(value interface{}, args ...interface{}) (interface{}, error) {
	stringVal := value.(string)

	total, _ := strconv.ParseFloat(stringVal, 64)

	bigTotal := big.NewFloat(total)

	bigTotalF64, _ := bigTotal.Float64()

	for i := 0; i < len(args); i++ {
		num := args[i].(float64)
		bigNum := big.NewFloat(num)
		bigNumF64, _ := bigNum.Float64()

		bigTotalF64 = bigTotalF64 / bigNumF64
	}

	bigTotalStr := fmt.Sprintf("%.0f", bigTotal)
	return bigTotalStr, nil
}