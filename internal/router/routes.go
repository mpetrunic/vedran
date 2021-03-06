package router

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"

	"github.com/NodeFactoryIo/vedran/internal/auth"
	"github.com/NodeFactoryIo/vedran/internal/controllers"
	"github.com/gorilla/mux"
)

func createRoute(route string, method string, handler http.HandlerFunc, router *mux.Router, authorized bool) {
	var r *mux.Route
	if authorized {
		r = router.Handle(route, auth.AuthMiddleware(handler))
	} else {
		r = router.Handle(route, handler)
	}
	r.Methods(method)
	r.Name(route)

	log.Debugf("Created route %s\t%s", method, route)
}

func createTrackedRoute(route string, method string, handler http.Handler, router *mux.Router) {
	r := router.Handle(route, handler)
	r.Methods(method)
	r.Name(route)

	log.Debugf("Created route %s\t%s", method, route)

}

func createRoutes(apiController *controllers.ApiController, router *mux.Router) {
	// Create a custom registry for prometheus.
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	createTrackedRoute("/", "POST", std.Handler("/", mdlw, http.HandlerFunc(apiController.RPCHandler)), router)
	createTrackedRoute("/ws", "GET", std.Handler("/ws", mdlw, http.HandlerFunc(apiController.WSHandler)), router)

	createRoute("/api/v1/nodes", "POST", apiController.RegisterHandler, router, false)
	createRoute("/api/v1/nodes/pings", "POST", apiController.PingHandler, router, true)
	createRoute("/api/v1/nodes/metrics", "PUT", apiController.SaveMetricsHandler, router, true)
	createRoute("/api/v1/stats", "POST", apiController.StatisticsHandlerAllStatsForLoadbalancer, router, false)
	createRoute("/api/v1/stats", "GET", apiController.StatisticsHandlerAllStats, router, false)
	createRoute("/api/v1/stats/node/{id}", "GET", apiController.StatisticsHandlerStatsForNode, router, false)

	createRoute("/metrics", "GET", promhttp.Handler().ServeHTTP, router, false)
}
