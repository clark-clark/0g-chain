package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	IterateDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(delegation stakingtypes.Delegation) (stop bool))
}
