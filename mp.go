package sdk

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/forward"
	io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/mp"
	ddxf_sdk "github.com/ont-bizsuite/ddxf-sdk"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
)

// Marketplace ...
type Marketplace struct {
	ddxfAPIAddr     string
	mpContract      string
	ddxfContractSdk *ddxf_sdk.DdxfSdk
}

func NewMarketplace(ddxfAPIAddr, ontologyApiAddr, mpContract string, payer *ontology_go_sdk.Account) *Marketplace {
	ddxfContractSdk := ddxf_sdk.NewDdxfSdk(ontologyApiAddr)
	ddxfContractSdk.SetPayer(payer)
	return &Marketplace{
		ddxfAPIAddr:     ddxfAPIAddr,
		mpContract:      mpContract,
		ddxfContractSdk: ddxfContractSdk,
	}
}

func (mp *Marketplace) PublishItem(ontIdAcc *ontology_go_sdk.Account, input io.PublishItemInput, seller, mpAcc *ontology_go_sdk.Account) (out io.PublishItemOutput, err error) {

	if input.MPContract == "" {
		input.MPContract = mp.mpContract
	}
	res, err := mp.handleInner(ontIdAcc, input, io.PublishItemURI, []*ontology_go_sdk.Account{seller, mpAcc})
	out = res.(io.PublishItemOutput)
	return
}
func (mp *Marketplace) UpdateItem(ontIdAcc *ontology_go_sdk.Account, input io.UpdateItemInput, seller *ontology_go_sdk.Account) (out io.UpdateItemOutput, err error) {

	if input.MPContract == "" {
		input.MPContract = mp.mpContract
	}
	res, err := mp.handleInner(ontIdAcc, input, io.UpdateItemURI, []*ontology_go_sdk.Account{seller})
	out = res.(io.UpdateItemOutput)
	return
}

func (mp *Marketplace) DeleteItem(ontIdAcc *ontology_go_sdk.Account, input io.DeleteItemInput, acc *ontology_go_sdk.Account) (out io.DeleteItemOutput, err error) {

	if input.MPContract == "" {
		input.MPContract = mp.mpContract
	}
	res, err := mp.handleInner(ontIdAcc, input, io.DeleteItemURI, []*ontology_go_sdk.Account{acc})
	out = res.(io.DeleteItemOutput)
	return
}

func (mp *Marketplace) BuyItem(ontIdAcc *ontology_go_sdk.Account, input io.BuyItemInput, buyer *ontology_go_sdk.Account) (out io.BuyItemOutput, err error) {

	if input.MPContract == "" {
		input.MPContract = mp.mpContract
	}
	res, err := mp.handleInner(ontIdAcc, input, io.BuyItemURI, []*ontology_go_sdk.Account{buyer})
	out = res.(io.BuyItemOutput)
	return
}

func (mp *Marketplace) GetItem(input io.GetItemInput) (*io.GetItemOutput, error) {
	if input.MPContract == "" {
		input.MPContract = mp.mpContract
	}
	bs, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	code, _, res, err := forward.PostJSONRequest(mp.ddxfAPIAddr+io.GetItemURI, bs, nil)
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

func (mp *Marketplace) handleInner(ontIDAcc *ontology_go_sdk.Account, input interface{}, uri string, controller []*ontology_go_sdk.Account) (out interface{}, err error) {
	ontID := "did:ont:" + ontIDAcc.Address.ToBase58()
	bs, err := json.Marshal(input)
	if err != nil {
		return
	}
	pk := keypair.SerializePublicKey(ontIDAcc.GetPublicKey())
	sig, err := ontIDAcc.Sign(bs)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"DDXF_ONTID": ontID,
		"DDXF_PK":    hex.EncodeToString(pk),
		"DDXF_SIGN":  hex.EncodeToString(sig),
	}
	code, _, res, err := forward.PostJSONRequest(mp.ddxfAPIAddr+uri, bs, header)
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
	if strings.Contains(uri, "publish") {
		output := io.PublishItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		out = output
	} else if strings.Contains(uri, "update") {
		output := io.PublishItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		out = output
	} else if strings.Contains(uri, "delete") {
		output := io.DeleteItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		out = output
	} else if strings.Contains(uri, "buy") {
		output := io.BuyItemOutput{}
		err = json.Unmarshal(res, output)
		if err != nil {
			return
		}
		txHex = output.Tx
		out = output
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
		err = mp.ddxfContractSdk.SignTx(mutTx, con)
		if err != nil {
			return
		}
	}
	txHash, err := mp.ddxfContractSdk.GetOntologySdk().SendTransaction(mutTx)
	if err != nil {
		return
	}
	event, err := mp.ddxfContractSdk.GetSmartCodeEvent(txHash.ToHexString())
	if err != nil {
		return
	}
	if event.State != 1 {
		err = fmt.Errorf("tx event state is not 1")
		return
	}
	return
}
