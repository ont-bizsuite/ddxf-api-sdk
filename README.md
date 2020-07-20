# sdk for ddxf api

This repository contains 4 sdks, one each for `meta`, `data meta`, `marketplace` and `token`.

Each sdk is paired with respective restful apis(swagger doc):
1. [`meta api`](https://github.com/ont-bizsuite/ddxf-api-sdk/blob/master/pkg/io/meta/swagger.json)
2. [`data meta api`](https://github.com/ont-bizsuite/ddxf-api-sdk/blob/master/pkg/io/datameta/swagger.json)
3. [`marketplace api`](https://github.com/ont-bizsuite/ddxf-api-sdk/blob/master/pkg/io/mp/swagger.json)
4. [`token api`](https://github.com/ont-bizsuite/ddxf-api-sdk/blob/master/pkg/io/token/swagger.json)

`meta` sdk is used to save various meta info.

Other sdks mainly do 2 things:
1. call ddxf api to generate a transaction
2. send the transaction and retrieve events

