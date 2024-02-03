package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/venus/testutil/keeper"
	"github.com/ignite/venus/x/venus/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VenusKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
