package migration

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	cron0 "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	cid "github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"

	cron2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
)

type cronMigrator struct {
}

func (m *cronMigrator) MigrateState(ctx context.Context, store cbor.IpldStore, head cid.Cid, _ abi.TokenAmount) (cid.Cid, abi.TokenAmount, error) {
	var inState cron0.State
	if err := store.Get(ctx, head, &inState); err != nil {
		return cid.Undef, big.Zero(), err
	}

	outState := cron2.State{Entries: make([]cron2.Entry, len(inState.Entries))}
	for i, e := range inState.Entries {
		outState.Entries[i] = cron2.Entry(e) // Identical
	}
	newHead, err := store.Put(ctx, &outState)
	return newHead, big.Zero(), err
}
