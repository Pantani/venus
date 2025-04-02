package keeper

import (
	"venus/x/venus/types"
)

var _ types.QueryServer = Keeper{}
