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

// CreateSaslUser invokes the alikafka.CreateSaslUser API synchronously
func (client *Client) CreateSaslUser(request *CreateSaslUserRequest) (response *CreateSaslUserResponse, err error) {
	response = CreateCreateSaslUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSaslUserWithChan invokes the alikafka.CreateSaslUser API asynchronously
func (client *Client) CreateSaslUserWithChan(request *CreateSaslUserRequest) (<-chan *CreateSaslUserResponse, <-chan error) {
	responseChan := make(chan *CreateSaslUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSaslUser(request)
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

// CreateSaslUserWithCallback invokes the alikafka.CreateSaslUser API asynchronously
func (client *Client) CreateSaslUserWithCallback(request *CreateSaslUserRequest, callback func(response *CreateSaslUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSaslUserResponse
		var err error
		defer close(result)
		response, err = client.CreateSaslUser(request)
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

// CreateSaslUserRequest is the request struct for api CreateSaslUser
type CreateSaslUserRequest struct {
	*requests.RpcRequest
	Type       string `position:"Query" name:"Type"`
	Password   string `position:"Query" name:"Password"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Username   string `position:"Query" name:"Username"`
}

// CreateSaslUserResponse is the response struct for api CreateSaslUser
type CreateSaslUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateSaslUserRequest creates a request to invoke CreateSaslUser API
func CreateCreateSaslUserRequest() (request *CreateSaslUserRequest) {
	request = &CreateSaslUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "CreateSaslUser", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSaslUserResponse creates a response to parse from CreateSaslUser response
func CreateCreateSaslUserResponse() (response *CreateSaslUserResponse) {
	response = &CreateSaslUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
