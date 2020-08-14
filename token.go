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
	tokenContract   string
	ddxfContractSdk *ddxf_sdk.DdxfSdk
}

func NewTokenSdk(ddxfAPIAddr, ontologyApiAddr, tokenContract string) *TokenSdk {
	ddxfContractSdk := ddxf_sdk.NewDdxfSdk(ontologyApiAddr)
	return &TokenSdk{
		ddxfAPIAddr:     ddxfAPIAddr,
		tokenContract:   tokenContract,
		ddxfContractSdk: ddxfContractSdk,
	}
}

func (ts *TokenSdk) SetDDXFAPIAddr(ddxfAPIAddr string) {
	ts.ddxfAPIAddr = ddxfAPIAddr
}

func (ts *TokenSdk) VerifyToken(input io.VerifyTokenInput) (err error) {
	res, err := ts.request(nil, input, io.VerifyTokenURI)
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

func (ts *TokenSdk) UseToken(ontIdAcc *ontology_go_sdk.Account, input io.UseTokenInput) (out io.UseTokenOutput, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	if input.Address == "" {
		input.Address = ontIdAcc.Address.ToHexString()
	}

	res, err := ts.request(ontIdAcc, input, io.UseTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) TransferToken(ontIdAcc *ontology_go_sdk.Account, input io.TransferTokenInput) (out io.TransferTokenOutput, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.TransferTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) GenerateToken(ontIdAcc *ontology_go_sdk.Account, input io.GenerateTokenInput) (txHash, tokenId string, err error) {
	if input.Auth == "" {
		input.Auth = ontIdAcc.Address.ToHexString()
	}
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}

	res, err := ts.request(ontIdAcc, input, io.GenerateTokenURI)
	if err != nil {
		return
	}
	var out io.GenerateTokenOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	txHash, err = ts.handTx(out.Tx, ontIdAcc)
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

func (ts *TokenSdk) SetTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.SetTokenAgentInput) (txHash string, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.SetTokenAgentURI)
	if err != nil {
		return
	}
	var out io.SetTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	txHash, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) AddTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.AddTokenAgentInput) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.AddTokenAgentURI)
	if err != nil {
		return
	}
	var out io.AddTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) RemoveTokenAgent(ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenAgentInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.RemoveTokenAgentURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenAgentOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) GetToken(ontIdAcc *ontology_go_sdk.Account, input io.GetTokenInput) (out io.GetTokenOutput, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.GetTokenURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) CreateTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.CreateTokenTemplateInput) (tokenTemplateId string, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	if input.Address == "" {
		input.Address = ontIdAcc.Address.ToHexString()
	}
	if input.Template.Endpoint == "" {
		input.Template.Endpoint = ts.ddxfAPIAddr
	}
	res, err := ts.request(ontIdAcc, input, io.CreateTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.CreateTokenTemplateOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	txHash, err := ts.handTx(out.Tx, ontIdAcc)
	if err != nil {
		return
	}
	evt, _ := ts.ddxfContractSdk.GetSmartCodeEvent(txHash)
	for _, notify := range evt.Notify {
		states := notify.States.([]interface{})
		if len(states) == 4 && states[0] == "createTokenTemplate" {
			tokenTemplateId = states[3].(string)
		}
	}
	return
}

func (ts *TokenSdk) UpdateTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.UpdateTokenTemplateInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.UpdateTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.UpdateTokenTemplateOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}
func (ts *TokenSdk) RemoveTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.RemoveTokenTemplateURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenTemplateOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) GetTokenTemplate(ontIdAcc *ontology_go_sdk.Account, input io.GetTokenTemplateInput) (out io.GetTokenTemplateOutput, err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.GetTokenTemplateURI)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	return
}

func (ts *TokenSdk) SetTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.SetTokenTemplateAuthInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.SetTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.SetTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) AddTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.AddTokenTemplateAuthInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.AddTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.AddTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) RemoveTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.RemoveTokenTemplateAuthInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.RemoveTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.RemoveTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) ClearTokenTemplateAuth(ontIdAcc *ontology_go_sdk.Account, input io.ClearTokenTemplateAuthInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.ClearTokenTemplateAuthURI)
	if err != nil {
		return
	}
	var out io.ClearTokenTemplateAuthOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (ts *TokenSdk) SetMPContract(ontIdAcc *ontology_go_sdk.Account, input io.SetMPContractInput) (err error) {
	if input.TokenContract == "" {
		input.TokenContract = ts.tokenContract
	}
	res, err := ts.request(ontIdAcc, input, io.SetMPContractURI)
	if err != nil {
		return
	}
	var out io.SetMPContractOutput
	err = json.Unmarshal(res, &out)
	if err != nil {
		return
	}
	_, err = ts.handTx(out.Tx, ontIdAcc)
	return
}

func (m *TokenSdk) request(ontIDAcc *ontology_go_sdk.Account, input interface{}, uri string) (res []byte, err error) {

	bs, err := json.Marshal(input)
	if err != nil {
		return
	}
	var header map[string]string
	if ontIDAcc != nil {
		ontID := "did:ont:" + ontIDAcc.Address.ToBase58()
		pk := keypair.SerializePublicKey(ontIDAcc.GetPublicKey())
		var sig []byte
		sig, err = ontIDAcc.Sign(bs)
		if err != nil {
			return
		}
		header = map[string]string{
			"DDXF_ONTID": ontID,
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
