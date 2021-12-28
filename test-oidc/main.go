package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var secrets map[string]interface{}

func readJson() {

	data, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &secrets)
	return

}

func login(w http.ResponseWriter, req *http.Request) {
	endpoint := "https://accounts.google.com/o/oauth2/v2/auth?"
	v := url.Values{}
	v.Add("response_type", "code")
	v.Add("client_id", secrets["web"].(map[string]interface{})["client_id"].(string))
	v.Add("state", "xyz")
	v.Add("scope", "openid profile")
	v.Add("redirect_uri", "http://127.0.0.1:8080/callback")
	v.Add("nonce", "abc")

	log.Printf("http redirect to: %s", fmt.Sprintf("%s%s", endpoint, v.Encode()))
	http.Redirect(w, req, fmt.Sprintf("%s%s", endpoint, v.Encode()), 302)
}

func tokenRequest(query url.Values) (map[string]interface{}, error) {

	endpoint := "https://www.googleapis.com/oauth2/v4/token"
	v := url.Values{}
	v.Add("client_id", secrets["web"].(map[string]interface{})["client_id"].(string))
	v.Add("client_secret", secrets["web"].(map[string]interface{})["client_secret"].(string))
	v.Add("grant_type", "authorization_code")
	v.Add("code", query.Get("code"))
	v.Add("redirect_uri", "http://127.0.0.1:8080/callback")

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var token map[string]interface{}
	json.Unmarshal(body, &token)

	log.Printf("token response :%s\n", string(body))

	return token, nil
}

// https://qiita.com/ekzemplaro/items/e24d6af29ddc01683417
// 文字数が 4 の倍数になるように調整
func adjustB64Length(jwt string) (b64 string) {
	replace := strings.NewReplacer("-", "+", "_", "/")
	b64 = replace.Replace(jwt)

	if len(jwt)%4 != 0 {
		addLength := len(jwt) % 4
		for i := 0; i < addLength; i++ {
			b64 += "="
		}
	}

	return b64
}

func decodeJWT(idToken string) {
	tmp := strings.Split(idToken, ".")

	header := adjustB64Length(tmp[0])
	payload := adjustB64Length(tmp[1])

	decHeader, _ := base64.StdEncoding.DecodeString(string(header))
	decPayload, _ := base64.StdEncoding.DecodeString(string(payload))

	fmt.Println(string(decHeader))
	fmt.Println(string(decPayload))
}

func callback(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	token, err := tokenRequest(query)
	if err != nil {
		log.Println(err)
	}
	decodeJWT(token["id_token"].(string))

	userInfoURL := "https://openidconnect.googleapis.com/v1/userinfo"
	req, err = http.NewRequest("GET", userInfoURL, nil)
	if nil != err {
		log.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token["access_token"].(string)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
	w.Write(body)

}

func main() {

	log.Println("start server...")
	readJson()
	http.HandleFunc("/login", login)
	http.HandleFunc("/callback", callback)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
