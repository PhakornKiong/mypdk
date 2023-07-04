package pdk

import (
	"github.com/extism/go-pdk"
)

func AllocateString(data string) pdk.Memory {
	return pdk.AllocateBytes([]byte(data))
}
