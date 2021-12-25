package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/mkXultra/checkers/testutil/keeper"
	"github.com/mkXultra/checkers/x/checkers/keeper"
	"github.com/mkXultra/checkers/x/checkers/types"
)

func createTestLeaderboard(keeper *keeper.Keeper, ctx sdk.Context) types.Leaderboard {
	item := types.Leaderboard{}
	keeper.SetLeaderboard(ctx, item)
	return item
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	rst, found := keeper.GetLeaderboard(ctx)
	require.True(t, found)
	require.Equal(t, item, rst)
}
func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	createTestLeaderboard(keeper, ctx)
	keeper.RemoveLeaderboard(ctx)
	_, found := keeper.GetLeaderboard(ctx)
	require.False(t, found)
}
