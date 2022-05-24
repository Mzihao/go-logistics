package service

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
  )  

type Kuaidi100Server struct{}

func (k Kuaidi100Server) SearchRouter(barcode string) (int, map[string]interface{}) {
	result := make(map[string]interface{})
	posttype, postid := getCodeAndNum(barcode)
	if posttype == "" {
		return 500, result
	}
	reqUrl := "https://www.kuaidi100.com/query?type=" + posttype + "&postid=" + postid + "&temp=0.8926864614539474&phone="
	result["weblink"] = ""
	result["carrier_code"] = posttype + "-100"
	result["barcode"] = barcode
	return 200, result
}

func getCodeAndNum(id string) (code string, num string) {
	url := "https://www.kuaidi100.com/autonumber/autoComNum?resultv2=1&text=" + id
	client := &http.Client {}
	payload := strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", ""
	}
	json := string(body)
	return "", ""
}