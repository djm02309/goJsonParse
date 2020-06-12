package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)
type Functions struct {
	FunctionInfos []struct {
		Argu string
		ConstVar uint
		Sym string
		Signature string
	} `json:"functions"`
}
/*
type FunctionInfo struct {
	Argu string
	ConstVar uint
	Sym string
	Signature string
}

func ReadJSON( inputString string ) *FunctionInfo {
	f := FunctionInfo{}
	//json.Unmarshal([]byte(inputString), &f)
	err:= json.Unmarshal([]byte(inputString), &f)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &f
}
*/

func testFunctions(file *os.File) Functions {
	var fun Functions
	decoder := json.NewDecoder(file)
	for {
		if err := decoder.Decode(&fun); err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}
		fmt.Println(fun)
	}
	return fun
//출처: https://beji.tistory.com/entry/Json-Data-to-Go-Struct-unmarshal-vs-decode [¡Hola, Mondo!]
}

func main(){

	//ReadJSON testcode
	//str := `{"argu": "vari", "constVar": 3, "sym": ">", "signature": "0x123541" }`
	//str23:= `{"page": 1, "fruits": ["apple", "peach"]}`
	//asd := A(str23)
	//f := ReadJSON(str)
	//fmt.Println(f.Sym)

	file, _ := os.Open("emp.json")  //파일 읽기 (파일 위치는 현재파일이랑 같은 곳)
	testFunctions(file)

}