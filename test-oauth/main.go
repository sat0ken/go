package main

import (
	"crypto/sha256"
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

// https://tex2e.github.io/rfc-translater/html/rfc7636.html
// 付録B. S256 code_challenge_methodの例
const verifier string = "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"

func readJson() {

	data, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &secrets)
	return
}

// https://auth0.com/docs/authorization/flows/call-your-api-using-the-authorization-code-flow-with-pkce#javascript-sample
func base64URLEncode() string {

	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])

}

func login(w http.ResponseWriter, req *http.Request) {
	//authEndpoint := "https://accounts.google.com/o/oauth2/v2/auth?"
	authEndpoint := "http://localhost:8081/auth?"

	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("client_id", secrets["web"].(map[string]interface{})["client_id"].(string))
	values.Add("state", "xyz")
	values.Add("scope", "https://www.googleapis.com/auth/photoslibrary.readonly")
	values.Add("redirect_uri", "http://127.0.0.1:8080/callback")

	// PKCE用パラメータ
	values.Add("code_challenge_method", "S256")
	values.Add("code_challenge", base64URLEncode())

	// 認可エンドポイントにリダイレクト
	http.Redirect(w, req, authEndpoint+values.Encode(), 302)
}

func tokenRequest(query url.Values) (map[string]interface{}, error) {

	tokenEndpoint := "https://www.googleapis.com/oauth2/v4/token"
	values := url.Values{}
	values.Add("client_id", secrets["web"].(map[string]interface{})["client_id"].(string))
	values.Add("client_secret", secrets["web"].(map[string]interface{})["client_secret"].(string))
	values.Add("grant_type", "authorization_code")

	// 取得した認可コードをトークンのリクエストにセット
	values.Add("code", query.Get("code"))
	values.Add("redirect_uri", "http://127.0.0.1:8080/callback")
	// PKCE用パラメータ
	values.Add("code_verifier", verifier)

	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("何かエラー : %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	return data, nil
}

// 認可してからの戻るところ
func callback(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	// トークンをリクエストする
	result, err := tokenRequest(query)
	if err != nil {
		log.Println(err)
	}

	photoAPI := "https://photoslibrary.googleapis.com/v1/mediaItems"

	req, err = http.NewRequest("GET", photoAPI, nil)
	if err != nil {
		log.Println(err)
	}
	// 取得したアクセストークンをHeaderにセットしてリソースサーバにリクエストを送る
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", result["access_token"].(string)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("http status code is %d, err: %s", resp.StatusCode, err)
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

	log.Printf("start server...\n")
	readJson()
	http.HandleFunc("/login", login)
	http.HandleFunc("/callback", callback)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
