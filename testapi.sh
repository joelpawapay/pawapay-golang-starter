#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{"correspondent": "MTN_MOMO_GHA", "amount":2.1, "payer":{"type":"MSISDN", "address":{"value":"233541053593"}},"statementDescription":"Golang" }' "http://127.0.0.1:7070/deposits"
