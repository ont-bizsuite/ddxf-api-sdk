# DDXF Addon SDK Usage Documentation

## Overview

在addon store上部署完ddxf addon后，会拿到：

1. `mpContractAddress`(mp contract address)
2. `tokenContractAddress`(token contract address)
3. `ddxfAPIServer`(ddxf api server address)



将上述参数传给4个sdk的初始化方法后，即可使用！(参照：[ddxf api sdk使用文档](https://github.com/ont-bizsuite/ddxf-api-sdk/blob/master/usage.md))

### meta sdk 初始化

```golang

import "github.com/ont-bizsuite/ddxf-api-sdk"

ddxfAPIServer := ""
sdk := sdk.NewMetaSdk(ddxfAPIServer)
```



### data meta sdk 初始化

```golang

import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
	"github.com/ontio/ontology-crypto/signature"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	)

ddxfAPIServer := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
payer, _ := ontology_go_sdk.NewAccountFromPrivateKey("private key ...", signature.SHA256withECDSA)
sdk := sdk.NewDataMetaSdk(ddxfAPIServer, ontologyApiAddr, payer)
```


### marketplace sdk 初始化

```golang

import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
	"github.com/ontio/ontology-crypto/signature"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	)

ddxfAPIServer := ""
mpContract := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
payer, _ := ontology_go_sdk.NewAccountFromPrivateKey("private key ...", signature.SHA256withECDSA)
sdk := sdk.NewMarketplace(ddxfAPIServer, ontologyApiAddr, mpContract, payer)
```


### token sdk 初始化

```golang

import (
	"github.com/ont-bizsuite/ddxf-api-sdk"
	"github.com/ontio/ontology-crypto/signature"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	)

ddxfAPIServer := ""
tokenContract := ""
ontologyApiAddr := "http://dappnode1.ont.io:20336"
payer, _ := ontology_go_sdk.NewAccountFromPrivateKey("private key ...", signature.SHA256withECDSA)
sdk := sdk.NewTokenSdk(ddxfAPIServer, ontologyApiAddr, tokenContract, payer)
```