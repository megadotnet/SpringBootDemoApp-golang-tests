package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"net/url"
	"os"
	"testing"
	sw "testingappdemo/swagger"

	_ "go.uber.org/automaxprocs"
)

var client *sw.APIClient

func TestMain(m *testing.M) {
	cfg := sw.NewConfiguration()

	//creating the proxyURL
	//proxyStr := "http://localhost:8888"
	//proxyURL, err := url.Parse(proxyStr)
	//if err != nil {
	//	fmt.Print(err)
	//}

	//login process
	loginModel := CreateLoginModel()
	tr := &http.Transport{
		//Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	myclient := &http.Client{Transport: tr}
	payloadBytes, err := json.Marshal(loginModel)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", cfg.BasePath+"/api/authenticate", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := myclient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	localVarBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	var respObj sw.JWTToken
	jsonerr := json.Unmarshal(localVarBody, &respObj)
	if jsonerr != nil {
		fmt.Println("error:", jsonerr)
		return
	}

	cfg.AddDefaultHeader("Authorization", "Bearer "+respObj.Token)

	client = sw.NewAPIClient(cfg)
	retCode := m.Run()
	os.Exit(retCode)
}

func waitOnFunctions(t *testing.T, errc chan error, n int) {
	for i := 0; i < n; i++ {
		err := <-errc
		if err != nil {
			t.Errorf("Error performing concurrent test")
			t.Log(err)
		}
	}
}

func CreateLoginModel() sw.LoginVm {
	return sw.LoginVm{
		RememberMe: false,
		Username:   "admin",
		Password:   "admin"}
}
