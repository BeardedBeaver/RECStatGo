package main

import (
	statServer "RECStatGo/pkg/server"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Options struct describes server parameters
type Options struct {
	Address  string `json:"address"`
	Endpoint string `json:"endpoint"`
	DB       string `json:"db"`
	Uri      string `json:"uri"`
}

// getOptions func reads server parameters
// from a json file placed next to the binary.
// (See Options struct)
func getOptions() (Options, error) {
	binaryFileName, err := os.Executable()
	if err != nil {
		return Options{}, err
	}
	text, err := ioutil.ReadFile(filepath.Join(filepath.Dir(binaryFileName), "./RECStatGo.json"))
	if err != nil {
		return Options{}, err
	}
	var options Options
	err = json.Unmarshal(text, &options)
	if err != nil {
		return Options{}, err
	}
	return options, nil
}

func main() {
	options, err := getOptions()
	if err != nil {
		log.Fatal(err)
	}

	server, err := statServer.NewServer(options.DB, options.Uri)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	http.HandleFunc(options.Endpoint, server.Serve)
	err = http.ListenAndServe(options.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
