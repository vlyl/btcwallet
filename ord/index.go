package ord

import (
	"github.com/btcsuite/btcd/wire"
	"go.etcd.io/bbolt"
	"hash"
	"sync/atomic"
)

type Index struct {
	Auth                            Auth
	Client                          Client
	Database                        bbolt.DB
	Path                            string
	FistInscriptionHeight           uint64
	GenesisBlockCoinbaseTransaction wire.MsgTx
	GenesisBlockCoinbaseTxid        hash.Hash
	HeightLimit                     uint64
	Reorged                         atomic.Bool
	RpcUrl                          string
}
