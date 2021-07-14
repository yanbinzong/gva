package app

import (
"encoding/json"
"fmt"
"net/http"
)

func main() {
	http.HandleFunc("/openid", getOpenID)
	http.ListenAndServe(":2020", nil)
}

func getOpenID(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	if request.Method != http.MethodPost {
		return
	}

	var codeMap map[string]string
	err := json.NewDecoder(request.Body).Decode(&codeMap)
	if err != nil {
		return
	}

	defer request.Body.Close()

	code := codeMap["code"]
	openid, err := sendWxAuthAPI(code)
	if err != nil {
		return
	}
	fmt.Println("my openid", openid)
}

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx817442a15fb8c499"
	appSecret       = "2f9a834e8232294d80b2880d8f25bd96"
)

func sendWxAuthAPI(code string) (string, error) {
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	var wxMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return wxMap["openid"], nil
}