/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 创建loadbalancer的消息返回体
type LoadBalancer struct {
	// 功能描述：负载均衡器ID。
	Id string `json:"id"`
	// 功能描述：负载均衡器描述信息。
	Description string `json:"description"`
	// 功能描述：负载均衡器的配置状态。 取值范围：ACTIVE、PENDING_CREATE 或者ERROR。 约束：该字段为预留字段，暂未启用，默认为ACTIVE。
	ProvisioningStatus string `json:"provisioning_status"`
	// 功能描述：负载均衡器的管理状态。 约束：只支持设定为true。
	AdminStateUp bool `json:"admin_state_up"`
	// 功能描述：负载均衡器的生产者名称。 约束：只支持vlb。
	Provider string `json:"provider"`
	// 功能描述：负载均衡器关联的后端云服务器组ID的列表。
	Pools []PoolRef `json:"pools"`
	// 功能描述：负载均衡器关联的监听器ID的列表。
	Listeners []ListenerRef `json:"listeners"`
	// 功能描述：负载均衡器的操作状态。 取值范围：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。 约束：该字段为预留字段，暂未启用，默认为ONLINE。
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status"`
	// 功能描述：负载均衡器的虚拟IP。
	VipAddress string `json:"vip_address"`
	// 功能描述：负载均衡器所在的子网ID。 约束：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。
	VipSubnetCidrId string `json:"vip_subnet_cidr_id"`
	// 功能描述：负载均衡器的负载均衡器名称。
	Name string `json:"name"`
	// 负载均衡器所在的项目ID。
	ProjectId string `json:"project_id"`
	// 负载均衡器虚拟IP对应的端口ID。
	VipPortId string `json:"vip_port_id"`
	// 功能描述：负载均衡的标签列表。
	Tags []Tag `json:"tags"`
	// 功能描述：负载均衡器的创建时间。
	CreatedAt string `json:"created_at"`
	// 功能描述：负载均衡器的更新时间。
	UpdatedAt string `json:"updated_at"`
	// 功能描述：是否是性能保障性实例 取值范围：共享型：false；性能保障型：true
	Guaranteed bool `json:"guaranteed"`
	// 功能描述：实例对应的vpc属性。若无，则从vip_subnet_cidr_id获取。  约束：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。
	VpcId string `json:"vpc_id"`
	// 功能描述：公网ELB实例绑定EIP信息。
	Eips []EipInfo `json:"eips"`
	// 功能描述：双栈实例对应v6的ip地址。
	Ipv6VipAddress string `json:"ipv6_vip_address"`
	// 功能描述：双栈实例对应v6的网络id 。 约束： 1、默认为空，只有开启IPv6时才会传入。 2、vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。
	Ipv6VipVirsubnetId string `json:"ipv6_vip_virsubnet_id"`
	// 功能描述：IPv6的VIP端口id。
	Ipv6VipPortId string `json:"ipv6_vip_port_id"`
	// 功能描述：可用区列表。默认指定所有可利用的AZ。可调用nova接口（/v2/{project_id}/os-availability-zone）查询可用AZ
	AvailabilityZoneList []string `json:"availability_zone_list"`
	// 功能描述：企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 功能描述：预留资源账单信息，默认为空表示按需计费， 非空为包周期。 约束：admin权限才能更新此字段。
	BillingInfo string `json:"billing_info"`
	// 功能描述：四层Flavor。
	L4FlavorId string `json:"l4_flavor_id"`
	// 功能描述：预留L4 弹性flavor。
	L4ScaleFlavorId string `json:"l4_scale_flavor_id"`
	// 功能描述：七层Flavor。
	L7FlavorId string `json:"l7_flavor_id"`
	// 功能描述：预留弹性flavor。
	L7ScaleFlavorId string `json:"l7_scale_flavor_id"`
	// 功能描述：弹性公网EIP信息
	Publicips *[]PublicIpInfo `json:"publicips,omitempty"`
	// 功能描述：下联面子网ID  loadbalancer使用的下联面端口会动态的从这些网络中占用IP
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
	// 功能描述：下联面子网类型
	ElbVirsubnetType *LoadBalancerElbVirsubnetType `json:"elb_virsubnet_type,omitempty"`
	// 是否启用跨VPC后端转发
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 是否开启删除保护
	DeletionProtectionEnable *string `json:"deletion_protection_enable,omitempty"`
	// 负载均衡器的冻结场景。若负载均衡器有多个冻结场景，用逗号分隔 POLICE：公安冻结场景。 ILLEGAL：违规冻结场景。 VERIFY：客户未实名认证冻结场景。 PARTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。 ARREAR：欠费冻结场景。
	FrozenScene   string        `json:"frozen_scene"`
	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
}

func (o LoadBalancer) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"LoadBalancer", string(data)}, " ")
}

type LoadBalancerOperatingStatus struct {
	value string
}

type LoadBalancerOperatingStatusEnum struct {
	ONLINE     LoadBalancerOperatingStatus
	OFFLINE    LoadBalancerOperatingStatus
	DEGRADED   LoadBalancerOperatingStatus
	DISABLED   LoadBalancerOperatingStatus
	NO_MONITOR LoadBalancerOperatingStatus
}

func GetLoadBalancerOperatingStatusEnum() LoadBalancerOperatingStatusEnum {
	return LoadBalancerOperatingStatusEnum{
		ONLINE: LoadBalancerOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadBalancerOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadBalancerOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadBalancerOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadBalancerOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadBalancerOperatingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *LoadBalancerOperatingStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type LoadBalancerElbVirsubnetType struct {
	value string
}

type LoadBalancerElbVirsubnetTypeEnum struct {
	IPV4      LoadBalancerElbVirsubnetType
	DUALSTACK LoadBalancerElbVirsubnetType
}

func GetLoadBalancerElbVirsubnetTypeEnum() LoadBalancerElbVirsubnetTypeEnum {
	return LoadBalancerElbVirsubnetTypeEnum{
		IPV4: LoadBalancerElbVirsubnetType{
			value: "ipv4",
		},
		DUALSTACK: LoadBalancerElbVirsubnetType{
			value: "dualstack",
		},
	}
}

func (c LoadBalancerElbVirsubnetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *LoadBalancerElbVirsubnetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
