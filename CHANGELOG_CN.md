## 0.0.20-beta 2020-11-02
## HuaweiCloud SDK CES
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 创建告警规则接口增加资源分组字段
    - 查询告警历史接口响应体metadata修改，只返回total字段

## HuaweiCloud SDK ELB
 - ### 新增特性
    - 无
 - ### 解决问题
    - 创建转发规则接口参数名错误问题修复
 - ### 特性变更
    - 无

## 0.0.19-beta 2020-10-31
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - Fix: query参数中包含枚举变量时请求失败
    - [Issue 7](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/7) 解决json.Marshal变成{}对象的问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK CBR
 - ### 新增特性
    - 新增支持接口：TAG标签相关接口
 - ### 解决问题
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/17) 修复ListBackups接口响应体问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK CTS
 - ### 新增特性
    - 新增支持接口：ListQuotas（查询租户追踪器配额信息）
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK EPS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 接口名称调整，原有的`*EP`接口展开为`*EnterpriseProject`，涉及接口：
     1. ListEP → ListEnterpriseProject
     2. CreateEP → CreateEnterpriseProject
     3. ShowEP → ShowEnterpriseProject
     4. ModifyEP → ModifyEnterpriseProject
     5. EnableEP → EnableEnterpriseProject
     6. ShowEPQuota → ShowEnterpriseProjectQuota
     7. ShowResourceBindEP → ShowResourceBindEnterpriseProject
     8. DisableEP → DisableEnterpriseProject

## HuaweiCloud SDK Iam
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 接口名称调整，涉及接口：
     1. KeystoneCreateUserTokenByPasswordAndMFA → KeystoneCreateUserTokenByPasswordAndMfa
     2. CreateUnscopeTokenByIDPInitiated → CreateUnscopeTokenByIdpInitiated

## HuaweiCloud SDK ProjectMan
 - ### 新增特性
    - 新增支持接口：迭代信息、用户信息、项目成员、项目信息、项目指标、项目统计等相关方法的接口
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


## 0.0.18-beta 2020-10-20
## HuaweiCloud SDK ELB
 - ### 新增特性
    - 增加v2版本接口
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


## 0.0.17-beta 2020-10-14
## HuaweiCloud SDK BSS
 - ### 新增特性
    - 伙伴中心支持导出产品目录价
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK Live
 - ### 新增特性
    - 新增直播V2接口
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


## 0.0.16-beta 2020-10-12
## HuaweiCloud SDK CTS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 删除v1版本废弃接口

## HuaweiCloud SDK DevStar
 - ### 新增特性
    - 无
 - ### 解决问题
    - 服务客户端认证类型调整为全局认证，即GlobalCredentials
 - ### 特性变更
    - 无


## 0.0.15-beta 2020-09-30
## HuaweiCloud SDK AS
 - ### 新增特性
    - 支持弹性伸缩服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


## 0.0.14-beta 2020-09-24
## HuaweiCloud SDK BSS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复BssClient类无法加载的问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK EIP
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 接口ListPublicips调整，企业项目ID不支持多值查询


## 0.0.13-beta 2020-09-18
## HuaweiCloud SDK ECS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修正接口参数类型
 - ### 特性变更
    - 无

## HuaweiCloud SDK BSS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 运营能力开放接口更新

## HuaweiCloud SDK EIP
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复查询弹性公网IP不支持传入多个值的问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK ELB
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复部分接口参数错误的问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK IMS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 调整接口描述

## HuaweiCloud SDK Live
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 修改禁播参数描述


## 0.0.12.1-beta 2020-09-16
## HuaweiCloud SDK ECS
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复接口参数类型错误
 - ### 特性变更
    - 无

## HuaweiCloud SDK All
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 可选参数类型为list时，参数类型变更为指针，例如：[]string 将变更为 *[]string

## 0.0.12-beta 2020-09-11
## HuaweiCloud SDK KPS
 - ### 新增特性
    - 支持KPS服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK EVS
 - ### 新增特性
    - 支持EVS服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK CBR
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复返回值类型定义错误的问题
 - ### 特性变更
    - 无


# 0.0.11-beta 2020-09-09
## HuaweiCloud SDK CBR
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复资源相关接口类型错误的问题
 - ### 特性变更
    - 无

## HuaweiCloud SDK VPC
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复安全组相关类型错误的问题
 - ### 特性变更
    - 无


# 0.0.10-beta 2020-09-04
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - 修复yaml中没有定义format的整型枚举参数无法生成枚举代码的问题
 - ### 特性变更
    - 调整Http请求头的User-Agent信息


# 0.0.9-beta 2020-08-28
## HuaweiCloud SDK CloudPipeline
 - ### 新增特性
    - 支持流水线服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK EIP
 - ### 新增特性
    - 新增支持接口：弹性公网IP标签相关接口和共享带宽相关接口
 - ### 解决问题
    - 批量创建带宽接口，修改billing_info字段
 - ### 特性变更
    - 无

## HuaweiCloud SDK IAM
 - ### 新增特性
    - 新增支持接口：console一致性相关接口
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK ProjectMan
 - ### 新增特性
    - 支持项目管理服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK VPC
 - ### 新增特性
    - 新增支持接口：安全组、安全组规则和可用IP数
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


# 0.0.8-beta 2020-08-25
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - 无
 - ### 特性变更
    - 可选枚举变量类型变更为指针类型。

## HuaweiCloud SDK ELB
 - ### 新增特性
    - 支持弹性负载均衡服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


# 0.0.7-beta 2020-08-20
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - 解决部分枚举类型变量无法正常读取的问题。
 - ### 特性变更
    - 以_开头的枚举类型变量名称添加前缀 'E'。


# 0.0.6-beta 2020-08-20
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - 解决当结构体包含枚举类型变量场景下部分依赖丢失的问题。
 - ### 特性变更
    - 无


# 0.0.5-beta 2020-08-19
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - 解决枚举类型反序列化失败的问题。
    - 解决部分场景下云服务异常响应反序列化失败问题。
 - ### 特性变更
    - 无


# 0.0.4-beta 2020-08-18
## HuaweiCloud SDK Core
 - ### 新增特性
    - 无
 - ### 解决问题
    - Go 基础类型默认值序列化丢失问题修复
 - ### 特性变更
    - 无


# 0.0.3-beta 2020-08-14
## HuaweiCloud SDK APIG
 - ### 新增特性
    - 支持API网关
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK BSS
 - ### 新增特性
    - 支持能力开放服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


# 0.0.2-beta 2020-08-11
## HuaweiCloud SDK Core
 - ### 新增特性
    - 支持临时AK/SK认证模式
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK CBR
 - ### 新增特性
    - 支持云备份服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK CloudIDE
 - ### 新增特性
    - 支持云测服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK ECS
 - ### 新增特性
    - 支持弹性云服务器服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK EIP
 - ### 新增特性
    - 支持弹性公网IP服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK EVS
 - ### 新增特性
    - 支持云硬盘服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK IMS
 - ### 新增特性
    - 支持镜像服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK Live
 - ### 新增特性
    - 支持视频直播服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无

## HuaweiCloud SDK MPC
 - ### 新增特性
    - 支持媒体转码服务
 - ### 解决问题
    - 无
 - ### 特性变更
    - 无


# 3.0.1-beta 2020-07-30
## 首次发布
 - ### 已支持服务
    - Classroom
    - 云审计服务（CTS）
    - 模板引擎（DevStar）
    - 企业管理服务（EPS）
    - 统一身份认证服务（IAM）
    - 标签管理服务（TMS）
    - 虚拟私有云（VPC）
