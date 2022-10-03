package node

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

var Reducers = make(map[string]func(interface{}, ...interface{}) (interface{}, error) )

func LoadReducers() {
	//Reducers["ADD"] = ADD
	Reducers["MUL"] = MUL
	//Reducers["DIV"] = DIV
	//Reducers["BYTE32"] = BYTE32
	//Reducers["MEAN"] = MEAN
	Reducers["PARSE"] = PARSE
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

	return nil, fmt.Errorf("%v","Decode error")
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
	return bigTotalStr ,nil
}

//func ADD(value interface{}, args ...interface{}) interface{} {
//stringVal := value.(string)
//
//total, _ := strconv.ParseFloat(stringVal, 64)
//
//bigTotal := big.NewFloat(total)
//
//for i := 0; i < len(args); i++ {
//num := args[i].(float64)
//bigNum := big.NewFloat(num)
//
//bigTotal = bigTotal.Add(bigTotal, bigNum)
//}
//
//bigTotalStr := fmt.Sprintf("%.0f", bigTotal)
//return bigTotalStr ,nil
//}

//
//func MUL(value interface{}, args ...interface{}) interface{} {
//	return  nil
//}
//
//func DIV(value interface{}, args ...interface{}) interface{} {
//	return  nil
//}
//
//func BYTE32(value interface{}, args ...interface{}) interface{} {
//	return  nil
//}
//
//func MEAN(value interface{}, items ...interface{}) interface{} {
//	return  nil
//}
//
//func MEDIAN(items []interface{}) interface{} {
//	return  nil
//}
//
//func MODE(items []interface{}) {
//
//}
//
//func BASE64DECODE(value interface{}) {
//
//}
//
//func BASE64ENCODE(value interface{}) {
//
//}
//
//func UPPERCASE(value string) {
//
//}
//
//func LOWERCASE(value string) {
//
//}
//
//// HEXDECODE Given the input 0x12345678, the task will return [0x12, 0x34, 0x56, 0x78]
//func HEXDECODE(value string) {
//
//}
//
//// HEXENCODE Given the input string "xyz", the task will return "0x78797a"
//func HEXENCODE(value string) {
//
//}
//
//func ETHABIDECODELOG(data byte, arg ...interface{}){
//
//}