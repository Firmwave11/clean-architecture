package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ramasapto/clean-architecture/controllers"
	"github.com/ramasapto/clean-architecture/middleware"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/zephinzer/ezpromhttp"
)

// route struct with value Controllers Interface
type route struct {
	ctrl controllers.Controllers
}

// Router represent the Router contract
type Router interface {
	Router(port string)
}

/*NewRouter will create an object that represent the Router interface (Router)
 * @parameter
 * c - controllers Interface
 *
 * @represent
 * interface Router
 *
 * @return
 * struct route with value Controllers Interface
 */
func NewRouter(c controllers.Controllers) Router {
	return &route{ctrl: c}
}

func (c *route) Router(port string) {
	router := chi.NewRouter()

	/* 	router.Use(middleware.RecordMiddleware) */

	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler)

	})

	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckAPIKey)
	})

	// group router if need to check session
	/* 	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckSession)
	}) */

	router.MethodNotAllowed(middleware.NotAllowed)
	router.NotFound(middleware.NotFound)
	logrus.Infof("Server running on port : %s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), cors.AllowAll().Handler(router)))
}
