package routes

import (
	"github.com/emnify/cloud-trail-lake/pkg/plugin"
	"github.com/grafana/grafana-aws-sdk/pkg/sql/routes"
	"github.com/grafana/sqlds/v2"
	"net/http"
)

type CtlResourceHandler struct {
	routes.ResourceHandler
	plugin plugin.CtlDatasourceIface
}

func New(api plugin.CtlDatasourceIface) *CtlResourceHandler {
	return &CtlResourceHandler{routes.ResourceHandler{API: api}, api}
}

func (r *CtlResourceHandler) eventDataStores(rw http.ResponseWriter, req *http.Request) {
	eventDataStores, err := r.plugin.EventDataStores(req.Context(), sqlds.Options{})
	routes.SendResources(rw, eventDataStores, err)
}

func (r *CtlResourceHandler) Routes() map[string]func(http.ResponseWriter, *http.Request) {
	routes := r.DefaultRoutes()
	//routes["/catalogs"] = r.catalogs
	//routes["/workgroups"] = r.workgroups
	//routes["/workgroupEngineVersion"] = r.workgroupEngineVersion
	routes["/eventDataStores"] = r.eventDataStores
	return routes
}
