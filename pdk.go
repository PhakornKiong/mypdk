package pdk

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

func Input(v interface{}) error {
	input := pdk.Input()
	return json.Unmarshal(input, &v)
}

func MemoryToBytes(mem pdk.Memory) []byte {
	buf := make([]byte, mem.Length())
	mem.Load(buf)
	return buf
}

func MemoryToStruct(mem pdk.Memory, v interface{}) error {
	bytes := MemoryToBytes(mem)
	return json.Unmarshal(bytes, &v)
}

func BytesToMemory(b []byte) (mem pdk.Memory) {
	mem = pdk.AllocateBytes(b)
	return
}

func StructToMemory(v interface{}) (mem pdk.Memory, err error) {
	b, err := json.Marshal(&v)

	if err != nil {
		return
	}

	mem = BytesToMemory(b)
	return
}

func GetConfig(key string) (string, bool) {
	return pdk.GetConfig(key)
}
