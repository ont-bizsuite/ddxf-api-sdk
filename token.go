package sdk

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ont-bizsuite/ddxf-api-sdk/pkg/forward"
	io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/token"
	ddxf_sdk "github.com/ont-bizsuite/ddxf-sdk"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
)

type TokenSdk struct {
	ddxfAPIAddr     string
	ddxfContractSdk *ddxf_sdk.DdxfSdk
}

func NewTokenSdk(ddxfAPIAddr, ontologyApiAddr string, payer *ontology_go_sdk.Account) *TokenSdk {
	ddxfContractSdk := ddxf_sdk.NewDdxfSdk(ontologyApiAddr)
	ddxfContractSdk.SetPayer(payer)
	return &TokenSdk{
		ddxfAPIAddr:     ddxfAPIAddr,
		ddxfContractSdk: ddxfContractSdk,
	}
}

func (ts *TokenSdk) SetDDXFAPIAddr(ddxfAPIAddr string) {
	ts.ddxfAPIAddr = ddxfAPIAddr
}

func (ts *TokenSdk) VerifyToken(input io.VerifyTokenInput) (err error) {
	res, err := ts.request("", nil, input, io.VerifyTokenURI)
	if err != nil {
		return
	}
	var out io.VerifyTokenOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	if out.OK {
		return nil
	}
	return fmt.Errorf("failed")
}

func (ts *TokenSdk) UseToken(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.UseTokenInput, tokenOwner *ontology_go_sdk.Account) (out io.UseTokenOutput, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.UseTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, tokenOwner)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) TransferToken(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.TransferTokenInput, tokenOwner *ontology_go_sdk.Account) (out io.TransferTokenOutput, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.TransferTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, tokenOwner)
	return
}

func (ts *TokenSdk) GenerateToken(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.GenerateTokenInput, acc *ontology_go_sdk.Account) (tokenId string, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.GenerateTokenURI)
	if err != nil {
		return
	}
	var out io.GenerateTokenOutput
	err = json.Unmarshal(res, out)
	if err != nil {
		return
	}
	txHash, err := ts.handTx(out.Tx, acc)
	if err != nil {
		return
	}
	evt, err := ts.ddxfContractSdk.GetSmartCodeEvent(txHash)
	if err != nil {
		return
	}
	for _, notify := range evt.Notify {
		states := notify.States.([]interface{})
		if len(states) == 5 && states[0] == "generateDToken" {
			tokenId = states[4].(string)
		}
	}
	return
}

func (ts *TokenSdk) SetTokenAgent(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.SetTokenAgentInput, acc *ontology_go_sdk.Account) {
	res, err := ts.request(ontId, ontIdAcc, input, io.SetTokenAgentURI)
	if err != nil {
		return
	}
	var out io.SetTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) AddTokenAgent(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.AddTokenAgentInput, acc *ontology_go_sdk.Account) {
	res, err := ts.request(ontId, ontIdAcc, input, io.AddTokenAgentURI)
	if err != nil {
		return
	}
	var out io.AddTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) RemoveTokenAgent(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.AddTokenAgentInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.RemoveTokenAgentURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) GetToken(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.GetTokenInput, acc *ontology_go_sdk.Account) (out io.GetTokenOutput, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.GetTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) CreateTokenTemplate(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.CreateTokenTemplateInput, acc *ontology_go_sdk.Account) (tokenTemplateId string, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.CreateTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	txHash, err := ts.handTx(out.Tx, acc)
	if err != nil {
		return
	}
	evt, _ := ts.ddxfContractSdk.GetSmartCodeEvent(txHash)
	for _, notify := range evt.Notify {
		states := notify.States.([]interface{})
		if len(states) != 4 && states[0] == "createTokenTemplate" {
			tokenTemplateId = states[3].(string)
		}
	}
	return
}

func (ts *TokenSdk) UpdateTokenTemplate(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.UpdateTokenTemplateInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.UpdateTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.UpdateTokenTemplateOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}
func (ts *TokenSdk) RemoveTokenTemplate(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.RemoveTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenTemplateOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) GetTokenTemplate(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.GetTokenTemplateInput, acc *ontology_go_sdk.Account) (out io.GetTokenTemplateOutput, err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.GetTokenTemplateURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) SetTokenTemplateAuth(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.SetTokenTemplateAuthInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.SetTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.SetTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) AddTokenTemplateAuth(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.AddTokenTemplateAuthInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.AddTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.AddTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) RemoveTokenTemplateAuth(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.RemoveTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (ts *TokenSdk) ClearTokenTemplateAuth(ontId string, ontIdAcc *ontology_go_sdk.Account, input io.ClearTokenTemplateAuthInput, acc *ontology_go_sdk.Account) (err error) {
	res, err := ts.request(ontId, ontIdAcc, input, io.ClearTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.ClearTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, acc)
	return
}

func (m *TokenSdk) request(ontId string, ontIdAcc *ontology_go_sdk.Account, input interface{}, uri string) (res []byte, err error) {
	bs, err := json.Marshal(input)
	if err != nil {
		return
	}
	var header map[string]string
	if ontIdAcc != nil {
		pk := keypair.SerializePublicKey(ontIdAcc.GetPublicKey())
		var sig []byte
		sig, err = ontIdAcc.Sign(bs)
		if err != nil {
			return
		}
		header = map[string]string{
			"DDXF_ONTID": ontId,
			"DDXF_PK":    hex.EncodeToString(pk),
			"DDXF_SIGN":  hex.EncodeToString(sig),
		}
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
	return
}
func (ts *TokenSdk) handTx(txHex string, controller *ontology_go_sdk.Account) (txhash string, err error) {
	tx, err := utils.TransactionFromHexString(txHex)
	if err != nil {
		return
	}
	mutTx, err := tx.IntoMutable()
	if err != nil {
		return
	}
	err = ts.ddxfContractSdk.SignTx(mutTx, controller)
	if err != nil {
		return
	}
	txHash, err := ts.ddxfContractSdk.GetOntologySdk().SendTransaction(mutTx)
	if err != nil {
		return
	}
	event, err := ts.ddxfContractSdk.GetSmartCodeEvent(txHash.ToHexString())
	if err != nil {
		return
	}
	if event.State != 1 {
		err = fmt.Errorf("tx event state is not 1")
		return
	}
	txhash = txHash.ToHexString()
	return
}
