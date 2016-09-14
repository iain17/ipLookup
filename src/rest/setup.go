package rest

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"github.com/iain17/ipLookup/src/environment"
	"github.com/iain17/ipLookup/src/rest/lookup"
	"github.com/iain17/ipLookup/src/utils/logger"
)

func Start() {
	lookup := lookup.New()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	//api.Use(&rest.CorsMiddleware{
	//	RejectNonCorsRequests: false,
	//	OriginValidator: func(origin string, request *rest.Request) bool {
	//		if origin != environment.Env.Config.HTTP.AppUrl+environment.Env.Config.HTTP.AppAddress {
	//			logger.Warningf("Rejecting request because origin: %s != appAddress: %s", origin, environment.Env.Config.HTTP.AppAddress)
	//			return false
	//		}
	//		return true
	//	},
	//	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowedHeaders: []string{
	//		"Accept", "Content-Type", "X-Custom-Header", "Origin", "Authorization"},
	//	AccessControlAllowCredentials: true,
	//	AccessControlMaxAge:           3600,
	//})
	//
	//session_middleware := &SessionMiddleware{
	//	Timeout:    time.Hour,
	//}

	router, err := rest.MakeRouter(
		//Lookup
		rest.Get("/lookup/ip/:ip", lookup.GetIp),
	)

	if err != nil {
		logger.Fatal(err)
	}
	api.SetApp(router)

	http.ListenAndServe(environment.Env.Config.HTTP.BindingAddress, api.MakeHandler())
	logger.Infof("[HTTP]: rest api at: %s", environment.Env.Config.HTTP.BindingAddress)
}