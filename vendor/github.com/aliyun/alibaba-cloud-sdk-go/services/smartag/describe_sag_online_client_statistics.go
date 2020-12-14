package smartag

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

// DescribeSagOnlineClientStatistics invokes the smartag.DescribeSagOnlineClientStatistics API synchronously
func (client *Client) DescribeSagOnlineClientStatistics(request *DescribeSagOnlineClientStatisticsRequest) (response *DescribeSagOnlineClientStatisticsResponse, err error) {
	response = CreateDescribeSagOnlineClientStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagOnlineClientStatisticsWithChan invokes the smartag.DescribeSagOnlineClientStatistics API asynchronously
func (client *Client) DescribeSagOnlineClientStatisticsWithChan(request *DescribeSagOnlineClientStatisticsRequest) (<-chan *DescribeSagOnlineClientStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeSagOnlineClientStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagOnlineClientStatistics(request)
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

// DescribeSagOnlineClientStatisticsWithCallback invokes the smartag.DescribeSagOnlineClientStatistics API asynchronously
func (client *Client) DescribeSagOnlineClientStatisticsWithCallback(request *DescribeSagOnlineClientStatisticsRequest, callback func(response *DescribeSagOnlineClientStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagOnlineClientStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagOnlineClientStatistics(request)
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

// DescribeSagOnlineClientStatisticsRequest is the request struct for api DescribeSagOnlineClientStatistics
type DescribeSagOnlineClientStatisticsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SmartAGIds           *[]string        `position:"Query" name:"SmartAGIds"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSagOnlineClientStatisticsResponse is the response struct for api DescribeSagOnlineClientStatistics
type DescribeSagOnlineClientStatisticsResponse struct {
	*responses.BaseResponse
	RequestId     string                                           `json:"RequestId" xml:"RequestId"`
	SagStatistics SagStatisticsInDescribeSagOnlineClientStatistics `json:"SagStatistics" xml:"SagStatistics"`
}

// CreateDescribeSagOnlineClientStatisticsRequest creates a request to invoke DescribeSagOnlineClientStatistics API
func CreateDescribeSagOnlineClientStatisticsRequest() (request *DescribeSagOnlineClientStatisticsRequest) {
	request = &DescribeSagOnlineClientStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagOnlineClientStatistics", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagOnlineClientStatisticsResponse creates a response to parse from DescribeSagOnlineClientStatistics response
func CreateDescribeSagOnlineClientStatisticsResponse() (response *DescribeSagOnlineClientStatisticsResponse) {
	response = &DescribeSagOnlineClientStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
