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

package v20170320

import (
    "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/anchnet/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-20"

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


func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "AssociateSecurityGroups")
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(AssociateSecurityGroups)用于安全组批量绑定实例。
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWanServiceRequest() (request *CloseWanServiceRequest) {
    request = &CloseWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CloseWanService")
    return
}

func NewCloseWanServiceResponse() (response *CloseWanServiceResponse) {
    response = &CloseWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CloseWanService)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问。
func (c *Client) CloseWanService(request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    if request == nil {
        request = NewCloseWanServiceRequest()
    }
    response = NewCloseWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateBackup)用于创建数据库备份。
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBImportJobRequest() (request *CreateDBImportJobRequest) {
    request = &CreateDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBImportJob")
    return
}

func NewCreateDBImportJobResponse() (response *CreateDBImportJobResponse) {
    response = &CreateDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateDBImportJob)用于创建云数据库数据导入任务。
func (c *Client) CreateDBImportJob(request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    if request == nil {
        request = NewCreateDBImportJobRequest()
    }
    response = NewCreateDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBInstance")
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateDBInstance)用于创建包年包月的云数据库实例（包括主实例、灾备实例和只读实例），可通过传入实例规格、MySQL 版本号、购买时长和数量等信息创建云数据库实例。
// 
// 您还可以使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872)接口查询该实例的详细信息。
// 
// 1. 首先请使用[获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229)接口查询可创建的实例规格信息，然后请使用[查询价格（包年包月）](https://cloud.tencent.com/document/api/236/1332)接口查询可创建实例的售卖价格；
// 
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
// 
// 3. 支持创建 MySQL5.5 、 MySQL5.6 、 MySQL5.7 版本；
// 
// 4. 支持创建主实例、只读实例、灾备实例；
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceHourRequest() (request *CreateDBInstanceHourRequest) {
    request = &CreateDBInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBInstanceHour")
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateDBInstanceHour)用于创建按量计费的实例，可通过传入实例规格、MySQL 版本号和数量等信息创建云数据库实例，支持主实例、灾备实例和只读实例的创建。
// 
// 您还可以使用[查询实例列表](https://cloud.tencent.com/document/api/236/15872)接口查询该实例的详细信息。
// 
// 1. 首先请使用[获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229)接口查询可创建的实例规格信息，然后请使用[查询价格（按量计费）](https://cloud.tencent.com/document/api/253/5176)接口查询可创建实例的售卖价格；
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
// 3. 支持创建 MySQL5.5、MySQL5.6和MySQL5.7 版本；
// 4. 支持创建主实例、灾备实例和只读实例；
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
    request = &DeleteBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteBackup")
    return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
    response = &DeleteBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DeleteBackup)用于删除数据库备份。
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRequest()
    }
    response = NewDeleteBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupConfig")
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackupConfig)用于查询数据库备份配置信息。
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDatabasesRequest() (request *DescribeBackupDatabasesRequest) {
    request = &DescribeBackupDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDatabases")
    return
}

func NewDescribeBackupDatabasesResponse() (response *DescribeBackupDatabasesResponse) {
    response = &DescribeBackupDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackupDatabases)用于查询备份数据库列表。
func (c *Client) DescribeBackupDatabases(request *DescribeBackupDatabasesRequest) (response *DescribeBackupDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDatabasesRequest()
    }
    response = NewDescribeBackupDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadDbTableCodeRequest() (request *DescribeBackupDownloadDbTableCodeRequest) {
    request = &DescribeBackupDownloadDbTableCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDownloadDbTableCode")
    return
}

func NewDescribeBackupDownloadDbTableCodeResponse() (response *DescribeBackupDownloadDbTableCodeResponse) {
    response = &DescribeBackupDownloadDbTableCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackupDownloadDbTableCode)用于查询备份数据分库分表下载位点。
func (c *Client) DescribeBackupDownloadDbTableCode(request *DescribeBackupDownloadDbTableCodeRequest) (response *DescribeBackupDownloadDbTableCodeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadDbTableCodeRequest()
    }
    response = NewDescribeBackupDownloadDbTableCodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupTablesRequest() (request *DescribeBackupTablesRequest) {
    request = &DescribeBackupTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupTables")
    return
}

func NewDescribeBackupTablesResponse() (response *DescribeBackupTablesResponse) {
    response = &DescribeBackupTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackupTables)用于查询指定的数据库的备份数据表名。
func (c *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (response *DescribeBackupTablesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupTablesRequest()
    }
    response = NewDescribeBackupTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackups")
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackups)用于查询云数据库实例的备份数据。
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogsRequest() (request *DescribeBinlogsRequest) {
    request = &DescribeBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBinlogs")
    return
}

func NewDescribeBinlogsResponse() (response *DescribeBinlogsResponse) {
    response = &DescribeBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBinlogs)用于查询云数据库实例的二进制数据。
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBImportRecordsRequest() (request *DescribeDBImportRecordsRequest) {
    request = &DescribeDBImportRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBImportRecords")
    return
}

func NewDescribeDBImportRecordsResponse() (response *DescribeDBImportRecordsResponse) {
    response = &DescribeDBImportRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBImportRecords)用于查询云数据库导入任务操作日志。
func (c *Client) DescribeDBImportRecords(request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBImportRecordsRequest()
    }
    response = NewDescribeDBImportRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceCharsetRequest() (request *DescribeDBInstanceCharsetRequest) {
    request = &DescribeDBInstanceCharsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceCharset")
    return
}

func NewDescribeDBInstanceCharsetResponse() (response *DescribeDBInstanceCharsetResponse) {
    response = &DescribeDBInstanceCharsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstanceCharset)用于查询云数据库实例的字符集，获取字符集的名称。
func (c *Client) DescribeDBInstanceCharset(request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceCharsetRequest()
    }
    response = NewDescribeDBInstanceCharsetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceGTIDRequest() (request *DescribeDBInstanceGTIDRequest) {
    request = &DescribeDBInstanceGTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceGTID")
    return
}

func NewDescribeDBInstanceGTIDResponse() (response *DescribeDBInstanceGTIDResponse) {
    response = &DescribeDBInstanceGTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstanceGTID)用于查询云数据库实例是否开通了GTID，不支持版本为5.5以及以下的实例。
func (c *Client) DescribeDBInstanceGTID(request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceGTIDRequest()
    }
    response = NewDescribeDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceRebootTimeRequest() (request *DescribeDBInstanceRebootTimeRequest) {
    request = &DescribeDBInstanceRebootTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceRebootTime")
    return
}

func NewDescribeDBInstanceRebootTimeResponse() (response *DescribeDBInstanceRebootTimeResponse) {
    response = &DescribeDBInstanceRebootTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstanceRebootTime)用于查询云数据库实例重启预计所需的时间。
func (c *Client) DescribeDBInstanceRebootTime(request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRebootTimeRequest()
    }
    response = NewDescribeDBInstanceRebootTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstances)用于查询云数据库实例列表，支持通过项目ID、实例ID、访问地址、实例状态等来筛选实例。
// 
// 1. 不指定任何过滤条件, 则默认返回20条实例记录，单次请求最多支持返回100条实例记录；
// 2. 支持查询主实例、灾备实例和只读实例信息列表。
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBZoneConfigRequest() (request *DescribeDBZoneConfigRequest) {
    request = &DescribeDBZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBZoneConfig")
    return
}

func NewDescribeDBZoneConfigResponse() (response *DescribeDBZoneConfigResponse) {
    response = &DescribeDBZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBZoneConfig)用于查询可创建的云数据库各地域可售卖的规格配置。
func (c *Client) DescribeDBZoneConfig(request *DescribeDBZoneConfigRequest) (response *DescribeDBZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBZoneConfigRequest()
    }
    response = NewDescribeDBZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSlowLogs")
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeSlowLogs)用于获取云数据库实例的慢查询日志。
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeTasks)用于查询云数据库实例任务列表。
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "DisassociateSecurityGroups")
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "InitDBInstances")
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(InitDBInstances)用于初始化云数据库实例，包括初始化密码、默认字符集、实例端口号等
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "IsolateDBInstance")
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(IsolateDBInstance)用于销毁云数据库实例，销毁之后不能通过IP和端口访问数据库，按量计费实例销毁后直接下线。
// 
// 本接口不支持包年包月实例；
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupConfig")
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyBackupConfig)用于修改数据库备份配置信息。
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceName)用于修改云数据库实例的名称。
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceProject")
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceProject)用于修改云数据库实例的所属项目。
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceVipVportRequest() (request *ModifyDBInstanceVipVportRequest) {
    request = &ModifyDBInstanceVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceVipVport")
    return
}

func NewModifyDBInstanceVipVportResponse() (response *ModifyDBInstanceVipVportResponse) {
    response = &ModifyDBInstanceVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceVipVport)用于修改云数据库实例的IP和端口号，也可进行基础网络转VPC网络和VPC网络下的子网变更。
func (c *Client) ModifyDBInstanceVipVport(request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVipVportRequest()
    }
    response = NewModifyDBInstanceVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstanceParam")
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyInstanceParam)用于修改云数据库实例的参数。
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWanServiceRequest() (request *OpenWanServiceRequest) {
    request = &OpenWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "OpenWanService")
    return
}

func NewOpenWanServiceResponse() (response *OpenWanServiceResponse) {
    response = &OpenWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(OpenWanService)用于开通实例外网访问
func (c *Client) OpenWanService(request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    if request == nil {
        request = NewOpenWanServiceRequest()
    }
    response = NewOpenWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewStopDBImportJobRequest() (request *StopDBImportJobRequest) {
    request = &StopDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "StopDBImportJob")
    return
}

func NewStopDBImportJobResponse() (response *StopDBImportJobResponse) {
    response = &StopDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(StopDBImportJob)用于终止数据导入任务。
func (c *Client) StopDBImportJob(request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    if request == nil {
        request = NewStopDBImportJobRequest()
    }
    response = NewStopDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchForUpgradeRequest() (request *SwitchForUpgradeRequest) {
    request = &SwitchForUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchForUpgrade")
    return
}

func NewSwitchForUpgradeResponse() (response *SwitchForUpgradeResponse) {
    response = &SwitchForUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(SwitchForUpgrade)用于切换访问新实例，针对主升级中的实例处于待切换状态时，用户可主动发起该流程
func (c *Client) SwitchForUpgrade(request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    if request == nil {
        request = NewSwitchForUpgradeRequest()
    }
    response = NewSwitchForUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstance")
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(UpgradeDBInstance)用于升级云数据库实例，实例类型支持主实例、灾备实例和只读实例
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
    request = &UpgradeDBInstanceEngineVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstanceEngineVersion")
    return
}

func NewUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
    response = &UpgradeDBInstanceEngineVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(UpgradeDBInstanceEngineVersion)用于升级云数据库实例版本，实例类型支持主实例、灾备实例和只读实例。
func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    response = NewUpgradeDBInstanceEngineVersionResponse()
    err = c.Send(request, response)
    return
}
