package models

import "fmt"

type MobileMoneyProvider struct {
	MNO              string
	Correspondent    string
	Country          string
	Currency         string
	DecimalsInAmount int
}

func createProvider(country_MNO string) MobileMoneyProvider {
	var provider MobileMoneyProvider

	switch country_MNO {
	case "Benin_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_BEN",
			Country:          "BEN",
			Currency:         "XOF",
			DecimalsInAmount: 0,
		}
	case "Cameroon_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_CMR",
			Country:          "CMR",
			Currency:         "XAF",
			DecimalsInAmount: 0,
		}
	case "Cameroon_Orange":
		provider = MobileMoneyProvider{
			MNO:              "Orange",
			Correspondent:    "ORANGE_CMR",
			Country:          "CMR",
			Currency:         "XAF",
			DecimalsInAmount: 0,
		}
	case "IvoryCoast_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_CIV",
			Country:          "CIV",
			Currency:         "XOF",
			DecimalsInAmount: 0,
		}
	case "IvoryCoast_Orange":
		provider = MobileMoneyProvider{
			MNO:              "Orange",
			Correspondent:    "ORANGE_CIV",
			Country:          "CIV",
			Currency:         "XOF",
			DecimalsInAmount: 0,
		}
	case "Congo_Vodacom":
		provider = MobileMoneyProvider{
			MNO:              "Vodacom",
			Correspondent:    "VODACOM_MPESA_COD",
			Country:          "COD",
			Currency:         "CDF",
			DecimalsInAmount: 2,
		}
	case "Congo_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_COD",
			Country:          "COD",
			Currency:         "CDF",
			DecimalsInAmount: 2,
		}
	case "Congo_Orange":
		provider = MobileMoneyProvider{
			MNO:              "Orange",
			Correspondent:    "ORANGE_COD",
			Country:          "COD",
			Currency:         "CDF",
			DecimalsInAmount: 2,
		}
	case "Ghana_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_GHA",
			Country:          "GHA",
			Currency:         "GHS",
			DecimalsInAmount: 2,
		}
	case "Ghana_AT":
		provider = MobileMoneyProvider{
			MNO:              "AT",
			Correspondent:    "AIRTELTIGO_GHA",
			Country:          "GHA",
			Currency:         "GHS",
			DecimalsInAmount: 2,
		}
	case "Ghana_Vodafone":
		provider = MobileMoneyProvider{
			MNO:              "Vodafone",
			Correspondent:    "VODAFONE_GHA",
			Country:          "GHA",
			Currency:         "GHS",
			DecimalsInAmount: 2,
		}
	case "Kenya_MPesa":
		provider = MobileMoneyProvider{
			MNO:              "MPesa",
			Correspondent:    "MPESA_KEN",
			Country:          "KEN",
			Currency:         "KES",
			DecimalsInAmount: 0, // Not supported
		}
	case "Malawi_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_MWI",
			Country:          "MWI",
			Currency:         "MWK",
			DecimalsInAmount: 2,
		}
	case "Malawi_TNM":
		provider = MobileMoneyProvider{
			MNO:              "TNM",
			Correspondent:    "TNM_MWI",
			Country:          "MWI",
			Currency:         "MWK",
			DecimalsInAmount: 2,
		}
	case "Nigeria_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_NGA",
			Country:          "NGA",
			Currency:         "NGN",
			DecimalsInAmount: 0, // Not supported
		}
	case "Nigeria_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_NGA",
			Country:          "NGA",
			Currency:         "NGN",
			DecimalsInAmount: 2,
		}
	case "Rwanda_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_RWA",
			Country:          "RWA",
			Currency:         "RWF",
			DecimalsInAmount: 0, // Not supported
		}
	case "Rwanda_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_RWA",
			Country:          "RWA",
			Currency:         "RWF",
			DecimalsInAmount: 0, // Not supported
		}
	case "Senegal_Free":
		provider = MobileMoneyProvider{
			MNO:              "Free",
			Correspondent:    "FREE_SEN",
			Country:          "SEN",
			Currency:         "XOF",
			DecimalsInAmount: 0, // Not supported
		}
	case "Senegal_Orange":
		provider = MobileMoneyProvider{
			MNO:              "Orange",
			Correspondent:    "ORANGE_SEN",
			Country:          "SEN",
			Currency:         "XOF",
			DecimalsInAmount: 0, // Not supported
		}
	case "Tanzania_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_TZA",
			Country:          "TZA",
			Currency:         "TZS",
			DecimalsInAmount: 2,
		}
	case "Tanzania_Vodacom":
		provider = MobileMoneyProvider{
			MNO:              "Vodacom",
			Correspondent:    "VODACOM_TZA",
			Country:          "TZA",
			Currency:         "TZS",
			DecimalsInAmount: 2,
		}
	case "Tanzania_Tigo":
		provider = MobileMoneyProvider{
			MNO:              "Tigo",
			Correspondent:    "TIGO_TZA",
			Country:          "TZA",
			Currency:         "TZS",
			DecimalsInAmount: 0, // Not supported
		}
	case "Tanzania_Halotel":
		provider = MobileMoneyProvider{
			MNO:              "Halotel",
			Correspondent:    "HALOTEL_TZA",
			Country:          "TZA",
			Currency:         "TZS",
			DecimalsInAmount: 0, // Not supported
		}
	case "Uganda_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_OAPI_UGA",
			Country:          "UGA",
			Currency:         "UGX",
			DecimalsInAmount: 0, // Not supported
		}
	case "Uganda_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_UGA",
			Country:          "UGA",
			Currency:         "UGX",
			DecimalsInAmount: 2,
		}
	case "Zambia_Airtel":
		provider = MobileMoneyProvider{
			MNO:              "Airtel",
			Correspondent:    "AIRTEL_OAPI_ZMB",
			Country:          "ZMB",
			Currency:         "ZMW",
			DecimalsInAmount: 2,
		}
	case "Zambia_MTN":
		provider = MobileMoneyProvider{
			MNO:              "MTN",
			Correspondent:    "MTN_MOMO_ZMB",
			Country:          "ZMB",
			Currency:         "ZMW",
			DecimalsInAmount: 2,
		}
	case "Zambia_Zamtel":
		provider = MobileMoneyProvider{
			MNO:              "Zamtel",
			Correspondent:    "ZAMTEL_ZMB",
			Country:          "ZMB",
			Currency:         "ZMW",
			DecimalsInAmount: 2,
		}
	default:
		fmt.Println("Invalid Country_MNO:", country_MNO)
	}

	return provider
}

func printProvider(provider MobileMoneyProvider) {
	fmt.Printf("%-10s %-20s %-10s %-5s %d\n", provider.MNO, provider.Correspondent, provider.Country, provider.Currency, provider.DecimalsInAmount)
}
