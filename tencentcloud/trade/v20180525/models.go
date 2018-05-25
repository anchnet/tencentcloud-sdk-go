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
	//"time"
)

type DescribeDealsByCond struct {
	*tchttp.BaseRequest
	Page      int    `json:"page" name:"page"`
	PageSize  int    `json:"pageSize" name:"pageSize"`
	StartTime string `json:"startTime" name:"startTime"`
	EndTime   string `json:"endTime" name:"endTime"`
	Status    int    `json:"status" name:"status"`
	DealId    int    `json:"dealId" name:"dealId"`
}

func (r *DescribeDealsByCond) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDealsByCond) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDealsByCondResponse struct {
	*tchttp.BaseResponse
	Code     int    `json:"code" name:"code"`
	Message  string `json:"message" name:"message"`
	CodeDesc string `json:"codeDesc" name:"codeDesc"`
	Data     *struct {
		TotalNum string  `json:"totalNum" name:"totalNum"`
		Deals    []*Deal `json:"deals" name:"deals" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
	} `json:"data"`
}

type Deal struct {
	Id       string `json:"id" name:"id"`
	OwnerUin string `json:"ownerUin" name:"ownerUin"`
	AppId    string `json:"appId" name:"appId"`

	DealName        string `json:"dealName" name:"dealName"`
	DealId          string `json:"dealId" name:"dealId"`
	GoodsCategoryId string `json:"goodsCategoryId" name:"goodsCategoryId"`
	GoodsDetail     struct {
		// "domain":"dnspod.cn",
		// "years":"1",
		// "service":"DR_CN_NEW"
	} `json:"goodsDetail" name:"goodsDetail"`
	CreatTime  string `json:"creatTime" name:"creatTime"`
	BillId     string `json:"billId" name:"billId"`
	TaskDetail struct {
		// "flowId":"3333"
	} `json:"taskDetail" name:"taskDetail"`
	RealTotalCost  string `json:"realTotalCost" name:"realTotalCost"`
	VoucherDecline string `json:"voucherDecline" name:"voucherDecline"`
	ProjectId      string `json:"projectId" name:"projectId"`
}

/*
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
	BillId  string `json:"payMode" name:"payMode"`
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
*/
