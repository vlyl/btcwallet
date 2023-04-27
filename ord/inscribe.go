package ord

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr/musig2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/wire"
)

func Inscribe(option *Config, path string) {

	inscription, err := FromFile(option.ChainParam, path)
	if err != nil {
		panic(err)
	}

	// todo update index

	wallet
}

func CreateInscriptionTransaction(satPoint SatPoint,
	inscription Inscription,
	inscriptions map[SatPoint]Inscription,
	network string,
	utxos map[wire.OutPoint]btcutil.Amount,
	change [2]btcutil.Address,
	destination btcutil.Address,
	commitFeeRate FeeRate,
	revealFeeRate FeeRate,
	noLimit bool) (*wire.MsgTx, *wire.MsgTx, []byte /*TweakedKeyPair*/, error) {

	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Error("CreateInscriptionTransaction: btcec.NewPrivateKey error: ", err)
	}
	ctx, err := musig2.NewContext(privKey, true)

}
