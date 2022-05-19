package service

import (
	"fmt"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/types"
	"github.com/lestrrat-go/libxml2/xpath"
	"go-logistics/model"
	"go-logistics/utils"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

type MapleLogisticsServer struct{}

func (m MapleLogisticsServer) SearchRouter(barcode string) (int, map[string]interface{}) {
	reqUrl := "https://www.25431010.tw/Search.php"
	if len(barcode) != 9 && len(barcode) != 10 && len(barcode) != 12 {
		return 4001, nil
	}
	result := make(map[string]interface{})

	// 自动更新cookie
	var jar, _ = cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	tik := getTik(reqUrl, client)
	payload := strings.NewReader("tik=" + tik + "&BARCODE1=" + barcode + "&BARCODE2=&BARCODE3=")

	req, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		fmt.Println(err)
		return 500, nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://www.25431010.tw")
	req.Header.Add("Referer", "https://www.25431010.tw/Search.php")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 500, nil
	}
	defer res.Body.Close()

	//body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 500, nil
	}
	doc, err := libxml2.ParseHTMLReader(res.Body)
	if err != nil {
		panic("failed to parse HTML: " + err.Error())
	}
	defer doc.Free()
	doc.Walk(func(n types.Node) error {
		return nil
	})
	nodes := xpath.NodeList(doc.Find(`//tr/td/text()`))
	messageList := make([]string, 5)
	for i := 0; i < len(nodes); i++ {
		if len(nodes[i].String()) == 2 {
			break
		} else {
			messageList = append(messageList, nodes[i].String())
		}
	}
	messageList = messageList[5:]
	var trackInfo []map[string]string
	for i := 0; i < len(messageList)/3; i++ {
		trackInfo = append(trackInfo, map[string]string{"Date": strings.Replace(messageList[3*i+1], "/", "-", -1), "StatusDescription": messageList[3*i+2]})
	}
	trackInfo = rev(trackInfo)
	result["weblink"] = "https://www.25431010.tw/Search.php"
	result["carrier_code"] = "bld-express"
	result["trackInfo"] = trackInfo
	data := model.Logistics{ID: utils.GetSnowflakeId(), TrackingNumber: barcode, CarrierCode: "bld-express"}
	num := model.GetLogisticsByTrackingNumber(barcode, "bld-express")
	if num == 0 {
		model.CreateLogistics(&data)
	}
	return 200, result
}

func getTik(reqUrl string, client *http.Client) string {
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Printf("get failed, err:%v\n\n", err)
		return ""
	}

	// 添加请求头
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://www.25431010.tw")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Add("Referer", "https://www.25431010.tw/Search.php")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n\n", err)
		return ""
	}
	myRegex := regexp.MustCompile(`<input name="tik" id="tik" type="hidden" value="(.*?)" />`)
	found := myRegex.FindAllStringSubmatch(string(body), -1)
	return found[0][1]
}

// rev 切片反转
func rev(slice []map[string]string) []map[string]string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
