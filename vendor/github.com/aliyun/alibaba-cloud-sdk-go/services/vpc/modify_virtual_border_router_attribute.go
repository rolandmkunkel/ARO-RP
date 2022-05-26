package vpc

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

// ModifyVirtualBorderRouterAttribute invokes the vpc.ModifyVirtualBorderRouterAttribute API synchronously
func (client *Client) ModifyVirtualBorderRouterAttribute(request *ModifyVirtualBorderRouterAttributeRequest) (response *ModifyVirtualBorderRouterAttributeResponse, err error) {
	response = CreateModifyVirtualBorderRouterAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVirtualBorderRouterAttributeWithChan invokes the vpc.ModifyVirtualBorderRouterAttribute API asynchronously
func (client *Client) ModifyVirtualBorderRouterAttributeWithChan(request *ModifyVirtualBorderRouterAttributeRequest) (<-chan *ModifyVirtualBorderRouterAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVirtualBorderRouterAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVirtualBorderRouterAttribute(request)
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

// ModifyVirtualBorderRouterAttributeWithCallback invokes the vpc.ModifyVirtualBorderRouterAttribute API asynchronously
func (client *Client) ModifyVirtualBorderRouterAttributeWithCallback(request *ModifyVirtualBorderRouterAttributeRequest, callback func(response *ModifyVirtualBorderRouterAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVirtualBorderRouterAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVirtualBorderRouterAttribute(request)
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

// ModifyVirtualBorderRouterAttributeRequest is the request struct for api ModifyVirtualBorderRouterAttribute
type ModifyVirtualBorderRouterAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CircuitCode                   string           `position:"Query" name:"CircuitCode"`
	AssociatedPhysicalConnections string           `position:"Query" name:"AssociatedPhysicalConnections"`
	VlanId                        requests.Integer `position:"Query" name:"VlanId"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	EnableIpv6                    requests.Boolean `position:"Query" name:"EnableIpv6"`
	Description                   string           `position:"Query" name:"Description"`
	VbrId                         string           `position:"Query" name:"VbrId"`
	PeerGatewayIp                 string           `position:"Query" name:"PeerGatewayIp"`
	PeerIpv6GatewayIp             string           `position:"Query" name:"PeerIpv6GatewayIp"`
	DetectMultiplier              requests.Integer `position:"Query" name:"DetectMultiplier"`
	PeeringSubnetMask             string           `position:"Query" name:"PeeringSubnetMask"`
	LocalGatewayIp                string           `position:"Query" name:"LocalGatewayIp"`
	MinTxInterval                 requests.Integer `position:"Query" name:"MinTxInterval"`
	PeeringIpv6SubnetMask         string           `position:"Query" name:"PeeringIpv6SubnetMask"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                     requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	MinRxInterval                 requests.Integer `position:"Query" name:"MinRxInterval"`
	LocalIpv6GatewayIp            string           `position:"Query" name:"LocalIpv6GatewayIp"`
	Name                          string           `position:"Query" name:"Name"`
}

// ModifyVirtualBorderRouterAttributeResponse is the response struct for api ModifyVirtualBorderRouterAttribute
type ModifyVirtualBorderRouterAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVirtualBorderRouterAttributeRequest creates a request to invoke ModifyVirtualBorderRouterAttribute API
func CreateModifyVirtualBorderRouterAttributeRequest() (request *ModifyVirtualBorderRouterAttributeRequest) {
	request = &ModifyVirtualBorderRouterAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVirtualBorderRouterAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVirtualBorderRouterAttributeResponse creates a response to parse from ModifyVirtualBorderRouterAttribute response
func CreateModifyVirtualBorderRouterAttributeResponse() (response *ModifyVirtualBorderRouterAttributeResponse) {
	response = &ModifyVirtualBorderRouterAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
