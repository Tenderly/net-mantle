package blobpool

import (
	"github.com/tenderly/net-mantle/common"
	"github.com/tenderly/net-mantle/core"
	"github.com/tenderly/net-mantle/core/txpool"
	"github.com/tenderly/net-mantle/core/types"
	"github.com/tenderly/net-mantle/event"
)

// SubscribeNewPreconfTxEvent subscribes to new preconf transaction events.
func (p *BlobPool) SubscribeNewPreconfTxEvent(ch chan<- core.NewPreconfTxEvent) event.Subscription {
	return p.preconfTxFeed.Subscribe(ch)
}

// SubscribeNewPreconfTxRequestEvent subscribes to new preconf transaction request events.
func (p *BlobPool) SubscribeNewPreconfTxRequestEvent(ch chan<- *core.NewPreconfTxRequest) event.Subscription {
	return p.preconfTxRequestFeed.Subscribe(ch)
}

func (p *BlobPool) PendingPreconfTxs(filter txpool.PendingFilter) ([]*types.Transaction, map[common.Address][]*txpool.LazyTransaction) {
	// Blob pool does not support preconf transactions
	return nil, p.Pending(filter)
}

// PreconfReady closes the preconfReadyCh channel to notify the miner that preconf is ready
func (p *BlobPool) PreconfReady() {
	// Do nothing
}

func (p *BlobPool) SetPreconfTxStatus(txHash common.Hash, status core.PreconfStatus) {
	// Do nothing
}
