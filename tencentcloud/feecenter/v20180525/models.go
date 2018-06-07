// Copyright 1999-2018 Tencent Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20170312

import (
	"encoding/json"

	tchttp "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeResourceBillsRequest struct {
	*tchttp.BaseRequest
	Platform   string `json:"platform" name:"platform"`
	PayMode    string `json:"payMode" name:"payMode"`
	PayType    string `json:"payType" name:"payType"`
	StartMonth string `json:"startMonth" name:"startMonth"`
	EndMonth   string `json:"endMonth" name:"endMonth"`
	Order      string `json:"order" name:"order"`
}

func (r *DescribeResourceBillsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceBillsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceBillsResponse struct {
	*tchttp.BaseResponse
	Code     int    `json:"code" name:"code"`
	Message  string `json:"message" name:"message"`
	CodeDesc string `json:"codeDesc" name:"codeDesc"`
	Data     *struct {
		Total           int      `json:"total" name:"total"`
		Amount          int      `json:"amount" name:"amount"`
		ProductCodeList []string `json:"productCodeList" name:"productCodeList" list`

		Data []*ResourceBill `json:"data" name:"data" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	} `json:"data"`
}

type ResourceBill struct {
	Id             string      `json:"id" name:"id"`
	OwnerUin       string      `json:"ownerUin" name:"ownerUin"`
	AppId          string      `json:"appId" name:"appId"`
	PayerMode      string      `json:"payerMode" name:"payerMode"`
	PayerUin       string      `json:"payerUin" name:"payerUin"`
	PayMode        string      `json:"payMode" name:"payMode"`
	PayType        string      `json:"payType" name:"payType"`
	Platform       string      `json:"platform" name:"platform"`
	Region         string      `json:"region" name:"region"`
	ZoneNumber     string      `json:"zoneNumber" name:"zoneNumber"`
	ProductCode    string      `json:"productCode" name:"productCode"`
	SubProductCode string      `json:"subProductCode" name:"subProductCode"`
	InstanceIdType string      `json:"instanceIdType" name:"instanceIdType"`
	InstanceId     string      `json:"instanceId" name:"instanceId"`
	ActionType     string      `json:"actionType" name:"actionType"`
	CalcUnit       string      `json:"calcUnit" name:"calcUnit"`
	TimeSpan       string      `json:"timeSpan" name:"timeSpan"`
	TimeUnit       string      `json:"timeUnit" name:"timeUnit"`
	UnitPrice      string      `json:"unitPrice" name:"unitPrice"`
	Price          string      `json:"price" name:"price"`
	Count          string      `json:"count" name:"count"`
	Month          string      `json:"month" name:"month"`
	StartTime      string      `json:"startTime" name:"startTime"`
	EndTime        string      `json:"endTime" name:"endTime"`
	Amount         string      `json:"amount" name:"amount"`
	Reduces        interface{} `json:"reduces" name:"reduces"`
	ProjectId      string      `json:"projectId" name:"projectId"`
	Projects       interface{} `json:"projects" name:"projects"`
	UsedTimeSpan   string      `json:"usedTimeSpan" name:"usedTimeSpan"`
	ReturnAmount   string      `json:"returnAmount" name:"returnAmount"`
	ModifyType     string      `json:"modifyType" name:"modifyType"`
}

type DescribeBillsRequest struct {
	*tchttp.BaseRequest
	PayerUin  string `json:"payerUin" name:"payerUin"`
	PayType   string `json:"payType" name:"payType"`
	StartTime string `json:"startTime" name:"startTime"`
	EndTime   string `json:"endTime" name:"endTime"`
	Order     string `json:"order" name:"order"`
}

func (r *DescribeBillsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBillsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillsResponse struct {
	*tchttp.BaseResponse
	Code     int    `json:"code" name:"code"`
	Message  string `json:"message" name:"message"`
	CodeDesc string `json:"codeDesc" name:"codeDesc"`
	Data     *struct {
		Total int `json:"total" name:"total"`
		In    int `json:"in" name:"in"`
		Out   int `json:"out" name:"productCodeList" list`

		Data []*Bill `json:"data" name:"data" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	} `json:"data"`
}

type Bill struct {
	FeeId   string `json:"feeId" name:"feeId"`
	BillId  string `json:"billId" name:"billId"`
	Status  string `json:"status" name:"status"`
	PayMode string `json:"payMode" name:"payMode"`
	PayType string `json:"payType" name:"payType"`

	OperationTime string `json:"operationTime" name:"operationTime"`
	OperationInfo string `json:"operationInfo" name:"operationInfo"`
	Amount        string `json:"amount" name:"amount"`
	Balance       string `json:"balance" name:"balance"`

	CalcUnit       string `json:"calcUnit" name:"calcUnit"`
	ActionType     string `json:"actionType" name:"actionType"`
	ProductCode    string `json:"productCode" name:"productCode"`
	SubProductCode string `json:"subProductCode" name:"subProductCode"`
}
