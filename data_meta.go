package sdk

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/forward"
	io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/datameta"
	ddxf_sdk "github.com/ont-bizsuite/ddxf-sdk"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
)

type DataMetaSdk struct {
	ddxfAPIAddr     string
	ddxfContractSdk *ddxf_sdk.DdxfSdk
}

func NewDataMetaSdk(ddxfAPIAddr, ontologyApiAddr string, payer *ontology_go_sdk.Account) *DataMetaSdk {

	ddxfContractSdk := ddxf_sdk.NewDdxfSdk(ontologyApiAddr)
	ddxfContractSdk.SetPayer(payer)
	return &DataMetaSdk{
		ddxfAPIAddr:     ddxfAPIAddr,
		ddxfContractSdk: ddxfContractSdk,
	}
}

func (m *DataMetaSdk) SetDDXFAPIAddr(ddxfAPIAddr string) {
	m.ddxfAPIAddr = ddxfAPIAddr
}

func (m *DataMetaSdk) CreateDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.CreateDataMetaInput, controller *ontology_go_sdk.Account) (out io.CreateDataMetaOutput, err error) {
	res, err := m.handleInner(ontId, ontIdAcc, input, io.CreateDataMetaURI, controller)
	out = res.(io.CreateDataMetaOutput)
	return
}

func (m *DataMetaSdk) UpdateDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.UpdateDataMetaInput, controller *ontology_go_sdk.Account) (out io.UpdateDataMetaOutput, err error) {
	res, err := m.handleInner(ontId, ontIdAcc, input, io.UpdateDataMetaURI, controller)
	out = res.(io.UpdateDataMetaOutput)
	return
}

func (m *DataMetaSdk) RemoveDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.RemoveDataMetaInput, controller *ontology_go_sdk.Account) (out io.RemoveDataMetaOutput, err error) {
	res, err := m.handleInner(ontId, ontIdAcc, input, io.RemoveDataMetaURI, controller)
	out = res.(io.RemoveDataMetaOutput)
	return
}

func (m *DataMetaSdk) handleInner(ontId string, ontIdAcc *ontology_go_sdk.Account, input interface{}, uri string, controller *ontology_go_sdk.Account) (data interface{}, err error) {
	bs, err := json.Marshal(input)
	if err != nil {
		return
	}
	pk := keypair.SerializePublicKey(ontIdAcc.GetPublicKey())
	sig, err := ontIdAcc.Sign(bs)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"DDXF_ONTID": ontId,
		"DDXF_PK":    hex.EncodeToString(pk),
		"DDXF_SIGN":  hex.EncodeToString(sig),
	}
	code, _, res, err := forward.PostJSONRequest(m.ddxfAPIAddr+uri, bs, header)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		err = fmt.Errorf("error code is not http.StatusOk")
		return
	}
	if res == nil {
		err = fmt.Errorf("res is nil")
		return
	}
	var txHex string
	if strings.Contains(uri, "remove") {
		output := io.CreateDataMetaOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		data = output
	} else if strings.Contains(uri, "update") {
		output := io.UpdateDataMetaOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		data = output
	} else if strings.Contains(uri, "create") {
		output := io.CreateDataMetaOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		data = output
	} else {
		err = fmt.Errorf("not support method: %s", uri)
		return
	}

	tx, err := utils.TransactionFromHexString(txHex)
	if err != nil {
		return
	}
	mutTx, err := tx.IntoMutable()
	if err != nil {
		return
	}
	err = m.ddxfContractSdk.SignTx(mutTx, controller)
	if err != nil {
		return
	}
	txHash, err := m.ddxfContractSdk.GetOntologySdk().SendTransaction(mutTx)
	if err != nil {
		return
	}
	event, err := m.ddxfContractSdk.GetSmartCodeEvent(txHash.ToHexString())
	if err != nil {
		return
	}
	if event.State != 1 {
		err = fmt.Errorf("tx event state is not 1")
		return
	}
	return
}

//TODO
func (m *DataMetaSdk) getDataMeta() {
}
