package ord

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"go.etcd.io/bbolt"
	"path/filepath"
)

type Index struct {
	Auth                            Auth
	Database                        *bbolt.DB
	Path                            string
	FistInscriptionHeight           uint64
	GenesisBlockCoinbaseTransaction wire.MsgTx
	GenesisBlockCoinbaseTxid        chainhash.Hash
	HeightLimit                     uint64
	Reorged                         bool
	RpcUrl                          string
	Err                             error
}

type Info struct {
	BlocksIndexed    uint64            `json:"blocks_indexed"`
	BranchPages      int               `json:"branch_pages"`
	FragmentedBytes  int               `json:"fragmented_bytes"`
	IndexFileSize    uint64            `json:"index_file_size"`
	IndexPath        string            `json:"index_path"`
	LeafPages        int               `json:"leaf_pages"`
	MetadataBytes    int               `json:"metadata_bytes"`
	OutputsTraversed uint64            `json:"outputs_traversed"`
	PageSize         int               `json:"page_size"`
	SatRanges        uint64            `json:"sat_ranges"`
	StoredBytes      int               `json:"stored_bytes"`
	Transactions     []TransactionInfo `json:"transactions"`
	TreeHeight       int               `json:"tree_height"`
	UtxosIndexed     int               `json:"utxos_indexed"`
}

type TransactionInfo struct {
	StartingBlockCount uint64 `json:"starting_block_count"`
	StartingTimestamp  uint64 `json:"starting_timestamp"`
}

const (
	HeightToBlockHash = "height_to_block_hash"
)

func (i *Index) Error() error {
	return i.Err
}

func (i *Index) Open(cfg Config) *Index {
	const dbName = "ord_index.db"
	db, err := bbolt.Open(filepath.Join(cfg.DataDir, dbName), 0600, bbolt.DefaultOptions)
	if err != nil {
		log.Critical("Open: bbolt.Open error: ", err)
		return nil
	}

	dbtx, err := db.Begin(true)

	if err != nil {
		log.Critical("Open: db.Begin error: ", err)
		return nil
	}
	defer dbtx.Rollback()

	// todo create buckets
	_, err = dbtx.CreateBucketIfNotExists([]byte(HeightToBlockHash))
	if err != nil {
		log.Critical("Open: CreateBucketIfNotExists error: ", err)
		return nil
	}

	if err := dbtx.Commit(); err != nil {
		log.Critical("Open: dbtx.Commit error: ", err)
		return nil
	}

	genesisBlockCoinbaseTransaction := cfg.ChainParam.GenesisBlock.Transactions[0]

	return &Index{
		Auth:                            Auth{CookieFile: cfg.CookieFile},
		Database:                        db,
		Path:                            cfg.DataDir,
		FistInscriptionHeight:           cfg.FirstIncantationHeight,
		GenesisBlockCoinbaseTransaction: *genesisBlockCoinbaseTransaction,
		GenesisBlockCoinbaseTxid:        genesisBlockCoinbaseTransaction.TxHash(),
		HeightLimit:                     cfg.HeightLimit,
		RpcUrl:                          cfg.RpcUrl,
		Reorged:                         false,
	}
}
