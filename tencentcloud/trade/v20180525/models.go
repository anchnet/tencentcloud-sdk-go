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

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Code        int    `json:"code" name:"code"`
	Message     string `json:"message" name:"message"`
	BalanceInfo int    `json:"balanceInfo" name:"balanceInfo"`
}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountBalanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountBalanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Code     int    `json:"code" name:"code"`
	Message  string `json:"message" name:"message"`
	UserInfo *struct {
		Name       string `json:"name" name:"name"`
		IsOwner    int    `json:"isOwner" name:"isOwner"`
		MailStatus int    `json:"mailStatus" name:"mailStatus"`
		Mail       string `json:"mail" name:"mail"`
		Phone      string `json:"phone" name:"phone"`
	} `json:"userInfo"`
}
