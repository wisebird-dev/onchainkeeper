package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "wasmapp/testutil/keeper"
	"wasmapp/x/onchainkeeper/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OnchainkeeperKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
