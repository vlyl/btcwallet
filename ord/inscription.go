package ord

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"os"
)

const ProtocolOrd = "ord"

var ContentTypeTag = []byte{1}
var BodyTag = []byte{}

const UnknownOrdError = 7303780

const (
	EmptyWitness = iota + UnknownOrdError + 1
	InvalidInscription
	KeyPathSpend
	NoInscription
)

type Inscription struct {
	Body        []byte
	ContentType []byte
}

func FromTransaction(tx *wire.MsgTx) ([]byte, error) {
	witness := tx.TxIn[0].Witness
	if len(witness) == 0 {
		return nil, errors.New("no witness")
	}

	if len(witness) == 1 {
		return nil, errors.New("no script")
	}

	annex := false
	lastElement := (witness)[len(witness)-1]
	if len(lastElement) > 0 && lastElement[0] == txscript.TaprootAnnexTag {
		annex = true
	}

	if len(witness) == 2 && annex {
		return nil, errors.New("no taproot annex")
	}

	script := (witness)[len(witness)-2]
	return txscript.NewScriptBuilder().AddData(script).Script()

}

func FromFile(params chaincfg.Params, path string) (Inscription, error) {
	var inscription Inscription
	// read file content into body
	file, err := os.Open(path)
	if err != nil {
		return inscription, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return inscription, err
	}

	fileSize := stat.Size()
	if fileSize > int64(InscriptionContentSizeLimit(params)) {
		return inscription, errors.New("file too large")
	}

	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return inscription, err
	}

	contentType, err := ContentTypeForPath(path)
	if err != nil {
		return inscription, err
	}

	return Inscription{Body: buffer, ContentType: []byte(contentType)}, nil

}

func (inscription *Inscription) AppendRevealScriptToBuilder(builder *txscript.ScriptBuilder) *txscript.ScriptBuilder {
	return builder.AddOp(txscript.OP_FALSE).
		AddOp(txscript.OP_IF).
		AddData([]byte(ProtocolOrd)).
		AddData(ContentTypeTag).AddData(inscription.ContentType).
		AddData(BodyTag).AddData(inscription.Body).
		AddOp(txscript.OP_ENDIF)
}

func (inscription *Inscription) AppendRevealScript(builder *txscript.ScriptBuilder) ([]byte, error) {
	return inscription.AppendRevealScriptToBuilder(builder).Script()
}

func (inscription *Inscription) ToWitness() (wire.TxWitness, error) {
	builder := txscript.NewScriptBuilder()
	script, err := inscription.AppendRevealScript(builder)
	if err != nil {
		return nil, err
	}
	return wire.TxWitness{script}, nil
}
