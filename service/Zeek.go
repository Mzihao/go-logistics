package service

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"go-logistics/model"
	"go-logistics/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type ZeekServer struct{}

func (z ZeekServer) SearchRouter(barcode string) (int, map[string]interface{}) {
	reqUrl := "https://ap1-zeek2door-api.ks-it.co/Index/get_order_log"
	
	result := make(map[string]interface{})
	result["weblink"] = "https://www.zeek.one"
	result["carrier_code"] = "zeek"
	result["barcode"] = barcode

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("order_sn", barcode)
	_ = writer.WriteField("language", "zh_HK")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return 500, result
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", reqUrl, payload)

	if err != nil {
		fmt.Println(err)
		return 500, result
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Origin", "https://www.zeek.one")
	req.Header.Add("Referer", "https://www.zeek.one/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
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
	dataStr := string(body)
	//fmt.Println(dataStr)
	isError := gjson.Get(dataStr, "error")
	var trackInfo []map[string]string
	result["trackInfo"] = trackInfo
	if isError.Num != 0 {
		trackInfo = append(
			trackInfo, map[string]string{"Date": time.Now().Format("2006-01-02 15:04:05"), "StatusDescription": "訂單正在創建"})
	} else {
		data := gjson.Get(dataStr, "data.#")
		for i := 0; i < int(data.Num); i++ {
			if gjson.Get(dataStr, "data."+strconv.Itoa(i)+".detail").Str != "" {
				trackInfo = append(trackInfo, map[string]string{
					"Date": gjson.Get(dataStr, "data."+strconv.Itoa(i)+".addtime").Str,
					"StatusDescription": gjson.Get(dataStr, "data."+strconv.Itoa(i)+".detail").Str
				})
			}
		}
	}
	num, logistics := model.GetLogisticsByTrackingNumber(barcode, "zeek")
	if num == 0 {
		id := utils.GetSnowflakeId()
		data := model.Logistics{ID: id, TrackingNumber: barcode, CarrierCode: "zeek"}
		model.CreateLogistics(&data)
		result["id"] = id
	} else {
		result["id"] = logistics.ID
	}
	result["trackInfo"] = trackInfo
	return 200, result
}
