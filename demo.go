package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	ApiUrl        = "https://pay...."
	PartnerId     = "2000056..."
	privateKeyPem = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC19KqOI+f0cEcUAuEtE1Se0Cvbz5rmgsbOVH/KrIHkSyjAJD84
eqZGOKliAthUv7k/2XZe4dFzvIOPooK8ZwMzRqELar9XpTQGVIWJzlNyv+sswIvq
4FfGf9+OmjEwdvbBZHslAbDsqPXtxUg0TP/bBcpwTWyKNs3TlK++hKULTwIDAQAB
...
-----END RSA PRIVATE KEY-----`
)

type PostData struct {
	Service      string  `json:"Service,omitempty"`
	Version      string  `json:"Version,omitempty"`
	PartnerId    string  `json:"PartnerId,omitempty"`
	InputCharset string  `json:"InputCharset,omitempty"`
	TradeDate    string  `json:"TradeDate,omitempty"`
	TradeTime    string  `json:"TradeTime,omitempty"`
	ReturnUrl    string  `json:"ReturnUrl,omitempty"`
	Memo         string  `json:"Memo,omitempty"`
	TrxId        string  `json:"TrxId,omitempty"`
	SellerId     string  `json:"SellerId,omitempty"`
	ExpiredTime  string  `json:"ExpiredTime,omitempty"`
	MerUserId    string  `json:"MerUserId,omitempty"`
	TradeType    string  `json:"TradeType,omitempty"`
	CardBegin    string  `json:"CardBegin,omitempty"`
	CardEnd      string  `json:"CardEnd,omitempty"`
	TrxAmt       float32 `json:"TrxAmt,omitempty"`
	SmsFlag      string  `json:"SmsFlag,omitempty"`
	OrdrName     string  `json:"OrdrName,omitempty"`
	Sign         string  `json:"sign,omitempty"`
	SignType     string  `json:"SignType,omitempty"`
}

func init() {
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("Demo")
}

func main() {

	testData := &PostData{}

	testData.Service = "nmg_biz_api_quick_payment"
	testData.Version = "1.0"
	testData.PartnerId = PartnerId
	testData.InputCharset = "UTF-8"
	testData.TradeDate = time.Now().Format("20060102")
	testData.TradeTime = time.Now().Format("150405")
	testData.ReturnUrl = "http://dev.chanpay.com/receive.php"
	testData.Memo = "暂无备注"
	testData.TrxId = fmt.Sprintf("%020d", time.Now().Unix())
	testData.SellerId = PartnerId
	testData.ExpiredTime = "90m"
	testData.MerUserId = "11"
	testData.TradeType = "11"
	testData.CardBegin = "621700"
	testData.CardEnd = "6570"
	testData.TrxAmt = 0.01
	testData.SmsFlag = "0"
	testData.OrdrName = "测试"

	signature, err := sign(testData)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	testData.Sign = base64.StdEncoding.EncodeToString(signature) + "s"
	testData.SignType = "RSA"

	log.Printf("singature:%s \n", testData.Sign)

	err = GetData(testData)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}

func GetData(data *PostData) (err error) {
	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)

	params := url.Values{}

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i)
		if !value.IsZero() {
			params.Add(key.Name, fmt.Sprintf("%s", value.Interface()))
		}
	}

	fmt.Printf("params: %s \n", params.Encode())
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", ApiUrl, params.Encode()), nil)
	if err != nil {
		log.Fatal("Error:%s \n", err.Error())
		return
	}
	fmt.Printf("req: %s\n", req.URL.RequestURI())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error:%s \n", err.Error())
		return
	}

	fmt.Printf("Resp: code:%d,codeString:%s\n", resp.StatusCode, resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:%s \n", err.Error())
		return
	}
	fmt.Printf("Resp%s \n", body)
	return
}

func sign(data *PostData) (signature []byte, err error) {

	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	var params []string

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i)
		if !value.IsZero() {
			params = append(params, fmt.Sprintf("%s=%v", key.Name, value.Interface()))
		}
	}

	parmsData := sha256.Sum256([]byte(strings.Join(params, "&")))

	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		err = fmt.Errorf("nil block")
		return
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Errorf("Error:%s \n", err.Error())
		return
	}
	signature, err = rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, parmsData[:])
	return
}
