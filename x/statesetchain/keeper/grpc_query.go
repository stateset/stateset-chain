package keeper

import (
	"github.com/stateset/statesetchain/x/statesetchain/types"
)

var _ types.QueryServer = Keeper{}
