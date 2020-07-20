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
	io "github.com/zhiqiangxu/ddxf-sdk/pkg/io/mp"
)

// Marketplace ...
type Marketplace struct {
	addr string
}

func NewMarketplace(addr string) *Marketplace {
	return &Marketplace{
		addr: addr,
	}
}

func (mp *Marketplace) PublishItem(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.PublishItemInput, seller, mpAcc *ontology_go_sdk.Account) error {
	return mp.handleInner(ontId, ontIdAcc, input, io.PublishItemURI, []*ontology_go_sdk.Account{seller, mpAcc})
}
func (mp *Marketplace) UpdateItem(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.PublishItemInput, seller *ontology_go_sdk.Account) error {
	return mp.handleInner(ontId, ontIdAcc, input, io.UpdateItemURI, []*ontology_go_sdk.Account{seller})
}

func (mp *Marketplace) DeleteItem(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.DeleteItemInput, acc *ontology_go_sdk.Account) error {
	return mp.handleInner(ontId, ontIdAcc, input, io.DeleteItemURI, []*ontology_go_sdk.Account{acc})
}

func (mp *Marketplace) BuyItem(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.BuyItemInput, buyer *ontology_go_sdk.Account) error {
	return mp.handleInner(ontId, ontIdAcc, input, io.BuyItemURI, []*ontology_go_sdk.Account{buyer})
}

func (mp *Marketplace) GetItem(input io.GetItemInput) (*io.GetItemOutput, error) {
	bs, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	code, _, res, err := forward.PostJSONRequest(mp.addr+io.GetItemURI, bs, nil)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		err = fmt.Errorf("error code is not http.StatusOk")
		return nil, err
	}
	if res == nil {
		err = fmt.Errorf("res is nil")
		return nil, err
	}
	var out io.GetItemOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (m *Marketplace) handleInner(ontId string, ontIdAcc *ontology_go_sdk.Account, input interface{}, uri string, controller []*ontology_go_sdk.Account) (err error) {
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
	}
	if res == nil {
		err = fmt.Errorf("res is nil")
		return
	}
	var txHex string
	if strings.Contains(uri, "publish") {
		output := io.PublishItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
	} else if strings.Contains(uri, "update") {
		output := io.PublishItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
	} else if strings.Contains(uri, "delete") {
		output := io.DeleteItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
	} else if strings.Contains(uri, "buy") {
		output := io.BuyItemOutput{}
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
	for _, con := range controller {
		err = instance.DDXFSdk().SignTx(mutTx, con)
		if err != nil {
			return
		}
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
