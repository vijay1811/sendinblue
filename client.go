package sendinblue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vijay1811/sendinblue/protocol"
)

const apiEndPoint = "https://api.sendinblue.com/v2.0/"

type SIBClient struct {
	*http.Client
	apiKey string
}

const clientfilePath = "sendinblue/client.go:"

func NewSIBClient(apiKey string) *SIBClient {
	return &SIBClient{
		Client: http.DefaultClient,
		apiKey: apiKey,
	}
}

func (c *SIBClient) sendRequest(path string, body []byte) (*protocol.Response, error) {
	//	fmt.Println(string(body))
	req, err := http.NewRequest("POST", apiEndPoint+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("%s error while making http request: %s", clientfilePath, err)
	}
	req.Header["api-key"] = []string{c.apiKey}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s error while sending http request: %s", clientfilePath, err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s error while sending http request: %s", clientfilePath, err)
	}

	//	fmt.Println(string(respBody))
	var response *protocol.Response

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("%s error while Un marshaling http response: %s", clientfilePath, err)
	}

	return response, nil
}
