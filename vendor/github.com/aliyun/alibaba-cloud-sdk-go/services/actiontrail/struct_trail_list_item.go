package actiontrail

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

// TrailListItem is a nested struct in actiontrail response
type TrailListItem struct {
	Name                string `json:"Name" xml:"Name"`
	HomeRegion          string `json:"HomeRegion" xml:"HomeRegion"`
	RoleName            string `json:"RoleName" xml:"RoleName"`
	OssBucketName       string `json:"OssBucketName" xml:"OssBucketName"`
	OssKeyPrefix        string `json:"OssKeyPrefix" xml:"OssKeyPrefix"`
	EventRW             string `json:"EventRW" xml:"EventRW"`
	SlsWriteRoleArn     string `json:"SlsWriteRoleArn" xml:"SlsWriteRoleArn"`
	SlsProjectArn       string `json:"SlsProjectArn" xml:"SlsProjectArn"`
	Status              string `json:"Status" xml:"Status"`
	TrailRegion         string `json:"TrailRegion" xml:"TrailRegion"`
	CreateTime          string `json:"CreateTime" xml:"CreateTime"`
	UpdateTime          string `json:"UpdateTime" xml:"UpdateTime"`
	StartLoggingTime    string `json:"StartLoggingTime" xml:"StartLoggingTime"`
	StopLoggingTime     string `json:"StopLoggingTime" xml:"StopLoggingTime"`
	MnsTopicArn         string `json:"MnsTopicArn" xml:"MnsTopicArn"`
	IsOrganizationTrail bool   `json:"IsOrganizationTrail" xml:"IsOrganizationTrail"`
}
