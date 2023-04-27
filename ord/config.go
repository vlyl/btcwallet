package ord

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// Config equals to Option struct in origin rust version ord
// In golang, we usually use Config to instead of Option
type Config struct {
	ChainParam             chaincfg.Params
	BitcoinDataDir         string // Load Bitcoin Core data dir from <BITCOIN_DATA_DIR>
	Config                 string // Load configuration from <CONFIG>
	ConfigDir              string // Load configuration from <CONFIG_DIR>
	CookieFile             string // Load Bitcoin Core RPC cookie file from <COOKIE_FILE>
	DataDir                string // Store index in <DATA_DIR>
	FirstIncantationHeight uint64 // Don't look for inscriptions below <FIRST_INSCRIPTION_HEIGHT>
	HeightLimit            uint64 // Limit index to <HEIGHT_LIMIT> blocks
	Index                  string // Use index at <INDEX>
	IndexSats              bool   // Track location of all satoshis
	RpcUrl                 string // Connect to Bitcoin Core(btcd) RPC at <RPC_URL>
	RegTest                bool   // Use regtest. Equivalent to `--chain regtest`.
	Signet                 bool   // Use signet. Equivalent to `--chain signet`.
	TestNet                bool   // Use testnet. Equivalent to `--chain testnet`.
	Wallet                 string // Use wallet named <WALLET> Default "ord"
}
