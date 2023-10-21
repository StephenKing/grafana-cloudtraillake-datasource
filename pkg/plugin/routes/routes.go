package routes

import (
	"github.com/emnify/cloud-trail-lake/pkg/plugin"
	"github.com/grafana/grafana-aws-sdk/pkg/sql/routes"
	"net/http"
)

type CtlResourceHandler struct {
	routes.ResourceHandler
	plugin plugin.CtlDatasourceIface
}

func New(api plugin.CtlDatasourceIface) *CtlResourceHandler {
	return &CtlResourceHandler{routes.ResourceHandler{API: api}, api}
}

func (r *CtlResourceHandler) Routes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := r.DefaultRoutes()
	//routes["/catalogs"] = r.catalogs
	//routes["/workgroups"] = r.workgroups
	//routes["/workgroupEngineVersion"] = r.workgroupEngineVersion
	return routes
}
