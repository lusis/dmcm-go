package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"
	"dmcm/utils"
	"crypto/tls"
)

type EnstratusError struct {
	// { "error" : { "message" : "Invalid access keys" } }
	Error struct {
		Message string
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf(`
		Usage: dmcm-cli namespace/Resource [foo=bar]
		`)
		os.Exit(1)
	}
	es_access_key, ok := syscall.Getenv("ES_ACCESS_KEY")
	if ok != true {
		fmt.Printf("You must set the environment variable: ES_ACCESS_KEY\n")
		os.Exit(1)
	}
	es_secret_key, ok := syscall.Getenv("ES_SECRET_KEY")
	if ok != true {
		fmt.Printf("You must set the environment variable: ES_SECRET_KEY\n")
		os.Exit(1)
	}
	es_detail := func() string {
		ed, ok := syscall.Getenv("ES_DETAIL")
		if ok != true {
			return "basic"
		} else {
			return ed
		}
	}
	es_endpoint := func() string {
		ep, ok := syscall.Getenv("ES_ENDPOINT")
		if ok != true {
			return "https://api.enstratus.com"
		} else {
			return ep
		}
	}
	es_api_version := func() string {
		ev, ok := syscall.Getenv("ES_API_VERSION")
		if ok != true {
			return "2012-06-15"
		} else {
			return ev
		}
	}
	es_resource := func() string {
		if len(os.Args) <= 1 {
			return "geography/Cloud"
		} else {
			return os.Args[1]
		}
	}

	sign_url := es_endpoint() + "/api/enstratus/" + es_api_version() + "/" + es_resource()
	var get_url string
	if len(os.Args) == 3 {
		get_url = sign_url + "?" + os.Args[2]
	} else {
		get_url = sign_url
	}
	request, _ := http.NewRequest("GET", get_url, nil)
	request.Header.Add("accept", "application/json")
	request.Header.Add("x-esauth-access", es_access_key)
	request.Header.Add("x-esauth-timestamp", string(utils.GetTimeString()))
	request.Header.Add("x-es-details", es_detail())
	request.Header.Add("user-agent", utils.ES_UA)
	sig := utils.SignRequest(es_access_key, es_secret_key,
		utils.ES_UA, "GET",
		es_resource(), es_api_version())
	request.Header.Add("x-esauth-signature", sig)
	verifySsl := func() bool {
		_, ok := syscall.Getenv("ES_NOVERIFY_SSL")
		if ok != true {
			return false
		} else {
			return true
		}
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: verifySsl()},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		switch {
		case response.StatusCode >= 400 && response.StatusCode <= 599:
			e := &EnstratusError{}
			err := json.Unmarshal([]byte(contents), &e)
			if err != nil {
				fmt.Printf("JSON Decoding error: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", e.Error.Message)
			os.Exit(1)
		}
		var c interface{}
		jsonerr := json.Unmarshal([]byte(contents), &c)
		if jsonerr != nil {
			fmt.Printf("JSON Decoding error: %s\n", jsonerr)
			os.Exit(1)
		}
		spew.Dump(c)
	}
}
