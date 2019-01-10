package v0

import (
	"fmt"

	bankInvariant "github.com/irisnet/irishub/modules/bank/invariants"
	distrInvariant "github.com/irisnet/irishub/modules/distribution/invariants"
	stakeInvariant "github.com/irisnet/irishub/modules/stake/invariants"
	sdk "github.com/irisnet/irishub/types"
)

func (p *ProtocolV0) runtimeInvariants() []sdk.Invariant {
	return []sdk.Invariant{
		bankInvariant.NonnegativeBalanceInvariant(p.accountMapper),
		distrInvariant.ValAccumInvariants(p.distrKeeper, p.StakeKeeper),
		stakeInvariant.SupplyInvariants(p.bankKeeper, p.StakeKeeper,
			p.feeKeeper, p.distrKeeper, p.accountMapper),
		stakeInvariant.PositivePowerInvariant(p.StakeKeeper),
	}
}

func (p *ProtocolV0) assertRuntimeInvariants(ctx sdk.Context) {
	if p.invariantLevel != sdk.InvariantError && p.invariantLevel != sdk.InvariantPanic {
		return
	}
	invariants := p.runtimeInvariants()
	for _, inv := range invariants {
		if err := inv(ctx); err != nil {
			if p.invariantLevel == sdk.InvariantPanic {
				panic(fmt.Errorf("invariant broken: %s", err))
			} else {
				p.logger.Error(fmt.Sprintf("Invariant broken: height %d, reason %s", ctx.BlockHeight(), err.Error()))
			}
		}
	}
}
