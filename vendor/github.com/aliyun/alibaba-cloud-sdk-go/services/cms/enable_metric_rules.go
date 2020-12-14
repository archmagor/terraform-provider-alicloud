package cms

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

// EnableMetricRules invokes the cms.EnableMetricRules API synchronously
func (client *Client) EnableMetricRules(request *EnableMetricRulesRequest) (response *EnableMetricRulesResponse, err error) {
	response = CreateEnableMetricRulesResponse()
	err = client.DoAction(request, response)
	return
}

// EnableMetricRulesWithChan invokes the cms.EnableMetricRules API asynchronously
func (client *Client) EnableMetricRulesWithChan(request *EnableMetricRulesRequest) (<-chan *EnableMetricRulesResponse, <-chan error) {
	responseChan := make(chan *EnableMetricRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableMetricRules(request)
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

// EnableMetricRulesWithCallback invokes the cms.EnableMetricRules API asynchronously
func (client *Client) EnableMetricRulesWithCallback(request *EnableMetricRulesRequest, callback func(response *EnableMetricRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableMetricRulesResponse
		var err error
		defer close(result)
		response, err = client.EnableMetricRules(request)
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

// EnableMetricRulesRequest is the request struct for api EnableMetricRules
type EnableMetricRulesRequest struct {
	*requests.RpcRequest
	RuleId *[]string `position:"Query" name:"RuleId"  type:"Repeated"`
}

// EnableMetricRulesResponse is the response struct for api EnableMetricRules
type EnableMetricRulesResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableMetricRulesRequest creates a request to invoke EnableMetricRules API
func CreateEnableMetricRulesRequest() (request *EnableMetricRulesRequest) {
	request = &EnableMetricRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "EnableMetricRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableMetricRulesResponse creates a response to parse from EnableMetricRules response
func CreateEnableMetricRulesResponse() (response *EnableMetricRulesResponse) {
	response = &EnableMetricRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
