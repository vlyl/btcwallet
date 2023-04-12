package ord

type Auth struct {
	UserPass struct {
		User string
		Pass string
	}
	CookieFile string
}

// Since we write ord integrate in btcwallet, we don't need to setup rpc connect once again.
// Theoretically, we can use btcwallet rpc client to connect to bitcoind, and utxo from btcwallet directly.
