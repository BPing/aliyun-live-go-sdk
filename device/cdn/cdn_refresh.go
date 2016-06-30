package cdn

import (
	"github.com/BPing/aliyun-live-go-sdk/client"
	"fmt"
	"github.com/BPing/aliyun-live-go-sdk/util/global"
	"errors"
)

//  刷新预热接口
// -------------------------------------------------------------------------------

//  RefreshObjectCaches 刷新节点上的文件内容。 刷新指定URL内容至Cache节点，每次只能提交一个url。
//  @param  RefreshObjectType   FileRefreshType|DirectoryRefreshType
//  @link https://help.aliyun.com/document_detail/27200.html?spm=0.0.0.0.InnEqg
func (c *CDN)RefreshObjectCaches(objectPath string, objectType RefreshObjectType, resp interface{}) (err error) {
	if (objectPath == global.EmptyString) {
		return errors.New("objectPath should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = RefreshObjectCachesAction
	req.SetArgs("ObjectPath", objectPath)
	if (global.EmptyString != objectType) {
		req.SetArgs("ObjectType", string(objectType))
	}
	err = c.rpc.Query(req, resp)
	return
}

//  PushObjectCache 将源站的内容主动预热到L2 Cache节点上，用户首次访问可直接命中缓存，缓解源站压力。
//
//  @link https://help.aliyun.com/document_detail/27201.html?spm=0.0.0.0.PJp719
func (c *CDN)PushObjectCache(objectPath string, resp interface{}) (err error) {
	if (objectPath == global.EmptyString) {
		return errors.New("objectPath should not be empty")
	}
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = PushObjectCacheAction
	req.SetArgs("ObjectPath", objectPath)
	err = c.rpc.Query(req, resp)
	return
}

//  GetRefreshTasks 查询刷新、预热状态，是否在全网生效。
//
//  @link https://help.aliyun.com/document_detail/27202.html?spm=0.0.0.0.Otm79O
func (c *CDN)GetRefreshTasks(taskId, objectPath string, pageSize, pageNumber int64, resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeRefreshTasksAction
	if (global.EmptyString != objectPath) {
		req.SetArgs("ObjectPath", objectPath)
	}
	if (global.EmptyString != taskId) {
		req.SetArgs("TaskId", taskId)
	}
	if (pageSize >= 1) {
		req.SetArgs("PageSize", fmt.Sprintf("%d", pageSize))
	}

	if (pageNumber >= 1) {
		req.SetArgs("PageNumber", fmt.Sprintf("%d", pageNumber))
	}
	err = c.rpc.Query(req, resp)
	return
}

//  GetRefreshQuota 刷新（包含预热）URL及目录的最大限制数量
//                  当日剩余刷新（含预热）URL及目录的次数
//
//  @link https://help.aliyun.com/document_detail/27203.html?spm=0.0.0.0.qSO8re
func (c *CDN)GetRefreshQuota(resp interface{}) (err error) {
	req := c.cdnReq.Clone().(*client.CDNRequest)
	req.Action = DescribeRefreshQuotaAction
	err = c.rpc.Query(req, resp)
	return
}

