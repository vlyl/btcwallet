package ordjson

type OrdInscribeCmd struct {
	File          string `json:"file"`
	Destination   string `json:"destination;omitempty"`
	FeeRate       int    `json:"feerate;omitempty"`
	CommitFeeRate int    `json:"commitfeerate;omitempty"`
	SatPoint      string `json:"satpoint;omitempty"`
	DryRun        bool   `json:"dryrun;omitempty"`
	NoBackup      bool   `json:"nobackup;omitempty"`
	NoLimit       bool   `json:"nolimit;omitempty"`
}
