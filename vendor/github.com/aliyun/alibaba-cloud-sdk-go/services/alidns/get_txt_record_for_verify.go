package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetTxtRecordForVerify invokes the alidns.GetTxtRecordForVerify API synchronously
func (client *Client) GetTxtRecordForVerify(request *GetTxtRecordForVerifyRequest) (response *GetTxtRecordForVerifyResponse, err error) {
	response = CreateGetTxtRecordForVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// GetTxtRecordForVerifyWithChan invokes the alidns.GetTxtRecordForVerify API asynchronously
func (client *Client) GetTxtRecordForVerifyWithChan(request *GetTxtRecordForVerifyRequest) (<-chan *GetTxtRecordForVerifyResponse, <-chan error) {
	responseChan := make(chan *GetTxtRecordForVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTxtRecordForVerify(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetTxtRecordForVerifyWithCallback invokes the alidns.GetTxtRecordForVerify API asynchronously
func (client *Client) GetTxtRecordForVerifyWithCallback(request *GetTxtRecordForVerifyRequest, callback func(response *GetTxtRecordForVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTxtRecordForVerifyResponse
		var err error
		defer close(result)
		response, err = client.GetTxtRecordForVerify(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetTxtRecordForVerifyRequest is the request struct for api GetTxtRecordForVerify
type GetTxtRecordForVerifyRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	Type         string `position:"Query" name:"Type"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// GetTxtRecordForVerifyResponse is the response struct for api GetTxtRecordForVerify
type GetTxtRecordForVerifyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainName string `json:"DomainName" xml:"DomainName"`
	RR         string `json:"RR" xml:"RR"`
	Value      string `json:"Value" xml:"Value"`
}

// CreateGetTxtRecordForVerifyRequest creates a request to invoke GetTxtRecordForVerify API
func CreateGetTxtRecordForVerifyRequest() (request *GetTxtRecordForVerifyRequest) {
	request = &GetTxtRecordForVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "GetTxtRecordForVerify", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTxtRecordForVerifyResponse creates a response to parse from GetTxtRecordForVerify response
func CreateGetTxtRecordForVerifyResponse() (response *GetTxtRecordForVerifyResponse) {
	response = &GetTxtRecordForVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
