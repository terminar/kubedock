package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/joyrex2001/kubedock/internal/container"
	"github.com/joyrex2001/kubedock/internal/kubernetes"
	"github.com/joyrex2001/kubedock/internal/server/httputil"
)

// containerRouter is the object that facilitate all container
// related API endpoints.
type Router struct {
	factory    container.Factory
	kubernetes kubernetes.Kubernetes
}

// New will instantiate a containerRouter object.
func New(version int, router *gin.Engine, factory container.Factory, kube kubernetes.Kubernetes) *Router {
	vprefix := ""
	if version != 0 {
		vprefix = fmt.Sprintf("/v1.%d", version)
	}
	cr := &Router{
		factory:    factory,
		kubernetes: kube,
	}
	cr.initRoutes(vprefix, router)
	return cr
}

// initRoutes will add all suported routes.
func (cr *Router) initRoutes(version string, router *gin.Engine) {
	router.POST(version+"/containers/create", cr.ContainerCreate)
	router.POST(version+"/containers/:id/start", cr.ContainerStart)
	router.GET(version+"/containers/:id/json", cr.ContainerInfo)
	router.DELETE(version+"/containers/:id", cr.ContainerDelete)
	router.POST(version+"/containers/:id/exec", cr.ContainerExec)
	router.GET(version+"/containers/:id/logs", cr.ContainerLogs)
	router.POST(version+"/exec/:id/start", cr.ExecStart)
	router.GET(version+"/exec/:id/json", cr.ExecInfo)
	router.PUT(version+"/containers/:id/archive", cr.PutArchive)
	router.GET(version+"/networks", cr.NetworksList)
	router.POST(version+"/networks/create", cr.NetworksCreate)
	router.GET(version+"/images/json", cr.ImageList)
	router.POST(version+"/images/create", cr.ImageCreate)
	router.GET(version+"/images/:image/*json", cr.ImageJSON)
	router.GET(version+"/_ping", cr.Ping)
	router.HEAD(version+"/_ping", cr.Ping)
	router.GET(version+"/info", cr.Info)
	router.GET(version+"/version", cr.Version)
	router.GET(version+"/healthz", cr.Healthz)

	// not supported at the moment
	router.POST(version+"/containers/:id/stop", httputil.NotImplemented)
	router.POST(version+"/containers/:id/kill", httputil.NotImplemented)
	router.GET(version+"/containers/json", httputil.NotImplemented)
	router.GET(version+"/containers/:id/top", httputil.NotImplemented)
	router.GET(version+"/containers/:id/changes", httputil.NotImplemented)
	router.GET(version+"/containers/:id/export", httputil.NotImplemented)
	router.GET(version+"/containers/:id/stats", httputil.NotImplemented)
	router.POST(version+"/containers/:id/resize", httputil.NotImplemented)
	router.POST(version+"/containers/:id/restart", httputil.NotImplemented)
	router.POST(version+"/containers/:id/update", httputil.NotImplemented)
	router.POST(version+"/containers/:id/rename", httputil.NotImplemented)
	router.POST(version+"/containers/:id/pause", httputil.NotImplemented)
	router.POST(version+"/containers/:id/unpause", httputil.NotImplemented)
	router.POST(version+"/containers/:id/attach", httputil.NotImplemented)
	router.GET(version+"/containers/:id/attach/ws", httputil.NotImplemented)
	router.POST(version+"/containers/:id/wait", httputil.NotImplemented)
	router.HEAD(version+"/containers/:id/archive", httputil.NotImplemented)
	router.GET(version+"/containers/:id/archive", httputil.NotImplemented)
	router.POST(version+"/containers/prune", httputil.NotImplemented)
	router.GET(version+"/networks/reaper_default", httputil.NotImplemented)
}