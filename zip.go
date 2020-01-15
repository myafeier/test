package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Test struct {
	Id   int64
	Name string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header.Get("Content-type"))
		// b, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	log.Println(err)
		// 	w.WriteHeader(400)
		// 	return
		// }
		// log.Printf("%s\n", b)
		gzipReader, err := gzip.NewReader(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		defer gzipReader.Close()

		body, err := ioutil.ReadAll(gzipReader)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		log.Printf("%s", body)

	})
	go http.ListenAndServe(":9090", mux)
	time.Sleep(3 * time.Second)
	go UploadZippedJSON()
	time.Sleep(1 * time.Second)
}
func UploadZippedJSON() {

	data := new(Test)
	data.Id = 1
	data.Name = "test"
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	zipBuf, err := zip(jsonData)
	if err != nil {
		panic(err)
	}

	req, err := http.Post("http://localhost:9090", "Application/json", zipBuf)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(req.StatusCode)
}

func zip(data []byte) (bf *bytes.Buffer, err error) {

	bf = new(bytes.Buffer)
	writer, err := gzip.NewWriterLevel(bf, 9)
	if err != nil {
		return
	}

	_, err = writer.Write(data)
	if err != nil {
		return
	}
	defer writer.Close()
	return
}
