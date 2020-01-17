package main

import "crypto/x509"
import "github.com/myafeier/log"
import "io/ioutil"

import "net/http"

import "crypto/tls"

const (
	CaCert     = "/Users/xiafei/test/gotest/twoWaySSL/keys/ca.crt"
	ServerKey  = "/Users/xiafei/test/gotest/twoWaySSL/keys/srv.key"
	ServerCert = "/Users/xiafei/test/gotest/twoWaySSL/keys/srv.crt"
)

func init() {
	log.SetLogLevel(log.INFO)
	log.SetPrefix("test")
}

func main() {
	pool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile(CaCert)
	if err != nil {
		log.Error("%s", err.Error())
		return
	}
	pool.AppendCertsFromPEM(caCert)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Info("req!")
		w.Write([]byte("Hello world!"))
	})
	tlsConfig := &tls.Config{
		ClientCAs:  pool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	s := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}
	err = s.ListenAndServeTLS(ServerCert, ServerKey)
	if err != nil {
		panic(err)
	}

}
