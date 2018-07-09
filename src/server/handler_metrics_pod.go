package server

import (
	"github.com/linkernetworks/vortex/src/entity"
	response "github.com/linkernetworks/vortex/src/net/http"
	"github.com/linkernetworks/vortex/src/net/http/query"
	pc "github.com/linkernetworks/vortex/src/prometheuscontroller"
	"github.com/linkernetworks/vortex/src/web"
)

func listPodMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	expression := pc.Expression{}
	expression.Metrics = []string{"kube_pod_info"}
	expression.QueryLabels = map[string]string{}

	if node, ok := query.Str("node"); ok {
		expression.QueryLabels["node"] = node
	} else {
		expression.QueryLabels["node"] = ".*"
	}

	if namespace, ok := query.Str("namespace"); ok {
		expression.QueryLabels["namespace"] = namespace
	} else {
		expression.QueryLabels["namespace"] = ".*"
	}

	if controller, ok := query.Str("controller"); ok {
		expression.QueryLabels["deployment"] = controller
	} else {
		expression.QueryLabels["deployment"] = ".*"
	}

	containerList, err := pc.ListResource(sp, "pod", expression)
	if err != nil {
		response.BadRequest(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(containerList)
}

func getPodMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	pod := entity.PodMetrics{}
	id := req.PathParameter("id")

	pod, err := pc.GetPod(sp, id)
	if err != nil {
		response.BadRequest(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(pod)
}
