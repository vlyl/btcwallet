package ord

import "github.com/btcsuite/btcd/wire"

type SatPoint struct {
	OutPoint wire.OutPoint
	Offset   uint64
}
