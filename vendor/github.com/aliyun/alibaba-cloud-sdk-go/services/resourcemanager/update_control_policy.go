package resourcemanager

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

// UpdateControlPolicy invokes the resourcemanager.UpdateControlPolicy API synchronously
func (client *Client) UpdateControlPolicy(request *UpdateControlPolicyRequest) (response *UpdateControlPolicyResponse, err error) {
	response = CreateUpdateControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateControlPolicyWithChan invokes the resourcemanager.UpdateControlPolicy API asynchronously
func (client *Client) UpdateControlPolicyWithChan(request *UpdateControlPolicyRequest) (<-chan *UpdateControlPolicyResponse, <-chan error) {
	responseChan := make(chan *UpdateControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateControlPolicy(request)
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

// UpdateControlPolicyWithCallback invokes the resourcemanager.UpdateControlPolicy API asynchronously
func (client *Client) UpdateControlPolicyWithCallback(request *UpdateControlPolicyRequest, callback func(response *UpdateControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.UpdateControlPolicy(request)
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

// UpdateControlPolicyRequest is the request struct for api UpdateControlPolicy
type UpdateControlPolicyRequest struct {
	*requests.RpcRequest
	NewPolicyName     string `position:"Query" name:"NewPolicyName"`
	PolicyId          string `position:"Query" name:"PolicyId"`
	NewPolicyDocument string `position:"Query" name:"NewPolicyDocument"`
	NewDescription    string `position:"Query" name:"NewDescription"`
}

// UpdateControlPolicyResponse is the response struct for api UpdateControlPolicy
type UpdateControlPolicyResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ControlPolicy ControlPolicy `json:"ControlPolicy" xml:"ControlPolicy"`
}

// CreateUpdateControlPolicyRequest creates a request to invoke UpdateControlPolicy API
func CreateUpdateControlPolicyRequest() (request *UpdateControlPolicyRequest) {
	request = &UpdateControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "UpdateControlPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateControlPolicyResponse creates a response to parse from UpdateControlPolicy response
func CreateUpdateControlPolicyResponse() (response *UpdateControlPolicyResponse) {
	response = &UpdateControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
