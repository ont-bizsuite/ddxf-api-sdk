package sdk

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/zhiqiangxu/ddxf-sdk/pkg/forward"
	"github.com/zhiqiangxu/ddxf-sdk/pkg/instance"
	io "github.com/zhiqiangxu/ddxf-sdk/pkg/io/datameta"
)

type DataMetaSdk struct {
	addr string
}

func NewDataMetaSdk(addr string) *DataMetaSdk {
	return &DataMetaSdk{
		addr: addr,
	}
}

func (m *DataMetaSdk) SetAddr(addr string) {
	m.addr = addr
}

func (m *DataMetaSdk) CreateDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.CreateDataMetaInput, controller *ontology_go_sdk.Account) error {
	return m.handleInner(ontId, ontIdAcc, input, io.CreateDataMetaURI, controller)
}

func (m *DataMetaSdk) UpdateDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.UpdateDataMetaInput, controller *ontology_go_sdk.Account) (err error) {
	err = m.handleInner(ontId, ontIdAcc, input, io.UpdateDataMetaURI, controller)
	return
}

func (m *DataMetaSdk) RemoveDataMeta(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.RemoveDataMetaInput, controller *ontology_go_sdk.Account) (err error) {
	err = m.handleInner(ontId, ontIdAcc, input, io.RemoveDataMetaURI, controller)
	return
}

func (m *DataMetaSdk) handleInner(ontId string, ontIdAcc *ontology_go_sdk.Account, input interface{}, uri string, controller *ontology_go_sdk.Account) (err error) {
	bs, err := json.Marshal(input)
	if err != nil {
		return
	}
	pk := keypair.SerializePublicKey(ontIdAcc.GetPublicKey())
	sig, err := ontIdAcc.Sign(bs)
	if err != nil {
		return err
	}
	header := map[string]string{
		"DDXF_ONTID": ontId,
		"DDXF_PK":    hex.EncodeToString(pk),
		"DDXF_SIGN":  hex.EncodeToString(sig),
	}
	code, _, res, err := forward.PostJSONRequest(m.addr+uri, bs, header)
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
	} else if strings.Contains(uri, "update") {
		output := io.UpdateDataMetaOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
	} else if strings.Contains(uri, "create") {
		output := io.CreateDataMetaOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
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
	err = instance.DDXFSdk().SignTx(mutTx, controller)
	if err != nil {
		return
	}
	txHash, err := instance.DDXFSdk().GetOntologySdk().SendTransaction(mutTx)
	if err != nil {
		return
	}
	event, err := instance.DDXFSdk().GetSmartCodeEvent(txHash.ToHexString())
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
