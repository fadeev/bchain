package keeper

import (
	"github.com/fadeev/bchain/x/bchain/types"
)

var _ types.QueryServer = Keeper{}
