package main

import (
	"fmt"
)

type Env struct{}

func (e *Env) GetCalldata() []byte {
	return []byte("switza")
}

func (e *Env) SetReturnData(data []byte) {
	fmt.Println("set", string(data))
}

func (e *Env) GetAskCount() int64 {
	return 10000
}

func (e *Env) GetMinCount() int64 {
	return 20000
}

func (e *Env) GetPrepareTime() int64 {
	return 1623403856
}

func (e *Env) GetExecuteTime() int64 {
	return 1623403868
}

func (e *Env) GetAnsCount() int64 {
	return 30000
}

func (e *Env) AskExternalData(eid int64, did int64, data []byte) {
	fmt.Println("asked", eid, did, string(data))
}

func (e *Env) GetExternalDataStatus(eid int64, vid int64) int64 {
	return 42
}

func (e *Env) GetExternalData(eid int64, vid int64) []byte {
	return []byte("switez")
}

// func Wat2Wasm(fileName string, spanSize int) error {
// 	code, _ := ioutil.ReadFile(fmt.Sprintf("./wasm/%s.wat", fileName))
// 	wasm, err := api.Wat2Wasm(code, spanSize)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("wasm:", wasm)
// 	f, err := os.Create(fmt.Sprintf("./wasm/%s.wasm", fileName))
// 	defer f.Close()
// 	n, err := f.Write(wasm)
// 	fmt.Printf("wrote %d bytes\n", n)
// 	if err != nil {
// 		return err
// 	}
// 	f.Sync()

// 	return nil
// }

func main() {
	fmt.Println("Hello, World!")
}
