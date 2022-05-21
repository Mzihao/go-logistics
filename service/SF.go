package service

//import "github.com/spf13/viper"

type ShunFengService struct{}

func (s ShunFengService) SearchRouter(barcode string) (int, map[string]interface{}) {
	// url := viper.GetString("logistics.SF.url")
	// partnerID := viper.GetString("logistics.SF.partnerID")
	// key := viper.GetString("logistics.SF.key")
	// serviceCode := "FOP_RECE_LTL_SEARCH_ROUTER"

	result := make(map[string]interface{})
	result["weblink"] = "https://www.sf-express.com/"
	result["carrier_code"] = "sf-express"
	var trackInfo []map[string]string
	result["trackInfo"] = trackInfo
	return 200, result
}

func (s ShunFengService) CreateOrder() (int, map[string]string) {
	result := make(map[string]interface{})
	result["weblink"] = "https://www.sf-express.com/"
	result["carrier_code"] = "sf-express"
	result["barcode "] = "SF1234567890"
	return 200, result
}
