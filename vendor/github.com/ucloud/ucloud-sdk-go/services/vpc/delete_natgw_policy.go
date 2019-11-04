// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteNATGWPolicyRequest is request schema for DeleteNATGWPolicy action
type DeleteNATGWPolicyRequest struct {
	request.CommonBase

	// [公共参数] 项目Id。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NAT网关Id
	NATGWId *string `required:"true"`

	// 端口转发规则Id
	PolicyId *string `required:"true"`
}

// DeleteNATGWPolicyResponse is response schema for DeleteNATGWPolicy action
type DeleteNATGWPolicyResponse struct {
	response.CommonBase
}

// NewDeleteNATGWPolicyRequest will create request of DeleteNATGWPolicy action.
func (c *VPCClient) NewDeleteNATGWPolicyRequest() *DeleteNATGWPolicyRequest {
	req := &DeleteNATGWPolicyRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteNATGWPolicy - 删除NAT网关端口转发规则
func (c *VPCClient) DeleteNATGWPolicy(req *DeleteNATGWPolicyRequest) (*DeleteNATGWPolicyResponse, error) {
	var err error
	var res DeleteNATGWPolicyResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteNATGWPolicy", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}