package service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"go-logistics/global"
	"go-logistics/model"
	"go-logistics/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Kuaidi100Server struct{}

func (k Kuaidi100Server) SearchRouter(barcode string) (int, map[string]interface{}) {
	result := make(map[string]interface{})
	//client := &http.Client{}
	posttype, postid := getCodeAndNum(barcode)
	if posttype == "" {
		return 500, result
	}
	reqUrl := "https://www.kuaidi100.com/query?type=" + posttype + "&postid=" + postid + "&temp=0.8926864614539474&phone="

	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		fmt.Println(err)
		return 500, result
	}
	req.Header.Add("Referer", "https://www.kuaidi100.com/?from=openv")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")

	res, err := global.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return 500, result
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 500, result
	}
	response := string(body)

	result["weblink"] = ""
	result["carrier_code"] = posttype + "-100"
	result["barcode"] = barcode
	var trackInfo []map[string]string
	result["trackInfo"] = trackInfo
	hasMessage := gjson.Get(response, "message").Str
	if hasMessage != "ok" {
		return 4002, result
	}
	data := gjson.Get(response, "data.#")
	for i := 0; i < int(data.Num); i++ {
		trackInfo = append(trackInfo, map[string]string{
			"Date":              gjson.Get(response, "data."+strconv.Itoa(i)+".time").Str,
			"StatusDescription": gjson.Get(response, "data."+strconv.Itoa(i)+".context").Str,
		})
	}
	num, logistics := model.GetLogisticsByTrackingNumber(barcode, posttype+"-100")
	if num == 0 {
		id := utils.GetSnowflakeId()
		data := model.Logistics{ID: id, TrackingNumber: barcode, CarrierCode: posttype + "-100"}
		model.CreateLogistics(&data)
		result["id"] = id
	} else {
		result["id"] = logistics.ID
	}
	result["trackInfo"] = utils.Rev(trackInfo)
	return 200, result
}

func getCodeAndNum(id string) (code string, num string) {
	url := "https://www.kuaidi100.com/autonumber/autoComNum?resultv2=1&text=" + id
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := global.HttpClient.Do(req)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", ""
	}
	json := string(body)
	data := gjson.Get(json, "auto").Array()
	if len(data) == 0 {
		return "", ""
	}
	posttype := gjson.Get(json, "auto.0.comCode").Str
	postid := gjson.Get(json, "num").Str
	return posttype, postid
}
