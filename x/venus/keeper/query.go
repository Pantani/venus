package keeper

import (
	"github.com/ignite/venus/x/venus/types"
)

var _ types.QueryServer = Keeper{}
