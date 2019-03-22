// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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

package v20180321

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgentAuditedClient struct {

	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客审核通过时间戳
	AgentTime *string `json:"AgentTime,omitempty" name:"AgentTime"`

	// 代客类型，可能值为a/b/c
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// 代客备注
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 代客名称（首选实名认证名称）
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 认证类型, 0：个人，1：企业；其他：未认证
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 代客APPID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 上月消费金额
	LastMonthAmt *uint64 `json:"LastMonthAmt,omitempty" name:"LastMonthAmt"`

	// 本月消费金额
	ThisMonthAmt *uint64 `json:"ThisMonthAmt,omitempty" name:"ThisMonthAmt"`

	// 是否欠费,0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 客户类型：可以为new(新拓)/assign(指定)/old(存量)/空
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`
}

type AgentBillElem struct {

	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 订单号，仅对预付费账单有意义
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 云产品名称
	GoodsType *string `json:"GoodsType,omitempty" name:"GoodsType"`

	// 预付费/后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 支付月份
	SettleMonth *string `json:"SettleMonth,omitempty" name:"SettleMonth"`

	// 支付金额，单位分
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// agentpay：代付；selfpay：自付
	PayerMode *string `json:"PayerMode,omitempty" name:"PayerMode"`

	// 客户类型：可以为new(新拓)/assign(指定)/old(存量)/空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`
}

type AgentClientElem struct {

	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客申请时间戳
	ApplyTime *uint64 `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 代客类型，可能值为a/b/c
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// 代客邮箱，打码显示
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 代客手机，打码显示
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 0表示不欠费，1表示欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 1:待代理商审核;2:待腾讯云审核
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type AgentPayDealsRequest struct {
	*tchttp.BaseRequest

	// 订单所有者uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 代付标志，1：代付；0：自付
	AgentPay *uint64 `json:"AgentPay,omitempty" name:"AgentPay"`

	// 订单号数组
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

func (r *AgentPayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AgentPayDealsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AgentPayDealsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AgentPayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AgentPayDealsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AgentTransferMoneyRequest struct {
	*tchttp.BaseRequest

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 转账金额，单位分
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
}

func (r *AgentTransferMoneyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AgentTransferMoneyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AgentTransferMoneyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AgentTransferMoneyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AgentTransferMoneyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditApplyClientRequest struct {
	*tchttp.BaseRequest

	// 待审核客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult,omitempty" name:"AuditResult"`

	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *AuditApplyClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditApplyClientRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AuditApplyClientResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理商账号ID
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// 客户账号ID
		ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

		// 审核结果，包括accept/reject/qcloudaudit（腾讯云审核）
		AuditResult *string `json:"AuditResult,omitempty" name:"AuditResult"`

		// 关联时间对应的时间戳
		AgentTime *uint64 `json:"AgentTime,omitempty" name:"AgentTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuditApplyClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AuditApplyClientResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAuditedClientsRequest struct {
	*tchttp.BaseRequest

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按审核通过时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 客户账号ID列表
	ClientUins []*string `json:"ClientUins,omitempty" name:"ClientUins" list`

	// 是否欠费。0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 客户类型：可以为new(新拓)/assign(指定)/old(存量)/空
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`
}

func (r *DescribeAgentAuditedClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentAuditedClientsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentAuditedClientsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已审核代客列表
		AgentClientSet []*AgentAuditedClient `json:"AgentClientSet,omitempty" name:"AgentClientSet" list`

		// 符合条件的代客总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentAuditedClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentAuditedClientsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentBillsRequest struct {
	*tchttp.BaseRequest

	// 支付月份，如2018-02
	SettleMonth *string `json:"SettleMonth,omitempty" name:"SettleMonth"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 预付费订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgentBillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentBillsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentBillsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件列表总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 业务明细列表
		AgentBillSet []*AgentBillElem `json:"AgentBillSet,omitempty" name:"AgentBillSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentBillsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentBillsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentClientsRequest struct {
	*tchttp.BaseRequest

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgentClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentClientsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentClientsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 待审核代客列表
		AgentClientSet []*AgentClientElem `json:"AgentClientSet,omitempty" name:"AgentClientSet" list`

		// 符合条件的代客总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAgentClientsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientBalanceRequest struct {
	*tchttp.BaseRequest

	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *DescribeClientBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientBalanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientBalanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账户余额，单位分
		Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClientBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientBalanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRebateInfosRequest struct {
	*tchttp.BaseRequest

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRebateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRebateInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRebateInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返佣信息列表
		RebateInfoSet []*RebateInfoElem `json:"RebateInfoSet,omitempty" name:"RebateInfoSet" list`

		// 符合查询条件返佣信息数目
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRebateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRebateInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClientRemarkRequest struct {
	*tchttp.BaseRequest

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *ModifyClientRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClientRemarkRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClientRemarkResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClientRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClientRemarkResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebateInfoElem struct {

	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 返佣金额，单位分
	Amt *uint64 `json:"Amt,omitempty" name:"Amt"`

	// 月度业绩，单位分
	MonthSales *uint64 `json:"MonthSales,omitempty" name:"MonthSales"`

	// 季度业绩，单位分
	QuarterSales *uint64 `json:"QuarterSales,omitempty" name:"QuarterSales"`

	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag,omitempty" name:"ExceptionFlag"`
}
