package node

import (
	"math/big"
)

var Reducers = make(map[string]func(interface{}, ...interface{}) interface{} )

func LoadReducers() {
	Reducers["ADD"] = ADD
	Reducers["MUL"] = MUL
	Reducers["DIV"] = DIV
	Reducers["BYTE32"] = BYTE32
	Reducers["MEAN"] = MEAN
}

func ADD(value interface{}, args ...interface{}) interface{} {
	total := value.(int64)
	bigTotal := big.NewInt(total)

 	for i := 0; i < len(args); i++ {
		 num := args[i].(int64)
		 bigNum := big.NewInt(num)

		bigTotal = bigTotal.Add(bigTotal, bigNum)
	}

	return bigTotal
}

func MUL(value interface{}, args ...interface{}) interface{} {
	return  nil
}

func DIV(value interface{}, args ...interface{}) interface{} {
	return  nil
}

func BYTE32(value interface{}, args ...interface{}) interface{} {
	return  nil
}

func MEAN(value interface{}, items ...interface{}) interface{} {
	return  nil
}

func MEDIAN(items []interface{}) interface{} {
	return  nil
}

func MODE(items []interface{}) {

}

func BASE64DECODE(value interface{}) {

}

func BASE64ENCODE(value interface{}) {

}

func UPPERCASE(value string) {

}

func LOWERCASE(value string) {

}

// HEXDECODE Given the input 0x12345678, the task will return [0x12, 0x34, 0x56, 0x78]
func HEXDECODE(value string) {

}

// HEXENCODE Given the input string "xyz", the task will return "0x78797a"
func HEXENCODE(value string) {

}

func ETHABIDECODELOG(data byte, arg ...interface{}){

}