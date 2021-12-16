package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidCreator   = sdkerrors.Register(ModuleName, 1100, "creator address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")
	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1104, "game by id not found: %s")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1105, "message creator is not a player: %s")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1106, "player tried to play out of turn: %s")
	ErrWrongMove        = sdkerrors.Register(ModuleName, 1107, "wrong move")
)
