package venus_test

import (
	"testing"

	keepertest "venus/testutil/keeper"
	"venus/testutil/nullify"
	venus "venus/x/venus/module"
	"venus/x/venus/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VenusKeeper(t)
	venus.InitGenesis(ctx, k, genesisState)
	got := venus.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
