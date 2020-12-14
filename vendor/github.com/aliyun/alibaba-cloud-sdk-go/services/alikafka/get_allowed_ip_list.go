package alikafka

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

// GetAllowedIpList invokes the alikafka.GetAllowedIpList API synchronously
func (client *Client) GetAllowedIpList(request *GetAllowedIpListRequest) (response *GetAllowedIpListResponse, err error) {
	response = CreateGetAllowedIpListResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllowedIpListWithChan invokes the alikafka.GetAllowedIpList API asynchronously
func (client *Client) GetAllowedIpListWithChan(request *GetAllowedIpListRequest) (<-chan *GetAllowedIpListResponse, <-chan error) {
	responseChan := make(chan *GetAllowedIpListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllowedIpList(request)
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

// GetAllowedIpListWithCallback invokes the alikafka.GetAllowedIpList API asynchronously
func (client *Client) GetAllowedIpListWithCallback(request *GetAllowedIpListRequest, callback func(response *GetAllowedIpListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllowedIpListResponse
		var err error
		defer close(result)
		response, err = client.GetAllowedIpList(request)
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

// GetAllowedIpListRequest is the request struct for api GetAllowedIpList
type GetAllowedIpListRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetAllowedIpListResponse is the response struct for api GetAllowedIpList
type GetAllowedIpListResponse struct {
	*responses.BaseResponse
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        int         `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	AllowedList AllowedList `json:"AllowedList" xml:"AllowedList"`
}

// CreateGetAllowedIpListRequest creates a request to invoke GetAllowedIpList API
func CreateGetAllowedIpListRequest() (request *GetAllowedIpListRequest) {
	request = &GetAllowedIpListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetAllowedIpList", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAllowedIpListResponse creates a response to parse from GetAllowedIpList response
func CreateGetAllowedIpListResponse() (response *GetAllowedIpListResponse) {
	response = &GetAllowedIpListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
