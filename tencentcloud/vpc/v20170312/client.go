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
    "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

type Client struct {
    common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithSecretId(credential.SecretId, credential.SecretKey).
        WithProfile(clientProfile)
    return
}


func NewAllocateAddressesRequest() (request *AllocateAddressesRequest) {
    request = &AllocateAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AllocateAddresses")
    return
}

func NewAllocateAddressesResponse() (response *AllocateAddressesResponse) {
    response = &AllocateAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (AllocateAddresses) 用于申请一个或多个[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）。
// * EIP 是专为动态云计算设计的静态 IP 地址。借助 EIP，您可以快速将 EIP 重新映射到您的另一个实例上，从而屏蔽实例故障。
// * 您的 EIP 与腾讯云账户相关联，而不是与某个实例相关联。在您选择显式释放该地址，或欠费超过七天之前，它会一直与您的腾讯云账户保持关联。
// * 平台对用户每地域能申请的 EIP 最大配额有所限制，可参见 [EIP 产品简介](https://cloud.tencent.com/document/product/213/5733)，上述配额可通过 DescribeAddressQuota 接口获取。
func (c *Client) AllocateAddresses(request *AllocateAddressesRequest) (response *AllocateAddressesResponse, err error) {
    if request == nil {
        request = NewAllocateAddressesRequest()
    }
    response = NewAllocateAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssignPrivateIpAddressesRequest() (request *AssignPrivateIpAddressesRequest) {
    request = &AssignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssignPrivateIpAddresses")
    return
}

func NewAssignPrivateIpAddressesResponse() (response *AssignPrivateIpAddressesResponse) {
    response = &AssignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AssignPrivateIpAddresses）用于弹性网卡申请内网 IP。
// * 一个弹性网卡支持绑定的IP地址是有限制的，更多资源限制信息详见<a href="https://cloud.tencent.com/document/product/215/6513">弹性网卡使用限制</a>。
// * 可以指定内网IP地址申请，内网IP地址类型不能为主IP，主IP已存在，不能修改，内网IP必须要弹性网卡所在子网内，而且不能被占用。
// * 在弹性网卡上申请一个到多个辅助内网IP，接口会在弹性网卡所在子网网段内返回指定数量的辅助内网IP。
func (c *Client) AssignPrivateIpAddresses(request *AssignPrivateIpAddressesRequest) (response *AssignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewAssignPrivateIpAddressesRequest()
    }
    response = NewAssignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateAddressRequest() (request *AssociateAddressRequest) {
    request = &AssociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AssociateAddress")
    return
}

func NewAssociateAddressResponse() (response *AssociateAddressResponse) {
    response = &AssociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (AssociateAddress) 用于将[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）绑定到实例或弹性网卡的指定内网 IP 上。
// * 将 EIP 绑定到实例上，其本质是将 EIP 绑定到实例上主网卡的主内网 IP 上。
// * 将 EIP 绑定到主网卡的主内网IP上，绑定过程会把其上绑定的普通公网 IP 自动解绑并释放。
// * 如果指定网卡的内网 IP 已经绑定了 EIP，则必须先解绑该 EIP，才能再绑定新的。
// * EIP 如果欠费或被封堵，则不能被绑定。
// * 只有状态为 UNBIND 的 EIP 才能够被绑定。
func (c *Client) AssociateAddress(request *AssociateAddressRequest) (response *AssociateAddressResponse, err error) {
    if request == nil {
        request = NewAssociateAddressRequest()
    }
    response = NewAssociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAttachClassicLinkVpcRequest() (request *AttachClassicLinkVpcRequest) {
    request = &AttachClassicLinkVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AttachClassicLinkVpc")
    return
}

func NewAttachClassicLinkVpcResponse() (response *AttachClassicLinkVpcResponse) {
    response = &AttachClassicLinkVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(AttachClassicLinkVpc)用于私有网络和基础网络设备互通。
// * 私有网络和基础网络设备必须在同一个地域。
// * 私有网络和基础网络的区别详见vpc产品文档-<a href="https://cloud.tencent.com/document/product/215/535#2.-.E7.A7.81.E6.9C.89.E7.BD.91.E7.BB.9C.E4.B8.8E.E5.9F.BA.E7.A1.80.E7.BD.91.E7.BB.9C">私有网络与基础网络</a>。
func (c *Client) AttachClassicLinkVpc(request *AttachClassicLinkVpcRequest) (response *AttachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewAttachClassicLinkVpcRequest()
    }
    response = NewAttachClassicLinkVpcResponse()
    err = c.Send(request, response)
    return
}

func NewAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
    request = &AttachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "AttachNetworkInterface")
    return
}

func NewAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
    response = &AttachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachNetworkInterface）用于弹性网卡绑定云主机。
// * 一个云主机可以绑定多个弹性网卡，但只能绑定一个主网卡。更多限制信息详见<a href="https://cloud.tencent.com/document/product/215/6513">弹性网卡使用限制</a>。
// * 一个弹性网卡只能同时绑定一个云主机。
// * 只有运行中或者已关机状态的云主机才能绑定弹性网卡，查看云主机状态详见<a href="https://cloud.tencent.com/document/api/213/9452#instance_state">腾讯云主机信息</a>。
// * 弹性网卡绑定的云主机必须是私有网络的，而且云主机所在可用区必须和弹性网卡子网的可用区相同。
func (c *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewAttachNetworkInterfaceRequest()
    }
    response = NewAttachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressTemplateRequest() (request *CreateAddressTemplateRequest) {
    request = &CreateAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplate")
    return
}

func NewCreateAddressTemplateResponse() (response *CreateAddressTemplateResponse) {
    response = &CreateAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建IP地址模版
func (c *Client) CreateAddressTemplate(request *CreateAddressTemplateRequest) (response *CreateAddressTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateRequest()
    }
    response = NewCreateAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAddressTemplateGroupRequest() (request *CreateAddressTemplateGroupRequest) {
    request = &CreateAddressTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateAddressTemplateGroup")
    return
}

func NewCreateAddressTemplateGroupResponse() (response *CreateAddressTemplateGroupResponse) {
    response = &CreateAddressTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建IP地址模版集合
func (c *Client) CreateAddressTemplateGroup(request *CreateAddressTemplateGroupRequest) (response *CreateAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateAddressTemplateGroupRequest()
    }
    response = NewCreateAddressTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetworkInterfaceRequest() (request *CreateNetworkInterfaceRequest) {
    request = &CreateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateNetworkInterface")
    return
}

func NewCreateNetworkInterfaceResponse() (response *CreateNetworkInterfaceResponse) {
    response = &CreateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateNetworkInterface）用于创建弹性网卡。
// * 创建弹性网卡时可以指定内网IP，并且可以指定一个主IP，指定的内网IP必须在弹性网卡所在子网内，而且不能被占用。
// * 创建弹性网卡时可以指定需要申请的内网IP数量，系统会随机生成内网IP地址。
// * 创建弹性网卡同时可以绑定已有安全组。
func (c *Client) CreateNetworkInterface(request *CreateNetworkInterfaceRequest) (response *CreateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewCreateNetworkInterfaceRequest()
    }
    response = NewCreateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteTableRequest() (request *CreateRouteTableRequest) {
    request = &CreateRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRouteTable")
    return
}

func NewCreateRouteTableResponse() (response *CreateRouteTableResponse) {
    response = &CreateRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateRouteTable)用于创建路由表。
// * 创建了VPC后，系统会创建一个默认路由表，所有新建的子网都会关联到默认路由表。默认情况下您可以直接使用默认路由表来管理您的路由策略。当您的路由策略较多时，您可以调用创建路由表接口创建更多路由表管理您的路由策略。
func (c *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
    if request == nil {
        request = NewCreateRouteTableRequest()
    }
    response = NewCreateRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoutesRequest() (request *CreateRoutesRequest) {
    request = &CreateRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateRoutes")
    return
}

func NewCreateRoutesResponse() (response *CreateRoutesResponse) {
    response = &CreateRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateRoutes)用于创建路由策略。
// * 向指定路由表批量新增路由策略。
func (c *Client) CreateRoutes(request *CreateRoutesRequest) (response *CreateRoutesResponse, err error) {
    if request == nil {
        request = NewCreateRoutesRequest()
    }
    response = NewCreateRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupRequest() (request *CreateSecurityGroupRequest) {
    request = &CreateSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroup")
    return
}

func NewCreateSecurityGroupResponse() (response *CreateSecurityGroupResponse) {
    response = &CreateSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateSecurityGroup）用于创建新的安全组（SecurityGroup）。
// * 每个账户下每个地域的每个项目的<a href="https://cloud.tencent.com/document/product/213/500#2.-.E5.AE.89.E5.85.A8.E7.BB.84.E7.9A.84.E9.99.90.E5.88.B6">安全组数量限制</a>。
// * 新建的安全组的入站和出站规则默认都是全部拒绝，在创建后通常您需要再调用CreateSecurityGroupPolicies将安全组的规则设置为需要的规则。
func (c *Client) CreateSecurityGroup(request *CreateSecurityGroupRequest) (response *CreateSecurityGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupRequest()
    }
    response = NewCreateSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityGroupPoliciesRequest() (request *CreateSecurityGroupPoliciesRequest) {
    request = &CreateSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSecurityGroupPolicies")
    return
}

func NewCreateSecurityGroupPoliciesResponse() (response *CreateSecurityGroupPoliciesResponse) {
    response = &CreateSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateSecurityGroupPolicies）用于创建安全组规则（SecurityGroupPolicy）。
// 
// * Version安全组规则版本号，用户每次更新安全规则版本会自动加1，防止你更新的路由规则已过期，不填不考虑冲突。
// * Protocol字段支持输入TCP, UDP, ICMP, GRE, ALL。
// * CidrBlock字段允许输入符合cidr格式标准的任意字符串。(展开)在基础网络中，如果CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * SecurityGroupId字段允许输入与待修改的安全组位于相同项目中的安全组ID，包括这个安全组ID本身，代表安全组下所有云服务器的内网IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个ID所关联的云服务器变化而变化，不需要重新修改。
// * Port字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当Protocol字段是TCP或UDP时，Port字段才被接受。
// * Action字段只允许输入ACCEPT或DROP。
// * CidrBlock, SecurityGroupId, AddressTemplate三者是排他关系，不允许同时输入，Protocol + Port和ServiceTemplate二者是排他关系，不允许同时输入。
// * 一次请求中只能创建单个方向的规则, 如果需要指定索引（PolicyIndex）参数, 多条规则的索引必须一致。
func (c *Client) CreateSecurityGroupPolicies(request *CreateSecurityGroupPoliciesRequest) (response *CreateSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewCreateSecurityGroupPoliciesRequest()
    }
    response = NewCreateSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceTemplateRequest() (request *CreateServiceTemplateRequest) {
    request = &CreateServiceTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplate")
    return
}

func NewCreateServiceTemplateResponse() (response *CreateServiceTemplateResponse) {
    response = &CreateServiceTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建协议端口模板
func (c *Client) CreateServiceTemplate(request *CreateServiceTemplateRequest) (response *CreateServiceTemplateResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateRequest()
    }
    response = NewCreateServiceTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceTemplateGroupRequest() (request *CreateServiceTemplateGroupRequest) {
    request = &CreateServiceTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateServiceTemplateGroup")
    return
}

func NewCreateServiceTemplateGroupResponse() (response *CreateServiceTemplateGroupResponse) {
    response = &CreateServiceTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建协议端口模板集合
func (c *Client) CreateServiceTemplateGroup(request *CreateServiceTemplateGroupRequest) (response *CreateServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewCreateServiceTemplateGroupRequest()
    }
    response = NewCreateServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubnetRequest() (request *CreateSubnetRequest) {
    request = &CreateSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateSubnet")
    return
}

func NewCreateSubnetResponse() (response *CreateSubnetResponse) {
    response = &CreateSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateSubnet)用于创建子网。
// * 创建子网前必须创建好 VPC。
// * 子网创建成功后，子网网段不能修改。子网网段必须在VPC网段内，可以和VPC网段相同（VPC有且只有一个子网时），建议子网网段在VPC网段内，预留网段给其他子网使用。
// * 你可以创建的最小网段子网掩码为28（有16个IP地址），最大网段子网掩码为16（65,536个IP地址）。
// * 同一个VPC内，多个子网的网段不能重叠。
// * 子网创建后会自动关联到默认路由表。
func (c *Client) CreateSubnet(request *CreateSubnetRequest) (response *CreateSubnetResponse, err error) {
    if request == nil {
        request = NewCreateSubnetRequest()
    }
    response = NewCreateSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVpcRequest() (request *CreateVpcRequest) {
    request = &CreateVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "CreateVpc")
    return
}

func NewCreateVpcResponse() (response *CreateVpcResponse) {
    response = &CreateVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateVpc)用于创建私有网络(VPC)。
// * 用户可以创建的最小网段子网掩码为28（有16个IP地址），最大网段子网掩码为16（65,536个IP地址）,如果规划VPC网段请参见VPC网段规划说明。
// * 创建VPC时可同时把子网创建好，创建子网也请规划好子网网段及子网所在可用区，同一个VPC内子网网段不能重叠，不同可用区可以做跨可用区容灾，详见VPC可用区说明。
// * 如果您同时创建了子网，系统会创建一个默认路由表，系统会把子网关联到这个默认路由表。
// * 同一个地域能创建的VPC资源个数也是有限制的，详见 <a href="https://cloud.tencent.com/doc/product/215/537" title="VPC使用限制">VPC使用限制</a>,如果需要扩充请联系在线客服。
func (c *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
    if request == nil {
        request = NewCreateVpcRequest()
    }
    response = NewCreateVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressTemplateRequest() (request *DeleteAddressTemplateRequest) {
    request = &DeleteAddressTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplate")
    return
}

func NewDeleteAddressTemplateResponse() (response *DeleteAddressTemplateResponse) {
    response = &DeleteAddressTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除IP地址模板
func (c *Client) DeleteAddressTemplate(request *DeleteAddressTemplateRequest) (response *DeleteAddressTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateRequest()
    }
    response = NewDeleteAddressTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAddressTemplateGroupRequest() (request *DeleteAddressTemplateGroupRequest) {
    request = &DeleteAddressTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteAddressTemplateGroup")
    return
}

func NewDeleteAddressTemplateGroupResponse() (response *DeleteAddressTemplateGroupResponse) {
    response = &DeleteAddressTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除IP地址模板集合
func (c *Client) DeleteAddressTemplateGroup(request *DeleteAddressTemplateGroupRequest) (response *DeleteAddressTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAddressTemplateGroupRequest()
    }
    response = NewDeleteAddressTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNetworkInterfaceRequest() (request *DeleteNetworkInterfaceRequest) {
    request = &DeleteNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteNetworkInterface")
    return
}

func NewDeleteNetworkInterfaceResponse() (response *DeleteNetworkInterfaceResponse) {
    response = &DeleteNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteNetworkInterface）用于创建弹性网卡。
// * 弹性网卡上绑定了云主机时，不能被删除。
// * 删除指定弹性网卡，弹性网卡必须先和子机解绑才能删除。删除之后弹性网卡上所有内网IP都将被退还。
func (c *Client) DeleteNetworkInterface(request *DeleteNetworkInterfaceRequest) (response *DeleteNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDeleteNetworkInterfaceRequest()
    }
    response = NewDeleteNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
    request = &DeleteRouteTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRouteTable")
    return
}

func NewDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
    response = &DeleteRouteTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除路由表
func (c *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTableRequest()
    }
    response = NewDeleteRouteTableResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoutesRequest() (request *DeleteRoutesRequest) {
    request = &DeleteRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteRoutes")
    return
}

func NewDeleteRoutesResponse() (response *DeleteRoutesResponse) {
    response = &DeleteRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DeleteRoutes)用于对某个路由表批量删除路由策略（Route）。
func (c *Client) DeleteRoutes(request *DeleteRoutesRequest) (response *DeleteRoutesResponse, err error) {
    if request == nil {
        request = NewDeleteRoutesRequest()
    }
    response = NewDeleteRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupRequest() (request *DeleteSecurityGroupRequest) {
    request = &DeleteSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroup")
    return
}

func NewDeleteSecurityGroupResponse() (response *DeleteSecurityGroupResponse) {
    response = &DeleteSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteSecurityGroup）用于删除安全组（SecurityGroup）。
// * 只有当前账号下的安全组允许被删除。
// * 安全组实例ID如果在其他安全组的规则中被引用，则无法直接删除。这种情况下，需要先进行规则修改，再删除安全组。
// * 删除的安全组无法再找回，请谨慎调用。
func (c *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (response *DeleteSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupRequest()
    }
    response = NewDeleteSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityGroupPoliciesRequest() (request *DeleteSecurityGroupPoliciesRequest) {
    request = &DeleteSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSecurityGroupPolicies")
    return
}

func NewDeleteSecurityGroupPoliciesResponse() (response *DeleteSecurityGroupPoliciesResponse) {
    response = &DeleteSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteSecurityGroupPolicies）用于用于删除安全组规则（SecurityGroupPolicy）。
// * SecurityGroupPolicySet.Version 用于指定要操作的安全组的版本。传入 Version 版本号若不等于当前安全组的最新版本，将返回失败；若不传 Version 则直接删除指定PolicyIndex的规则。
func (c *Client) DeleteSecurityGroupPolicies(request *DeleteSecurityGroupPoliciesRequest) (response *DeleteSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityGroupPoliciesRequest()
    }
    response = NewDeleteSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceTemplateRequest() (request *DeleteServiceTemplateRequest) {
    request = &DeleteServiceTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplate")
    return
}

func NewDeleteServiceTemplateResponse() (response *DeleteServiceTemplateResponse) {
    response = &DeleteServiceTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除协议端口模板
func (c *Client) DeleteServiceTemplate(request *DeleteServiceTemplateRequest) (response *DeleteServiceTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateRequest()
    }
    response = NewDeleteServiceTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceTemplateGroupRequest() (request *DeleteServiceTemplateGroupRequest) {
    request = &DeleteServiceTemplateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteServiceTemplateGroup")
    return
}

func NewDeleteServiceTemplateGroupResponse() (response *DeleteServiceTemplateGroupResponse) {
    response = &DeleteServiceTemplateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除协议端口模板集合
func (c *Client) DeleteServiceTemplateGroup(request *DeleteServiceTemplateGroupRequest) (response *DeleteServiceTemplateGroupResponse, err error) {
    if request == nil {
        request = NewDeleteServiceTemplateGroupRequest()
    }
    response = NewDeleteServiceTemplateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSubnetRequest() (request *DeleteSubnetRequest) {
    request = &DeleteSubnetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteSubnet")
    return
}

func NewDeleteSubnetResponse() (response *DeleteSubnetResponse) {
    response = &DeleteSubnetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteSubnet）用于用于删除子网(Subnet)。
// * 删除子网前，请清理该子网下所有资源，包括云主机、负载均衡、云数据、noSql、弹性网卡等资源。
func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (response *DeleteSubnetResponse, err error) {
    if request == nil {
        request = NewDeleteSubnetRequest()
    }
    response = NewDeleteSubnetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVpcRequest() (request *DeleteVpcRequest) {
    request = &DeleteVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DeleteVpc")
    return
}

func NewDeleteVpcResponse() (response *DeleteVpcResponse) {
    response = &DeleteVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteVpc）用于删除私有网络。
// * 删除前请确保 VPC 内已经没有相关资源，例如云主机、云数据库、NoSQL、VPN网关、专线网关、负载均衡、对等连接、与之互通的基础网络设备等。
// * 删除私有网络是不可逆的操作，请谨慎处理。
func (c *Client) DeleteVpc(request *DeleteVpcRequest) (response *DeleteVpcResponse, err error) {
    if request == nil {
        request = NewDeleteVpcRequest()
    }
    response = NewDeleteVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressQuotaRequest() (request *DescribeAddressQuotaRequest) {
    request = &DescribeAddressQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressQuota")
    return
}

func NewDescribeAddressQuotaResponse() (response *DescribeAddressQuotaResponse) {
    response = &DescribeAddressQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAddressQuota) 用于查询您账户的[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）在当前地域的配额信息。配额详情可参见 [EIP 产品简介](https://cloud.tencent.com/document/product/213/5733)。
func (c *Client) DescribeAddressQuota(request *DescribeAddressQuotaRequest) (response *DescribeAddressQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeAddressQuotaRequest()
    }
    response = NewDescribeAddressQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressTemplateGroupsRequest() (request *DescribeAddressTemplateGroupsRequest) {
    request = &DescribeAddressTemplateGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplateGroups")
    return
}

func NewDescribeAddressTemplateGroupsResponse() (response *DescribeAddressTemplateGroupsResponse) {
    response = &DescribeAddressTemplateGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询IP地址模板集合
func (c *Client) DescribeAddressTemplateGroups(request *DescribeAddressTemplateGroupsRequest) (response *DescribeAddressTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplateGroupsRequest()
    }
    response = NewDescribeAddressTemplateGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressTemplatesRequest() (request *DescribeAddressTemplatesRequest) {
    request = &DescribeAddressTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddressTemplates")
    return
}

func NewDescribeAddressTemplatesResponse() (response *DescribeAddressTemplatesResponse) {
    response = &DescribeAddressTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询IP地址模板
func (c *Client) DescribeAddressTemplates(request *DescribeAddressTemplatesRequest) (response *DescribeAddressTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressTemplatesRequest()
    }
    response = NewDescribeAddressTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAddressesRequest() (request *DescribeAddressesRequest) {
    request = &DescribeAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeAddresses")
    return
}

func NewDescribeAddressesResponse() (response *DescribeAddressesResponse) {
    response = &DescribeAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeAddresses) 用于查询一个或多个[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）的详细信息。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的 EIP。
func (c *Client) DescribeAddresses(request *DescribeAddressesRequest) (response *DescribeAddressesResponse, err error) {
    if request == nil {
        request = NewDescribeAddressesRequest()
    }
    response = NewDescribeAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClassicLinkInstancesRequest() (request *DescribeClassicLinkInstancesRequest) {
    request = &DescribeClassicLinkInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeClassicLinkInstances")
    return
}

func NewDescribeClassicLinkInstancesResponse() (response *DescribeClassicLinkInstancesResponse) {
    response = &DescribeClassicLinkInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeClassicLinkInstances)用于私有网络和基础网络设备互通。
func (c *Client) DescribeClassicLinkInstances(request *DescribeClassicLinkInstancesRequest) (response *DescribeClassicLinkInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClassicLinkInstancesRequest()
    }
    response = NewDescribeClassicLinkInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
    request = &DescribeNetworkInterfacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeNetworkInterfaces")
    return
}

func NewDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
    response = &DescribeNetworkInterfacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeNetworkInterfaces）用于查询弹性网卡列表。
func (c *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
    if request == nil {
        request = NewDescribeNetworkInterfacesRequest()
    }
    response = NewDescribeNetworkInterfacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
    request = &DescribeRouteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeRouteTables")
    return
}

func NewDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
    response = &DescribeRouteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  本接口（DescribeRouteTables）用于查询路由表。
func (c *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
    if request == nil {
        request = NewDescribeRouteTablesRequest()
    }
    response = NewDescribeRouteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupPoliciesRequest() (request *DescribeSecurityGroupPoliciesRequest) {
    request = &DescribeSecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroupPolicies")
    return
}

func NewDescribeSecurityGroupPoliciesResponse() (response *DescribeSecurityGroupPoliciesResponse) {
    response = &DescribeSecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSecurityGroupPolicies）用于查询安全组规则。
func (c *Client) DescribeSecurityGroupPolicies(request *DescribeSecurityGroupPoliciesRequest) (response *DescribeSecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupPoliciesRequest()
    }
    response = NewDescribeSecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
    request = &DescribeSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSecurityGroups")
    return
}

func NewDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
    response = &DescribeSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSecurityGroups）用于查询安全组。
func (c *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupsRequest()
    }
    response = NewDescribeSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTemplateGroupsRequest() (request *DescribeServiceTemplateGroupsRequest) {
    request = &DescribeServiceTemplateGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplateGroups")
    return
}

func NewDescribeServiceTemplateGroupsResponse() (response *DescribeServiceTemplateGroupsResponse) {
    response = &DescribeServiceTemplateGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询协议端口模板集合
func (c *Client) DescribeServiceTemplateGroups(request *DescribeServiceTemplateGroupsRequest) (response *DescribeServiceTemplateGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplateGroupsRequest()
    }
    response = NewDescribeServiceTemplateGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceTemplatesRequest() (request *DescribeServiceTemplatesRequest) {
    request = &DescribeServiceTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeServiceTemplates")
    return
}

func NewDescribeServiceTemplatesResponse() (response *DescribeServiceTemplatesResponse) {
    response = &DescribeServiceTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询协议端口模板
func (c *Client) DescribeServiceTemplates(request *DescribeServiceTemplatesRequest) (response *DescribeServiceTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeServiceTemplatesRequest()
    }
    response = NewDescribeServiceTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubnetsRequest() (request *DescribeSubnetsRequest) {
    request = &DescribeSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeSubnets")
    return
}

func NewDescribeSubnetsResponse() (response *DescribeSubnetsResponse) {
    response = &DescribeSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSubnets）用于查询子网列表。
func (c *Client) DescribeSubnets(request *DescribeSubnetsRequest) (response *DescribeSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeSubnetsRequest()
    }
    response = NewDescribeSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVpcsRequest() (request *DescribeVpcsRequest) {
    request = &DescribeVpcsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DescribeVpcs")
    return
}

func NewDescribeVpcsResponse() (response *DescribeVpcsResponse) {
    response = &DescribeVpcsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeVpcs）用于查询私有网络列表。
func (c *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
    if request == nil {
        request = NewDescribeVpcsRequest()
    }
    response = NewDescribeVpcsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachClassicLinkVpcRequest() (request *DetachClassicLinkVpcRequest) {
    request = &DetachClassicLinkVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DetachClassicLinkVpc")
    return
}

func NewDetachClassicLinkVpcResponse() (response *DetachClassicLinkVpcResponse) {
    response = &DetachClassicLinkVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DetachClassicLinkVpc)用于删除私有网络和基础网络设备互通。
func (c *Client) DetachClassicLinkVpc(request *DetachClassicLinkVpcRequest) (response *DetachClassicLinkVpcResponse, err error) {
    if request == nil {
        request = NewDetachClassicLinkVpcRequest()
    }
    response = NewDetachClassicLinkVpcResponse()
    err = c.Send(request, response)
    return
}

func NewDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
    request = &DetachNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DetachNetworkInterface")
    return
}

func NewDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
    response = &DetachNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachNetworkInterface）用于弹性网卡解绑云主机。
func (c *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewDetachNetworkInterfaceRequest()
    }
    response = NewDetachNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateAddressRequest() (request *DisassociateAddressRequest) {
    request = &DisassociateAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "DisassociateAddress")
    return
}

func NewDisassociateAddressResponse() (response *DisassociateAddressResponse) {
    response = &DisassociateAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DisassociateAddress) 用于解绑[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）。
// * 只有状态为 BIND 和 BIND_ENI 的 EIP 才能进行解绑定操作。
// * EIP 如果被封堵，则不能进行解绑定操作。
func (c *Client) DisassociateAddress(request *DisassociateAddressRequest) (response *DisassociateAddressResponse, err error) {
    if request == nil {
        request = NewDisassociateAddressRequest()
    }
    response = NewDisassociateAddressResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateNetworkInterfaceRequest() (request *MigrateNetworkInterfaceRequest) {
    request = &MigrateNetworkInterfaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "MigrateNetworkInterface")
    return
}

func NewMigrateNetworkInterfaceResponse() (response *MigrateNetworkInterfaceResponse) {
    response = &MigrateNetworkInterfaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（MigrateNetworkInterface）用于弹性网卡迁移。
func (c *Client) MigrateNetworkInterface(request *MigrateNetworkInterfaceRequest) (response *MigrateNetworkInterfaceResponse, err error) {
    if request == nil {
        request = NewMigrateNetworkInterfaceRequest()
    }
    response = NewMigrateNetworkInterfaceResponse()
    err = c.Send(request, response)
    return
}

func NewMigratePrivateIpAddressRequest() (request *MigratePrivateIpAddressRequest) {
    request = &MigratePrivateIpAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "MigratePrivateIpAddress")
    return
}

func NewMigratePrivateIpAddressResponse() (response *MigratePrivateIpAddressResponse) {
    response = &MigratePrivateIpAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

//  本接口（MigratePrivateIpAddress）用于弹性网卡内网IP迁移。
// 
// * 该接口用于将一个内网IP从一个弹性网卡上迁移到另外一个弹性网卡，主IP地址不支持迁移。
// * 迁移前后的弹性网卡必须在同一个子网内。
func (c *Client) MigratePrivateIpAddress(request *MigratePrivateIpAddressRequest) (response *MigratePrivateIpAddressResponse, err error) {
    if request == nil {
        request = NewMigratePrivateIpAddressRequest()
    }
    response = NewMigratePrivateIpAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressAttributeRequest() (request *ModifyAddressAttributeRequest) {
    request = &ModifyAddressAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressAttribute")
    return
}

func NewModifyAddressAttributeResponse() (response *ModifyAddressAttributeResponse) {
    response = &ModifyAddressAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyAddressAttribute) 用于修改[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）的名称。
func (c *Client) ModifyAddressAttribute(request *ModifyAddressAttributeRequest) (response *ModifyAddressAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressAttributeRequest()
    }
    response = NewModifyAddressAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressTemplateAttributeRequest() (request *ModifyAddressTemplateAttributeRequest) {
    request = &ModifyAddressTemplateAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateAttribute")
    return
}

func NewModifyAddressTemplateAttributeResponse() (response *ModifyAddressTemplateAttributeResponse) {
    response = &ModifyAddressTemplateAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改IP地址模板
func (c *Client) ModifyAddressTemplateAttribute(request *ModifyAddressTemplateAttributeRequest) (response *ModifyAddressTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateAttributeRequest()
    }
    response = NewModifyAddressTemplateAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAddressTemplateGroupAttributeRequest() (request *ModifyAddressTemplateGroupAttributeRequest) {
    request = &ModifyAddressTemplateGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyAddressTemplateGroupAttribute")
    return
}

func NewModifyAddressTemplateGroupAttributeResponse() (response *ModifyAddressTemplateGroupAttributeResponse) {
    response = &ModifyAddressTemplateGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改IP地址模板集合
func (c *Client) ModifyAddressTemplateGroupAttribute(request *ModifyAddressTemplateGroupAttributeRequest) (response *ModifyAddressTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAddressTemplateGroupAttributeRequest()
    }
    response = NewModifyAddressTemplateGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkInterfaceAttributeRequest() (request *ModifyNetworkInterfaceAttributeRequest) {
    request = &ModifyNetworkInterfaceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyNetworkInterfaceAttribute")
    return
}

func NewModifyNetworkInterfaceAttributeResponse() (response *ModifyNetworkInterfaceAttributeResponse) {
    response = &ModifyNetworkInterfaceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyNetworkInterfaceAttribute）用于修改弹性网卡属性。
func (c *Client) ModifyNetworkInterfaceAttribute(request *ModifyNetworkInterfaceAttributeRequest) (response *ModifyNetworkInterfaceAttributeResponse, err error) {
    if request == nil {
        request = NewModifyNetworkInterfaceAttributeRequest()
    }
    response = NewModifyNetworkInterfaceAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateIpAddressesAttributeRequest() (request *ModifyPrivateIpAddressesAttributeRequest) {
    request = &ModifyPrivateIpAddressesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyPrivateIpAddressesAttribute")
    return
}

func NewModifyPrivateIpAddressesAttributeResponse() (response *ModifyPrivateIpAddressesAttributeResponse) {
    response = &ModifyPrivateIpAddressesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyPrivateIpAddressesAttribute）用于修改弹性网卡内网IP属性。
func (c *Client) ModifyPrivateIpAddressesAttribute(request *ModifyPrivateIpAddressesAttributeRequest) (response *ModifyPrivateIpAddressesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyPrivateIpAddressesAttributeRequest()
    }
    response = NewModifyPrivateIpAddressesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRouteTableAttributeRequest() (request *ModifyRouteTableAttributeRequest) {
    request = &ModifyRouteTableAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyRouteTableAttribute")
    return
}

func NewModifyRouteTableAttributeResponse() (response *ModifyRouteTableAttributeResponse) {
    response = &ModifyRouteTableAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyRouteTableAttribute）用于修改路由表（RouteTable）属性。
func (c *Client) ModifyRouteTableAttribute(request *ModifyRouteTableAttributeRequest) (response *ModifyRouteTableAttributeResponse, err error) {
    if request == nil {
        request = NewModifyRouteTableAttributeRequest()
    }
    response = NewModifyRouteTableAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupAttributeRequest() (request *ModifySecurityGroupAttributeRequest) {
    request = &ModifySecurityGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupAttribute")
    return
}

func NewModifySecurityGroupAttributeResponse() (response *ModifySecurityGroupAttributeResponse) {
    response = &ModifySecurityGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifySecurityGroupAttribute）用于修改安全组（SecurityGroupPolicy）属性。
func (c *Client) ModifySecurityGroupAttribute(request *ModifySecurityGroupAttributeRequest) (response *ModifySecurityGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupAttributeRequest()
    }
    response = NewModifySecurityGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupPoliciesRequest() (request *ModifySecurityGroupPoliciesRequest) {
    request = &ModifySecurityGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySecurityGroupPolicies")
    return
}

func NewModifySecurityGroupPoliciesResponse() (response *ModifySecurityGroupPoliciesResponse) {
    response = &ModifySecurityGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifySecurityGroupPolicies）用于重置安全组出站和入站规则（SecurityGroupPolicy）。
// 
// * 接口是先删除当前所有的出入站规则，然后再添加 Egress 和 Ingress 规则，不支持自定义索引 PolicyIndex 。
// * 如果指定 SecurityGroupPolicySet.Version 为0, 表示清空所有规则，并忽略Egress和Ingress。
// * Protocol字段支持输入TCP, UDP, ICMP, GRE, ALL。
// * CidrBlock字段允许输入符合cidr格式标准的任意字符串。(展开)在基础网络中，如果CidrBlock包含您的账户内的云服务器之外的设备在腾讯云的内网IP，并不代表此规则允许您访问这些设备，租户之间网络隔离规则优先于安全组中的内网规则。
// * SecurityGroupId字段允许输入与待修改的安全组位于相同项目中的安全组ID，包括这个安全组ID本身，代表安全组下所有云服务器的内网IP。使用这个字段时，这条规则用来匹配网络报文的过程中会随着被使用的这个ID所关联的云服务器变化而变化，不需要重新修改。
// * Port字段允许输入一个单独端口号，或者用减号分隔的两个端口号代表端口范围，例如80或8000-8010。只有当Protocol字段是TCP或UDP时，Port字段才被接受。
// * Action字段只允许输入ACCEPT或DROP。
// * CidrBlock, SecurityGroupId, AddressTemplate三者是排他关系，不允许同时输入，Protocol + Port和ServiceTemplate二者是排他关系，不允许同时输入。
func (c *Client) ModifySecurityGroupPolicies(request *ModifySecurityGroupPoliciesRequest) (response *ModifySecurityGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupPoliciesRequest()
    }
    response = NewModifySecurityGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceTemplateAttributeRequest() (request *ModifyServiceTemplateAttributeRequest) {
    request = &ModifyServiceTemplateAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateAttribute")
    return
}

func NewModifyServiceTemplateAttributeResponse() (response *ModifyServiceTemplateAttributeResponse) {
    response = &ModifyServiceTemplateAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改协议端口模板
func (c *Client) ModifyServiceTemplateAttribute(request *ModifyServiceTemplateAttributeRequest) (response *ModifyServiceTemplateAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateAttributeRequest()
    }
    response = NewModifyServiceTemplateAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceTemplateGroupAttributeRequest() (request *ModifyServiceTemplateGroupAttributeRequest) {
    request = &ModifyServiceTemplateGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyServiceTemplateGroupAttribute")
    return
}

func NewModifyServiceTemplateGroupAttributeResponse() (response *ModifyServiceTemplateGroupAttributeResponse) {
    response = &ModifyServiceTemplateGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyServiceTemplateGroupAttribute）用于修改协议端口模板集合。
func (c *Client) ModifyServiceTemplateGroupAttribute(request *ModifyServiceTemplateGroupAttributeRequest) (response *ModifyServiceTemplateGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyServiceTemplateGroupAttributeRequest()
    }
    response = NewModifyServiceTemplateGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubnetAttributeRequest() (request *ModifySubnetAttributeRequest) {
    request = &ModifySubnetAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifySubnetAttribute")
    return
}

func NewModifySubnetAttributeResponse() (response *ModifySubnetAttributeResponse) {
    response = &ModifySubnetAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifySubnetAttribute）用于修改子网属性。
func (c *Client) ModifySubnetAttribute(request *ModifySubnetAttributeRequest) (response *ModifySubnetAttributeResponse, err error) {
    if request == nil {
        request = NewModifySubnetAttributeRequest()
    }
    response = NewModifySubnetAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVpcAttributeRequest() (request *ModifyVpcAttributeRequest) {
    request = &ModifyVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ModifyVpcAttribute")
    return
}

func NewModifyVpcAttributeResponse() (response *ModifyVpcAttributeResponse) {
    response = &ModifyVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyVpcAttribute）用于修改私有网络（VPC）的相关属性。
func (c *Client) ModifyVpcAttribute(request *ModifyVpcAttributeRequest) (response *ModifyVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyVpcAttributeRequest()
    }
    response = NewModifyVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseAddressesRequest() (request *ReleaseAddressesRequest) {
    request = &ReleaseAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReleaseAddresses")
    return
}

func NewReleaseAddressesResponse() (response *ReleaseAddressesResponse) {
    response = &ReleaseAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ReleaseAddresses) 用于释放一个或多个[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）。
// * 该操作不可逆，释放后 EIP 关联的 IP 地址将不再属于您的名下。
// * 只有状态为 UNBIND 的 EIP 才能进行释放操作。
func (c *Client) ReleaseAddresses(request *ReleaseAddressesRequest) (response *ReleaseAddressesResponse, err error) {
    if request == nil {
        request = NewReleaseAddressesRequest()
    }
    response = NewReleaseAddressesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRouteTableAssociationRequest() (request *ReplaceRouteTableAssociationRequest) {
    request = &ReplaceRouteTableAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRouteTableAssociation")
    return
}

func NewReplaceRouteTableAssociationResponse() (response *ReplaceRouteTableAssociationResponse) {
    response = &ReplaceRouteTableAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ReplaceRouteTableAssociation)用于修改子网（Subnet）关联的路由表（RouteTable）。
// * 一个子网只能关联一个路由表。
func (c *Client) ReplaceRouteTableAssociation(request *ReplaceRouteTableAssociationRequest) (response *ReplaceRouteTableAssociationResponse, err error) {
    if request == nil {
        request = NewReplaceRouteTableAssociationRequest()
    }
    response = NewReplaceRouteTableAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceRoutesRequest() (request *ReplaceRoutesRequest) {
    request = &ReplaceRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceRoutes")
    return
}

func NewReplaceRoutesResponse() (response *ReplaceRoutesResponse) {
    response = &ReplaceRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ReplaceRoutes）根据路由策略ID（RouteId）修改指定的路由策略（Route），支持批量修改。
func (c *Client) ReplaceRoutes(request *ReplaceRoutesRequest) (response *ReplaceRoutesResponse, err error) {
    if request == nil {
        request = NewReplaceRoutesRequest()
    }
    response = NewReplaceRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceSecurityGroupPolicyRequest() (request *ReplaceSecurityGroupPolicyRequest) {
    request = &ReplaceSecurityGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ReplaceSecurityGroupPolicy")
    return
}

func NewReplaceSecurityGroupPolicyResponse() (response *ReplaceSecurityGroupPolicyResponse) {
    response = &ReplaceSecurityGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ReplaceSecurityGroupPolicy）用于替换单条安全组规则（SecurityGroupPolicy）。
// 单个请求中只能替换单个方向的一条规则, 必须要指定索引（PolicyIndex）。
func (c *Client) ReplaceSecurityGroupPolicy(request *ReplaceSecurityGroupPolicyRequest) (response *ReplaceSecurityGroupPolicyResponse, err error) {
    if request == nil {
        request = NewReplaceSecurityGroupPolicyRequest()
    }
    response = NewReplaceSecurityGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResetRoutesRequest() (request *ResetRoutesRequest) {
    request = &ResetRoutesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "ResetRoutes")
    return
}

func NewResetRoutesResponse() (response *ResetRoutesResponse) {
    response = &ResetRoutesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResetRoutes）用于对某个路由表名称和所有路由策略（Route）进行重新设置。
// 注意: 调用本接口是先删除当前路由表中所有路由策略, 再保存新提交的路由策略内容, 会引起网络中断。
func (c *Client) ResetRoutes(request *ResetRoutesRequest) (response *ResetRoutesResponse, err error) {
    if request == nil {
        request = NewResetRoutesRequest()
    }
    response = NewResetRoutesResponse()
    err = c.Send(request, response)
    return
}

func NewTransformAddressRequest() (request *TransformAddressRequest) {
    request = &TransformAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "TransformAddress")
    return
}

func NewTransformAddressResponse() (response *TransformAddressResponse) {
    response = &TransformAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (TransformAddress) 用于将实例的普通公网 IP 转换为[弹性公网IP](https://cloud.tencent.com/document/product/213/1941)（简称 EIP）。
// * 平台对用户每地域每日解绑 EIP 重新分配普通公网 IP 次数有所限制（可参见 [EIP 产品简介](/document/product/213/1941)）。上述配额可通过 [DescribeAddressQuota](https://cloud.tencent.com/document/api/213/1378) 接口获取。
func (c *Client) TransformAddress(request *TransformAddressRequest) (response *TransformAddressResponse, err error) {
    if request == nil {
        request = NewTransformAddressRequest()
    }
    response = NewTransformAddressResponse()
    err = c.Send(request, response)
    return
}

func NewUnassignPrivateIpAddressesRequest() (request *UnassignPrivateIpAddressesRequest) {
    request = &UnassignPrivateIpAddressesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("vpc", APIVersion, "UnassignPrivateIpAddresses")
    return
}

func NewUnassignPrivateIpAddressesResponse() (response *UnassignPrivateIpAddressesResponse) {
    response = &UnassignPrivateIpAddressesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UnassignPrivateIpAddresses）用于弹性网卡退还内网 IP。
// * 退还弹性网卡上的辅助内网IP，接口自动解关联弹性公网 IP。不能退还弹性网卡的主内网IP。
func (c *Client) UnassignPrivateIpAddresses(request *UnassignPrivateIpAddressesRequest) (response *UnassignPrivateIpAddressesResponse, err error) {
    if request == nil {
        request = NewUnassignPrivateIpAddressesRequest()
    }
    response = NewUnassignPrivateIpAddressesResponse()
    err = c.Send(request, response)
    return
}
