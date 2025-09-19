package server

import (
	"github.com/anonychun/amartha-loan-api/internal/api"
	"github.com/anonychun/amartha-loan-api/internal/app"
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/samber/do"
)

func namespace(e *echo.Group, path string, f func(e *echo.Group)) {
	f(e.Group(path))
}

func routes(e *echo.Echo) error {
	m := do.MustInvoke[*middleware.Middleware](bootstrap.Injector)
	h := do.MustInvoke[*app.Handler](bootstrap.Injector)

	e.HTTPErrorHandler = api.HttpErrorHandler
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())

	apiRouter := e.Group("/api")
	namespace(apiRouter, "/v1", func(e *echo.Group) {
		namespace(e, "/admin", func(e *echo.Group) {
			e.Use(m.Auth.AuthenticateAdmin)

			e.POST("/auth/signin", h.Api.V1.Admin.Auth.SignIn)
			e.POST("/auth/signout", h.Api.V1.Admin.Auth.SignOut)
			e.GET("/auth/me", h.Api.V1.Admin.Auth.Me)

			e.GET("/admin", h.Api.V1.Admin.Admin.FindAll)
			e.GET("/admin/:id", h.Api.V1.Admin.Admin.FindById)
			e.POST("/admin", h.Api.V1.Admin.Admin.Create)
			e.PUT("/admin/:id", h.Api.V1.Admin.Admin.Update)
		})

		namespace(e, "/borrower", func(e *echo.Group) {
			e.Use(m.Auth.AuthenticateBorrower)

			e.POST("/auth/signup", h.Api.V1.Borrower.Auth.SignUp)
			e.POST("/auth/signin", h.Api.V1.Borrower.Auth.SignIn)
			e.POST("/auth/signout", h.Api.V1.Borrower.Auth.SignOut)
			e.GET("/auth/me", h.Api.V1.Borrower.Auth.Me)
		})

		namespace(e, "/investor", func(e *echo.Group) {
			e.Use(m.Auth.AuthenticateInvestor)

			e.POST("/auth/signup", h.Api.V1.Investor.Auth.SignUp)
			e.POST("/auth/signin", h.Api.V1.Investor.Auth.SignIn)
			e.POST("/auth/signout", h.Api.V1.Investor.Auth.SignOut)
			e.GET("/auth/me", h.Api.V1.Investor.Auth.Me)
		})
	})

	return nil
}
