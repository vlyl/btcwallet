package ord

import (
	"github.com/btcsuite/btcd/chaincfg"
)

func InscriptionContentSizeLimit(params chaincfg.Params) uint64 {
	if params.Name == "mainnet" || params.Name == "testnet" {
		return 4 * 1024 * 1024
	}
	if params.Name == "regtest" || params.Name == "signet" {
		return 1024
	}
	return 0
}
