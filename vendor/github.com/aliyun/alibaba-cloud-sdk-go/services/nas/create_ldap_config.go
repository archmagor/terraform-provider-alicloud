package nas

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

// CreateLDAPConfig invokes the nas.CreateLDAPConfig API synchronously
func (client *Client) CreateLDAPConfig(request *CreateLDAPConfigRequest) (response *CreateLDAPConfigResponse, err error) {
	response = CreateCreateLDAPConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLDAPConfigWithChan invokes the nas.CreateLDAPConfig API asynchronously
func (client *Client) CreateLDAPConfigWithChan(request *CreateLDAPConfigRequest) (<-chan *CreateLDAPConfigResponse, <-chan error) {
	responseChan := make(chan *CreateLDAPConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLDAPConfig(request)
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

// CreateLDAPConfigWithCallback invokes the nas.CreateLDAPConfig API asynchronously
func (client *Client) CreateLDAPConfigWithCallback(request *CreateLDAPConfigRequest, callback func(response *CreateLDAPConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLDAPConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateLDAPConfig(request)
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

// CreateLDAPConfigRequest is the request struct for api CreateLDAPConfig
type CreateLDAPConfigRequest struct {
	*requests.RpcRequest
	SearchBase   string `position:"Query" name:"SearchBase"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
	URI          string `position:"Query" name:"URI"`
	BindDN       string `position:"Query" name:"BindDN"`
}

// CreateLDAPConfigResponse is the response struct for api CreateLDAPConfig
type CreateLDAPConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLDAPConfigRequest creates a request to invoke CreateLDAPConfig API
func CreateCreateLDAPConfigRequest() (request *CreateLDAPConfigRequest) {
	request = &CreateLDAPConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateLDAPConfig", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLDAPConfigResponse creates a response to parse from CreateLDAPConfig response
func CreateCreateLDAPConfigResponse() (response *CreateLDAPConfigResponse) {
	response = &CreateLDAPConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
