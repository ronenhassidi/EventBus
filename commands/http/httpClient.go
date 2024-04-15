package commands

// The main need to be identical for controller and agent ad copied to both directories /controller and /agent

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/gilbarco-ai/event-bus/common/log"
	config "github.com/gilbarco-ai/event-bus/configuration"
)

func HttpClient(apiMsg config.ApiMsg) ([]byte, error) {
	// Create an HTTP client
	client := &http.Client{}

	body := GetBody(apiMsg.Body)
	req, err := GetRequest(apiMsg, body)
	if err != nil {
		log.Logger.Error("Error creating HTTP request:", err)
		return nil, err
	}
	req = AddHeaders(apiMsg.Headers, req)
	req = AddParameters(apiMsg.Params, req)

	resp, err := client.Do(req)
	if err != nil {
		log.Logger.Error("Error making HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Error("Error reading HTTP response:", err)
		return nil, err
	}
	res := string(response)
	log.Logger.Info("Respons:", res)

	//	jsonData, err := json.Marshal(res)
	//	if err != nil {
	//		logger.Error.Fatal("Error reading HTTP response:", err)
	//		return nil, err
	//	}
	//	logger.Info.Println("Respons after :", jsonData)

	//var data config.ApiResponse
	//json.Unmarshal(response, &data)
	//logger.Info.Println("Respons data:", data)

	return response, nil

}

func AddParameters(keyValue []config.KeyValue, req *http.Request) *http.Request {
	params := url.Values{}

	if keyValue != nil {
		for i := 0; i < len(keyValue); i++ {
			params.Add(keyValue[i].Key, keyValue[i].Value)
		}
		req.URL.RawQuery = params.Encode()
		return req
	}

	return req
}

func GetRequest(apiMsg config.ApiMsg, body *bytes.Buffer) (*http.Request, error) {

	req, err := http.NewRequest(apiMsg.MethodType, apiMsg.Url, body)
	if err != nil {
		log.Logger.Error("Error creating HTTP request:", err)
		return req, err
	}

	return req, err
}

func AddHeaders(headerInput []config.KeyValue, req *http.Request) *http.Request {

	if req != nil {
		for i := 0; i < len(headerInput); i++ {
			req.Header.Add(headerInput[i].Key, headerInput[i].Value)
		}
	}

	return req
}

func GetBody(mapStr map[string]interface{}) *bytes.Buffer {

	//bodyStr := ""

	byteData, err := json.Marshal(mapStr)
	if err != nil {
		log.Logger.Error(err)
	}

	var buffer bytes.Buffer
	_, err = buffer.Write(byteData)
	if err != nil {
		log.Logger.Error(err)
	}
	//body := bytes.NewBufferString(jsonData)

	return &buffer
}
