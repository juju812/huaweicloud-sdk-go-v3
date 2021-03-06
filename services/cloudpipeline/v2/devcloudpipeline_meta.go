package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpipeline/v2/model"
	"net/http"
)

func GenReqDefForBatchShowPipelinesStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/pipelines/status").
		WithResponse(new(model.BatchShowPipelinesStatusResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineIds").
		WithJsonTag("pipeline_ids").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreatePipelineByTemplate() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/templates/task").
		WithResponse(new(model.CreatePipelineByTemplateResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListTemplates() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/templates").
		WithResponse(new(model.ListTemplatesResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateType").
		WithJsonTag("template_type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IsBuildIn").
		WithJsonTag("is_build_in").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Sort").
		WithJsonTag("sort").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Asc").
		WithJsonTag("asc").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRegisterAgent() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/agentregister/v1/agent/register").
		WithResponse(new(model.RegisterAgentResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRemovePipeline() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/pipelines/{pipeline_id}").
		WithResponse(new(model.RemovePipelineResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineId").
		WithJsonTag("pipeline_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowAgentStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/agents/{agent_id}/status").
		WithResponse(new(model.ShowAgentStatusResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgentId").
		WithJsonTag("agent_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowInstanceStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/templates/{task_id}/status").
		WithResponse(new(model.ShowInstanceStatusResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TaskId").
		WithJsonTag("task_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPipleineStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/pipelines/{pipeline_id}/status").
		WithResponse(new(model.ShowPipleineStatusResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineId").
		WithJsonTag("pipeline_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("BuildId").
		WithJsonTag("build_id").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowTemplateDetail() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/templates/{template_id}").
		WithResponse(new(model.ShowTemplateDetailResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateId").
		WithJsonTag("template_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateType").
		WithJsonTag("template_type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Source").
		WithJsonTag("source").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForStartNewPipeline() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/pipelines/{pipeline_id}/start").
		WithResponse(new(model.StartNewPipelineResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineId").
		WithJsonTag("pipeline_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForStartPipeline() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/pipelines/start").
		WithResponse(new(model.StartPipelineResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineId").
		WithJsonTag("pipeline_id").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForStopPipeline() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/pipelines/stop").
		WithResponse(new(model.StopPipelineResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PipelineId").
		WithJsonTag("pipeline_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("BuildId").
		WithJsonTag("build_id").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XLanguage").
		WithJsonTag("X-Language").
		WithLocationType(def.Header))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}
