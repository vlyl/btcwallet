package wallet

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/ord/ordjson"
)

func (w *Wallet) Inscribe(cmd *ordjson.OrdInscribeCmd) ([]byte, error) {
	w.UnspentOutputs()
	wire.NewTxIn()
}
