package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/checkers/x/checkers/types"
	rules "github.com/mkXultra/checkers/x/checkers/rules"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
			panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)

	storedGame := types.StoredGame{
			Creator: msg.Creator,
			Index:   newIndex,
			Game:    rules.New().String(),
			Red:     msg.Red,
			Black:   msg.Black,
	}

	err := storedGame.Validate()
	if err != nil {
			return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
			IdValue: newIndex,
	}, nil
}
