package chain

import (
	"sync"

	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type SerdeOptions struct {
	SerDe types.SerDeOptions
}

var defaultSerDeOptions = SerdeOptions{}

var mu sync.RWMutex

// SetSerDeOptions overrides default serialise and deserialize options
func SetSerDeOptions(noPalletIndices bool) {
	defer mu.Unlock()
	mu.Lock()
	tsdo := types.SerDeOptions{
		NoPalletIndices: noPalletIndices,
	}
	types.SetSerDeOptions(tsdo)
	defaultSerDeOptions = SerdeOptions{
		SerDe: tsdo,
	}
}
