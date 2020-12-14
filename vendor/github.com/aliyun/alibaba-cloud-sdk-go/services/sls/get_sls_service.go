package sls

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

// GetSlsService invokes the sls.GetSlsService API synchronously
func (client *Client) GetSlsService(request *GetSlsServiceRequest) (response *GetSlsServiceResponse, err error) {
	response = CreateGetSlsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSlsServiceWithChan invokes the sls.GetSlsService API asynchronously
func (client *Client) GetSlsServiceWithChan(request *GetSlsServiceRequest) (<-chan *GetSlsServiceResponse, <-chan error) {
	responseChan := make(chan *GetSlsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSlsService(request)
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

// GetSlsServiceWithCallback invokes the sls.GetSlsService API asynchronously
func (client *Client) GetSlsServiceWithCallback(request *GetSlsServiceRequest, callback func(response *GetSlsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSlsServiceResponse
		var err error
		defer close(result)
		response, err = client.GetSlsService(request)
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

// GetSlsServiceRequest is the request struct for api GetSlsService
type GetSlsServiceRequest struct {
	*requests.RpcRequest
}

// GetSlsServiceResponse is the response struct for api GetSlsService
type GetSlsServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Enabled   bool   `json:"Enabled" xml:"Enabled"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateGetSlsServiceRequest creates a request to invoke GetSlsService API
func CreateGetSlsServiceRequest() (request *GetSlsServiceRequest) {
	request = &GetSlsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "GetSlsService", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSlsServiceResponse creates a response to parse from GetSlsService response
func CreateGetSlsServiceResponse() (response *GetSlsServiceResponse) {
	response = &GetSlsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
