package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/MhAhmadAliPL/hello/testutil/keeper"
	"github.com/MhAhmadAliPL/hello/x/hello/keeper"
	"github.com/MhAhmadAliPL/hello/x/hello/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HelloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
