package checkers_test

import (
	"testing"

	keepertest "github.com/mkXultra/checkers/testutil/keeper"
	"github.com/mkXultra/checkers/x/checkers"
	"github.com/mkXultra/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		NextGame: &types.NextGame{
			IdValue: 56,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.NextGame, got.NextGame)
	// this line is used by starport scaffolding # genesis/test/assert
}
