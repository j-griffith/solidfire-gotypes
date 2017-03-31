package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	sftypes "github.com/j-griffith/solidfire-gotypes"
)

func issueReq(endpoint, methodName string, params interface{}) (response []byte, err error) {
	id := 101 // Request ID, suggest using a random int64 here
	data, err := json.Marshal(map[string]interface{}{
		"method": methodName,
		"id":     id,
		"params": params,
	})

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http := &http.Client{Transport: tr}
	resp, err := http.Post(endpoint,
		"json-rpc",
		strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("error response from SolidFire API request: %v", err)
		return nil, errors.New("device API error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	errresp := sftypes.APIError{}
	json.Unmarshal([]byte(body), &errresp)
	if errresp.Error.Code != 0 {
		fmt.Printf("error detected in API response: %+v\n%+v", errresp, string(body))
		err = fmt.Errorf("device API error: %+v", errresp.Error.Name)
	}
	return body, err
}

func main() {
	endpoint := "https://adminLogin:adminPasswd@10.117.36.101/json-rpc/8.0"
	req := sftypes.ListVolumesForAccountRequest{AccountID: 75493}
	response, err := issueReq(endpoint, "ListVolumesForAccount", req)
	if err != nil {
		fmt.Printf("Bummer dude: %+v", err)
		os.Exit(1)
	}
	fmt.Printf("Yay volumes: %+v", string(response))

	// Oh... what's that you say, you want a response object?
	var result sftypes.ListVolumesResult
	json.Unmarshal([]byte(response), &result)

	for _, v := range result.Result.Volumes {
		fmt.Printf("Here's a volume: %+v", v)
	}
}
