# DDXF使用场景

## 作为marketplace provider参与

```golang

import (
    "github.com/ont-bizsuite/ddxf-api-sdk"
    io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/mp"
    "github.com/ontio/ontology-crypto/signature"
    "github.com/zhiqiangxu/ddxf"
    ontology_go_sdk "github.com/ontio/ontology-go-sdk"
    "github.com/ont-bizsuite/ddxf-sdk/split_policy_contract"
	)

ddxfAPIServer := ""
mpContract := ""
tokenContract := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
mpOntIDAcc, _ := ontology_go_sdk.NewAccountFromPrivateKey("mp private key ...", signature.SHA256withECDSA)
sdk := sdk.NewMarketplaceSdk(ddxfAPIServer, ontologyApiAddr, mpContract)


seller, _ := ontology_go_sdk.NewAccountFromPrivateKey("seller private key ...", signature.SHA256withECDSA)

itemID := "some item id"
itemMeta := map[string]interface{}{
    "id": 1,
    "desc": "...",
}
itemMetaHash, _ := ddxf.HashObject2Hex(itemMeta)
oep4ContractAddr := "195d72da6725e8243a52803f6de4cd93df48fc1f"
addr, _ := common2.AddressFromHexString(oep4ContractAddr)
_, err := sdk.PublishItem(
    mpOntIDAcc, 
    io.PublishItemInput{
        // item id, should be new
        ItemID : itemID,
        // meta hash of item meta
        ItemMetaHash: itemMetaHash,
        // token template ids is generated by /ddxf/dtoken/create_template
        TokenTemplateIDs []string{
            "token template id",
        }
        // token contract for each token
        TokenContracts []string{
            tokenContract,
        },
        // fee of item
        Fee: market_place_contract.Fee{
            ContractAddr: addr,
            ContractType: split_policy_contract.OEP4,
            Count:        1,
        }
        // stock of item
        Stock: 100,
        // valid until this timestamp, in seconds
        ExpiredDate: time.Now().Unix() + 30*24*3600,
    }, 
    seller)

buyer, _ := ontology_go_sdk.NewAccountFromPrivateKey("buyer private key ...", signature.SHA256withECDSA)
_, err = sdk.BuyItem(
    mpOntIDAcc, 
    io.BuyItemInput{
        // item id
        ItemID: itemID,
        // # of items to buy
        N: 1,
    }, 
    buyer)

```
## 作为data provider参与

首先为data生成dataID

```golang
import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
    "github.com/ontio/ontology-crypto/signature"
    io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/datameta"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	)

ddxfAPIServer := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
sdk := sdk.NewDataMetaSdk(ddxfAPIServer, ontologyApiAddr)

seller, _ := ontology_go_sdk.NewAccountFromPrivateKey("seller private key ...", signature.SHA256withECDSA)
dataMeta := map[string]interface{}{
    "desc": "...",
}
dataMetaHash, _ := ddxf.HashObject2Hex(dataMeta)

dataID, err := sdk.CreateDataMeta(seller, io.CreateDataMetaInput{
	// data fingerprint
	Fingerprint: "dataHash",
	// data meta id
	DataMetaHash: dataMetaHash,
})
```

然后便可以为dataID生成token template和token

```golang

import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
    "github.com/ontio/ontology-crypto/signature"
    io "github.com/ont-bizsuite/ddxf-api-sdk/pkg/io/token"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	)

ddxfAPIServer := ""
tokenContract := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
sdk := sdk.NewTokenSdk(ddxfAPIServer, ontologyApiAddr, tokenContract)

seller, _ := ontology_go_sdk.NewAccountFromPrivateKey("seller private key ...", signature.SHA256withECDSA)
buyer, _ := ontology_go_sdk.NewAccountFromPrivateKey("buyer private key ...", signature.SHA256withECDSA)




tokenMeta := map[string]interface{}{
    "read": true,
    "write": true,
}
tokenMetaHash, _ := ddxf.HashObject2Hex(tokenMeta)
tokenTemplateId, err := sdk.CreateTokenTemplate(
    seller, 
    io.CreateTokenTemplateInput{
        // template info
        Template: io.TokenTemplate {
            // data id is generated by /ddxf/data_meta/create
            DataID: dataID,
            // token meta hash for token meta
            TokenMetaHash: tokenMetaHash,
        },
    })

tokenID, err := sdk.GenerateToken(
    seller, 
    io.GenerateTokenInput{
        // template id, generated by /ddxf/dtoken/create_template
        TokenTemplateID: tokenTemplateId,
        // to account
        To: buyer.Address.ToHexString(),
        // amount of token
        N: 1,
    })

_, err = sdk.UseToken(buyer, io.UseTokenInput{
	// token id, generated by /ddxf/mp/buy_item or /ddxf/dtoken/generate
	TokenID: tokenID,
	// # of token
	N: 1,
})    
```