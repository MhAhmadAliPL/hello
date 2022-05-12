package keeper

import (
	"github.com/MhAhmadAliPL/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
