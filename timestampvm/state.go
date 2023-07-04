// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timestampvm

import (
	"github.com/lasthyphen/dijetalgo/database"
	"github.com/lasthyphen/dijetalgo/database/prefixdb"
	"github.com/lasthyphen/dijetalgo/database/versiondb"
	"github.com/lasthyphen/dijetalgo/vms/components/djtx"
)

var (
	// These are prefixes for db keys.
	// It's important to set different prefixes for each separate database objects.
	singletonStatePrefix = []byte("singleton")
	blockStatePrefix     = []byte("block")

	_ State = &state{}
)

// State is a wrapper around djtx.SingleTonState and BlockState
// State also exposes a few methods needed for managing database commits and close.
type State interface {
	// SingletonState is defined in avalanchego,
	// it is used to understand if db is initialized already.
	djtx.SingletonState
	BlockState

	Commit() error
	Close() error
}

type state struct {
	djtx.SingletonState
	BlockState

	baseDB *versiondb.Database
}

func NewState(db database.Database, vm *VM) State {
	// create a new baseDB
	baseDB := versiondb.New(db)

	// create a prefixed "blockDB" from baseDB
	blockDB := prefixdb.New(blockStatePrefix, baseDB)
	// create a prefixed "singletonDB" from baseDB
	singletonDB := prefixdb.New(singletonStatePrefix, baseDB)

	// return state with created sub state components
	return &state{
		BlockState:     NewBlockState(blockDB, vm),
		SingletonState: djtx.NewSingletonState(singletonDB),
		baseDB:         baseDB,
	}
}

// Commit commits pending operations to baseDB
func (s *state) Commit() error {
	return s.baseDB.Commit()
}

// Close closes the underlying base database
func (s *state) Close() error {
	return s.baseDB.Close()
}
