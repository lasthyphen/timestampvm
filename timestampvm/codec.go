// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timestampvm

import (
	"github.com/lasthyphen/dijetalgo/codec"
	"github.com/lasthyphen/dijetalgo/codec/linearcodec"
)

const (
	// CodecVersion is the current default codec version
	CodecVersion = 0
)

// Codecs do serialization and deserialization
var (
	Codec codec.Manager
)

func init() {
	// Create default codec and manager
	c := linearcodec.NewDefault()
	Codec = codec.NewDefaultManager()

	// Register codec to manager with CodecVersion
	if err := Codec.RegisterCodec(CodecVersion, c); err != nil {
		panic(err)
	}
}
