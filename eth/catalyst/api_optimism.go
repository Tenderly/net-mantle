package catalyst

import (
	"errors"

	"github.com/tenderly/net-mantle/beacon/engine"
	"github.com/tenderly/net-mantle/params"
)

// checkOptimismPayload performs Optimism-specific checks on the payload data (called during [(*ConsensusAPI).newPayload]).
func checkOptimismPayload(params engine.ExecutableData, cfg *params.ChainConfig) error {
	if cfg.IsMantleSkadi(params.Timestamp) {
		if params.WithdrawalsRoot == nil {
			return errors.New("nil withdrawalsRoot post-Skadi")
		}
	} else if params.WithdrawalsRoot != nil {
		return errors.New("non-nil withdrawalsRoot pre-Skadi")
	}

	return nil
}
