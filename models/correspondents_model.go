package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MobileMoneyProvider struct {
	DepositID            string  `json:"depositId"`
	MNO                  string  `json:"-"`
	Correspondent        string  `json:"correspondent"`
	Country              string  `json:"-"`
	Currency             string  `json:"currency"`
	DecimalsInAmount     int     `json:"-"`
	Amount               float64 `json:"amount"`
	Payer                Payer   `json:"payer"`
	CustomerTimestamp    string  `json:"customerTimestamp"`
	StatementDescription string  `json:"statementDescription"`
}

type Address struct {
	Value string `json:"value"`
}

type Payer struct {
	Type    string  `json:"type"`
	Address Address `json:"address"`
}

func (mobileMoney *MobileMoneyProvider) CreateProvider() {
	switch mobileMoney.Correspondent {
	case "MTN_MOMO_BEN":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_BEN"
		mobileMoney.Country = "BEN"
		mobileMoney.Currency = "XOF"
		mobileMoney.DecimalsInAmount = 0
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_CMR":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_CMR"
		mobileMoney.Country = "CMR"
		mobileMoney.Currency = "XAF"
		mobileMoney.DecimalsInAmount = 0
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "ORANGE_CMR":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Orange"
		mobileMoney.Correspondent = "ORANGE_CMR"
		mobileMoney.Country = "CMR"
		mobileMoney.Currency = "XAF"
		mobileMoney.DecimalsInAmount = 0
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_CIV":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_CIV"
		mobileMoney.Country = "CIV"
		mobileMoney.Currency = "XOF"
		mobileMoney.DecimalsInAmount = 0
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "ORANGE_CIV":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Orange"
		mobileMoney.Correspondent = "ORANGE_CIV"
		mobileMoney.Country = "CIV"
		mobileMoney.Currency = "XOF"
		mobileMoney.DecimalsInAmount = 0
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "VODACOM_MPESA_COD":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Vodacom"
		mobileMoney.Correspondent = "VODACOM_MPESA_COD"
		mobileMoney.Country = "COD"
		mobileMoney.Currency = "CDF"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_COD":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_COD"
		mobileMoney.Country = "COD"
		mobileMoney.Currency = "CDF"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "ORANGE_COD":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Orange"
		mobileMoney.Correspondent = "ORANGE_COD"
		mobileMoney.Country = "COD"
		mobileMoney.Currency = "CDF"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_GHA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_GHA"
		mobileMoney.Country = "GHA"
		mobileMoney.Currency = "GHS"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTELTIGO_GHA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "AT"
		mobileMoney.Correspondent = "AIRTELTIGO_GHA"
		mobileMoney.Country = "GHA"
		mobileMoney.Currency = "GHS"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "VODAFONE_GHA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Vodafone"
		mobileMoney.Correspondent = "VODAFONE_GHA"
		mobileMoney.Country = "GHA"
		mobileMoney.Currency = "GHS"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MPESA_KEN":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MPesa"
		mobileMoney.Correspondent = "MPESA_KEN"
		mobileMoney.Country = "KEN"
		mobileMoney.Currency = "KES"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_MWI":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_MWI"
		mobileMoney.Country = "MWI"
		mobileMoney.Currency = "MWK"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "TNM_MWI":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "TNM"
		mobileMoney.Correspondent = "TNM_MWI"
		mobileMoney.Country = "MWI"
		mobileMoney.Currency = "MWK"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_NGA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_NGA"
		mobileMoney.Country = "NGA"
		mobileMoney.Currency = "NGN"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_NGA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_NGA"
		mobileMoney.Country = "NGA"
		mobileMoney.Currency = "NGN"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_RWA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_RWA"
		mobileMoney.Country = "RWA"
		mobileMoney.Currency = "RWF"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_RWA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_RWA"
		mobileMoney.Country = "RWA"
		mobileMoney.Currency = "RWF"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "FREE_SEN":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Free"
		mobileMoney.Correspondent = "FREE_SEN"
		mobileMoney.Country = "SEN"
		mobileMoney.Currency = "XOF"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "ORANGE_SEN":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Orange"
		mobileMoney.Correspondent = "ORANGE_SEN"
		mobileMoney.Country = "SEN"
		mobileMoney.Currency = "XOF"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_TZA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_TZA"
		mobileMoney.Country = "TZA"
		mobileMoney.Currency = "TZS"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "VODACOM_TZA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Vodacom"
		mobileMoney.Correspondent = "VODACOM_TZA"
		mobileMoney.Country = "TZA"
		mobileMoney.Currency = "TZS"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "TIGO_TZA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Tigo"
		mobileMoney.Correspondent = "TIGO_TZA"
		mobileMoney.Country = "TZA"
		mobileMoney.Currency = "TZS"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "HALOTEL_TZA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Halotel"
		mobileMoney.Correspondent = "HALOTEL_TZA"
		mobileMoney.Country = "TZA"
		mobileMoney.Currency = "TZS"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_OAPI_UGA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_OAPI_UGA"
		mobileMoney.Country = "UGA"
		mobileMoney.Currency = "UGX"
		mobileMoney.DecimalsInAmount = 0 // Not supported
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_UGA":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_UGA"
		mobileMoney.Country = "UGA"
		mobileMoney.Currency = "UGX"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "AIRTEL_OAPI_ZMB":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Airtel"
		mobileMoney.Correspondent = "AIRTEL_OAPI_ZMB"
		mobileMoney.Country = "ZMB"
		mobileMoney.Currency = "ZMW"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "MTN_MOMO_ZMB":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "MTN"
		mobileMoney.Correspondent = "MTN_MOMO_ZMB"
		mobileMoney.Country = "ZMB"
		mobileMoney.Currency = "ZMW"
		mobileMoney.DecimalsInAmount = 2
		mobileMoney.CustomerTimestamp = time.Now().Format(time.RFC3339)
	case "ZAMTEL_ZMB":
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = "Zamtel"
		mobileMoney.Correspondent = "ZAMTEL_ZMB"
		mobileMoney.Country = "ZMB"
		mobileMoney.Currency = "ZMW"
		mobileMoney.DecimalsInAmount = 2
	default:
		mobileMoney.DepositID = uuid.New().String()
		mobileMoney.MNO = ""
		mobileMoney.Correspondent = ""
		mobileMoney.Country = ""
		mobileMoney.Currency = ""
		mobileMoney.DecimalsInAmount = -1
		fmt.Println("Invalid Country_MNO:")
	}
}

func printProvider(provider MobileMoneyProvider) {
	fmt.Printf("%-10s %-20s %-10s %-5s %d\n", provider.MNO, provider.Correspondent, provider.Country, provider.Currency, provider.DecimalsInAmount)
}
