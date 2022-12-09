//go:build tinygo.wasm

package main

import (
	"context"
	"fmt"

	host "github.com/cchamplin/plugin-bug/plugin"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	host.RegisterDataPasser(TestPlugin{})
}

type TestPlugin struct{}

func (p TestPlugin) PassData(ctx context.Context, request host.PassDataRequest) (result host.PassDataReply, err error) {
	if request.Data == nil || len(request.Data) == 0 {
		data := make(map[string]*host.State)
		for i := uint64(0); i < request.Entries; i++ {
			data[fmt.Sprintf("key-%d", i)] = &host.State{
				FieldA: "field-a",
				FieldB: "field-b",
				FieldC: "field-c",
				FieldD: "field-d",
				FieldE: "field-e",
				FieldF: "field-f",
				FieldG: "field-g",
				FieldH: "field-h",
				FieldI: "field-i",
				FieldJ: "field-j",
				FieldK: "field-k",
				FieldL: "field-l",
				FieldM: "field-m",
				FieldN: "field-n",
			}
		}
		return host.PassDataReply{
			Data: data,
		}, nil
	} else {

		data := request.Data
		for k, _ := range request.Data {
			data[k] = &host.State{
				FieldA: fmt.Sprintf("field-a-change-%d", request.Sequence),
				FieldB: fmt.Sprintf("field-b-change-%d", request.Sequence),
				FieldC: fmt.Sprintf("field-c-change-%d", request.Sequence),
				FieldD: fmt.Sprintf("field-d-change-%d", request.Sequence),
				FieldE: fmt.Sprintf("field-e-change-%d", request.Sequence),
				FieldF: fmt.Sprintf("field-f-change-%d", request.Sequence),
				FieldG: fmt.Sprintf("field-g-change-%d", request.Sequence),
				FieldH: fmt.Sprintf("field-h-change-%d", request.Sequence),
				FieldI: fmt.Sprintf("field-i-change-%d", request.Sequence),
				FieldJ: fmt.Sprintf("field-j-change-%d", request.Sequence),
				FieldK: fmt.Sprintf("field-k-change-%d", request.Sequence),
				FieldL: fmt.Sprintf("field-l-change-%d", request.Sequence),
				FieldM: fmt.Sprintf("field-m-change-%d", request.Sequence),
				FieldN: fmt.Sprintf("field-n-change-%d", request.Sequence),
			}
		}
		return host.PassDataReply{
			Data: data,
		}, nil
	}
}

func (p TestPlugin) PassDataError(ctx context.Context, request host.PassDataRequest) (result host.PassDataReply, err error) {
	return host.PassDataReply{}, fmt.Errorf("an error from the plugin")
}
