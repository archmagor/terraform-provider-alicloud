package cassandra

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

// DescribeIpWhitelist invokes the cassandra.DescribeIpWhitelist API synchronously
func (client *Client) DescribeIpWhitelist(request *DescribeIpWhitelistRequest) (response *DescribeIpWhitelistResponse, err error) {
	response = CreateDescribeIpWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpWhitelistWithChan invokes the cassandra.DescribeIpWhitelist API asynchronously
func (client *Client) DescribeIpWhitelistWithChan(request *DescribeIpWhitelistRequest) (<-chan *DescribeIpWhitelistResponse, <-chan error) {
	responseChan := make(chan *DescribeIpWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpWhitelist(request)
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

// DescribeIpWhitelistWithCallback invokes the cassandra.DescribeIpWhitelist API asynchronously
func (client *Client) DescribeIpWhitelistWithCallback(request *DescribeIpWhitelistRequest, callback func(response *DescribeIpWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpWhitelist(request)
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

// DescribeIpWhitelistRequest is the request struct for api DescribeIpWhitelist
type DescribeIpWhitelistRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeIpWhitelistResponse is the response struct for api DescribeIpWhitelist
type DescribeIpWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	IpList    IpListInDescribeIpWhitelist `json:"IpList" xml:"IpList"`
}

// CreateDescribeIpWhitelistRequest creates a request to invoke DescribeIpWhitelist API
func CreateDescribeIpWhitelistRequest() (request *DescribeIpWhitelistRequest) {
	request = &DescribeIpWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeIpWhitelist", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpWhitelistResponse creates a response to parse from DescribeIpWhitelist response
func CreateDescribeIpWhitelistResponse() (response *DescribeIpWhitelistResponse) {
	response = &DescribeIpWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
