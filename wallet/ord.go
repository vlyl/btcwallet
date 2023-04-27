package wallet

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/ord/ordjson"
)

func (w *Wallet) Inscribe(cmd *ordjson.OrdInscribeCmd) ([]byte, error) {
	utxos, err := w.UnspentOutputs()
	if err != nil {
		log.Error("Inscribe: UnspentOutputs error: ", err)
		return nil, err
	}

	wire.NewTxIn()
}
