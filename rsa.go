package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

var (
	privateKey string = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----`
	publicKey string = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----`
)

func RsaEncrypt(plainData []byte) (cipherText []byte, err error) {

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		err = fmt.Errorf("nil block")
		return
	}

	publicKeyData, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	rsaPubKey := publicKeyData.(*rsa.PublicKey)
	cipherText, err = rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, plainData)
	return
}
func RsaDecrypt(cipherData []byte) (plainData []byte, err error) {

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = fmt.Errorf("nil block")
		return
	}
	privateKeyData, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	plainData, err = rsa.DecryptPKCS1v15(rand.Reader, privateKeyData, cipherData)
	return
}

func main() {

	plainText := "Hello World"

	cipher, err := RsaEncrypt([]byte(plainText))
	if err != nil {
		fmt.Println(err)
		return
	}

	cipherText := base64.StdEncoding.EncodeToString(cipher)
	fmt.Println(cipherText)
	phpCipherBase64Encoded := "NKymQEECcVL6d4KMLhMDMkU7a6zd65oAQQ3x108kUnMHl8qc22JKLgiKXXNHW35QEMa3wixtafsn4+28OphZ3ov2IfG6f5iC4lB5vlND+fEkedeBESE/Jzi/z0JGQgIPi/d9sE83V7SHPk2Xh4RRMrr2aCsPh/hhbMkrg4KIeZk="
	phpCipherData, err := base64.StdEncoding.DecodeString(phpCipherBase64Encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	plainData, err := RsaDecrypt(phpCipherData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", plainData)
}
