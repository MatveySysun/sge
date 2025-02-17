package types

import "fmt"

// DefaultGenesis returns the default  genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:         DefaultParams(),
		DepositList:    []Deposit{},
		WithdrawalList: []Withdrawal{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, w := range gs.WithdrawalList {
		found := false
		for _, d := range gs.DepositList {
			if w.DepositorAddress == d.Creator &&
				w.SportEventUID == d.SportEventUID &&
				w.ParticipationIndex == d.ParticipationIndex {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("the deposit for the depositor address %s, sport event uid %s and participation index %d not found for the withdrawal",
				w.DepositorAddress,
				w.SportEventUID,
				w.ParticipationIndex)
		}
	}

	// TODO: extend validations for sport-event existence
	// and etc.

	return gs.Params.Validate()
}
