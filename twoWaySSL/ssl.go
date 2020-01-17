package main

import "crypto/x509"
import "github.com/myafeier/log"
import "io/ioutil"

import "net/http"

import "crypto/tls"

import "os/signal"

import "os"

import "syscall"

const (
	CaCert     = "./keys/ca.crt"
	ServerKey  = "./keys/srv.key"
	ServerCert = "./keys/srv.crt"
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
		Addr:      ":1433",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}
	go func() {
		log.Info("Listening 1433...")
		err = s.ListenAndServeTLS(ServerCert, ServerKey)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		log.Info("Listening 80...")
		http.ListenAndServe(":80", mux)
	}()
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			log.Info("system exit!")
			os.Exit(0)
		}
	}

}
