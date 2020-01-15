package gozabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RPCMessage contains the request data for the Zabbix API
type RPCMessage struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
	Auth    interface{} `json:"auth"`
}

//RPCResponse contains the data to be received after a request
type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Error   interface{} `json:"error"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

//ZabbixClient contains the config for our requests
type ZabbixClient struct {
	client *http.Client
	url    string
	auth   string
}

//ZabbixAPI returns an object to do further requests with
func ZabbixAPI(host string) *ZabbixClient {
	url := host + "/api_jsonrpc.php"
	return &ZabbixClient{
		client: &http.Client{},
		url:    url,
	}
}

//Signin signs in the user
func (c *ZabbixClient) Signin(username, password string) error {
	// do a request to login
	method := "user.login"
	params := map[string]interface{}{
		"user":     username,
		"password": password,
	}
	resp, err := c.Call(method, params, false)
	if err != nil {
		return fmt.Errorf("ZabbixClient.Signin error: %s", err.Error())
	}
	// set z.auth to the proper string
	if authString, ok := resp.Result.(string); ok {
		c.auth = authString
	} else {
		return fmt.Errorf("ZabbixClient.Signin error: error reading property 'auth'")
	}
	return nil
}

//Call transmits the message and returns a ZabbixResponse
func (c *ZabbixClient) Call(method string, params interface{}, useAuth bool) (response RPCResponse, err error) {
	var auth interface{}
	if useAuth {
		auth = c.auth
	}

	// create a json-rpc message
	txMessage := RPCMessage{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Auth:    auth,
		ID:      1,
	}

	txBytes, err := json.Marshal(txMessage)
	if err != nil {
		return response, fmt.Errorf("ZabbixClient.Call json.Marshal error: %s", err.Error())
	}

	// transmit the message
	resp, err := c.client.Post(c.url, "application/json", bytes.NewReader(txBytes))
	if err != nil {
		return response, fmt.Errorf("ZabbixClient.Call POST error: %s", err.Error())
	}
	defer resp.Body.Close()

	// handle the response
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("ZabbixClient.Call reading response error: %s", err.Error())
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return response, fmt.Errorf("ZabbixClient.Call Unmarshal error: %s", err.Error())
	}

	// handle API errors
	if response.Result == nil {
		if errorData, ok := response.Error.(map[string]interface{}); ok {
			return response, fmt.Errorf("ZabbixClient.Call API Error code: %f | message: %s | data: %s", errorData["code"], errorData["message"], errorData["data"])
		}
	}

	return response, nil
}
