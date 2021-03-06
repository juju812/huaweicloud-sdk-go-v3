package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/live/v2/model"
)

type LiveAPIClient struct {
	hcClient *http_client.HcHttpClient
}

func NewLiveAPIClient(hcClient *http_client.HcHttpClient) *LiveAPIClient {
	return &LiveAPIClient{hcClient: hcClient}
}

func LiveAPIClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//查询播放域名带宽数据。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListBandwidthDetailV2(request *model.ListBandwidthDetailV2Request) (*model.ListBandwidthDetailV2Response, error) {
	requestDef := GenReqDefForListBandwidthDetailV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthDetailV2Response), nil
	}
}

//查询指定时间范围内播放带宽峰值。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListDomainBandwidthSummaryV2(request *model.ListDomainBandwidthSummaryV2Request) (*model.ListDomainBandwidthSummaryV2Response, error) {
	requestDef := GenReqDefForListDomainBandwidthSummaryV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainBandwidthSummaryV2Response), nil
	}
}

//查询播放域名流量数据。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListDomainTrafficDetailV2(request *model.ListDomainTrafficDetailV2Request) (*model.ListDomainTrafficDetailV2Response, error) {
	requestDef := GenReqDefForListDomainTrafficDetailV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficDetailV2Response), nil
	}
}

//查询指定时间范围内流量汇总量。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListDomainTrafficSummaryV2(request *model.ListDomainTrafficSummaryV2Request) (*model.ListDomainTrafficSummaryV2Response, error) {
	requestDef := GenReqDefForListDomainTrafficSummaryV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTrafficSummaryV2Response), nil
	}
}

//查询历史推流列表。  最大查询跨度1天，最大查询周期7天。
func (c *LiveAPIClient) ListHistoryStreamsV2(request *model.ListHistoryStreamsV2Request) (*model.ListHistoryStreamsV2Response, error) {
	requestDef := GenReqDefForListHistoryStreamsV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryStreamsV2Response), nil
	}
}

//查询直播拉流HTTP状态码接口。  获取加速域名1分钟粒度的HTTP返回码  最大查询跨度不能超过24小时，最大查询周期7天。
func (c *LiveAPIClient) ListQueryHttpCode(request *model.ListQueryHttpCodeRequest) (*model.ListQueryHttpCodeResponse, error) {
	requestDef := GenReqDefForListQueryHttpCode()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryHttpCodeResponse), nil
	}
}

//查询直播租户每小时录制的最大并发数，计算1小时内每分钟的并发总路数，取最大值做为统计值。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListRecordDataV2(request *model.ListRecordDataV2Request) (*model.ListRecordDataV2Response, error) {
	requestDef := GenReqDefForListRecordDataV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordDataV2Response), nil
	}
}

//查询推流监控码率数据接口。  最大查询跨度6小时，最大查询周期7天。
func (c *LiveAPIClient) ListSingleStreamBitrateV2(request *model.ListSingleStreamBitrateV2Request) (*model.ListSingleStreamBitrateV2Response, error) {
	requestDef := GenReqDefForListSingleStreamBitrateV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamBitrateV2Response), nil
	}
}

//查询推流帧率数据接口。  最大查询跨度6小时，最大查询周期7天。
func (c *LiveAPIClient) ListSingleStreamFramerateV2(request *model.ListSingleStreamFramerateV2Request) (*model.ListSingleStreamFramerateV2Response, error) {
	requestDef := GenReqDefForListSingleStreamFramerateV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSingleStreamFramerateV2Response), nil
	}
}

//查询直播域名每小时的截图数量。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListSnapshotDataV2(request *model.ListSnapshotDataV2Request) (*model.ListSnapshotDataV2Response, error) {
	requestDef := GenReqDefForListSnapshotDataV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotDataV2Response), nil
	}
}

//查询直播域名每小时的转码时长数据。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ListTranscodeDataV2(request *model.ListTranscodeDataV2Request) (*model.ListTranscodeDataV2Response, error) {
	requestDef := GenReqDefForListTranscodeDataV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTranscodeDataV2Response), nil
	}
}

//查询观众趋势。  最大查询跨度7天，最大查询周期90天。
func (c *LiveAPIClient) ListUsersOfStreamV2(request *model.ListUsersOfStreamV2Request) (*model.ListUsersOfStreamV2Response, error) {
	requestDef := GenReqDefForListUsersOfStreamV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersOfStreamV2Response), nil
	}
}

//查询域名维度推流路数接口。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ShowStreamCountV2(request *model.ShowStreamCountV2Request) (*model.ShowStreamCountV2Response, error) {
	requestDef := GenReqDefForShowStreamCountV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamCountV2Response), nil
	}
}

//查询播放画像信息。  最大查询跨度1天，最大查询周期31天。
func (c *LiveAPIClient) ShowStreamPortrait(request *model.ShowStreamPortraitRequest) (*model.ShowStreamPortraitResponse, error) {
	requestDef := GenReqDefForShowStreamPortrait()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStreamPortraitResponse), nil
	}
}

//查询上行带宽数据。  最大查询跨度31天，最大查询周期90天。
func (c *LiveAPIClient) ShowUpBandwidthV2(request *model.ShowUpBandwidthV2Request) (*model.ShowUpBandwidthV2Response, error) {
	requestDef := GenReqDefForShowUpBandwidthV2()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpBandwidthV2Response), nil
	}
}
