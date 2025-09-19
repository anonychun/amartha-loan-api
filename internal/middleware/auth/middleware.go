package auth

import (
	"slices"

	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) AuthenticateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bypassedPaths := []string{
			"/api/v1/admin/auth/signin",
		}

		if slices.Contains(bypassedPaths, c.Request().URL.Path) {
			return next(c)
		}

		cookie, err := c.Cookie(consts.CookieAdminSession)
		if err != nil {
			return consts.ErrUnauthorized
		}

		adminSession, err := m.repository.AdminSession.FindByToken(c.Request().Context(), cookie.Value)
		if err != nil {
			return consts.ErrUnauthorized
		}

		admin, err := m.repository.Admin.FindById(c.Request().Context(), adminSession.AdminId.String())
		if err != nil {
			return consts.ErrUnauthorized
		}

		ctx := current.SetAdmin(c.Request().Context(), admin)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func (m *Middleware) AuthenticateBorrower(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bypassedPaths := []string{
			"/api/v1/borrower/auth/signup",
			"/api/v1/borrower/auth/signin",
		}

		if slices.Contains(bypassedPaths, c.Request().URL.Path) {
			return next(c)
		}

		cookie, err := c.Cookie(consts.CookieBorrowerSession)
		if err != nil {
			return consts.ErrUnauthorized
		}

		borrowerSession, err := m.repository.BorrowerSession.FindByToken(c.Request().Context(), cookie.Value)
		if err != nil {
			return consts.ErrUnauthorized
		}

		borrower, err := m.repository.Borrower.FindById(c.Request().Context(), borrowerSession.BorrowerId.String())
		if err != nil {
			return consts.ErrUnauthorized
		}

		ctx := current.SetBorrower(c.Request().Context(), borrower)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func (m *Middleware) AuthenticateInvestor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bypassedPaths := []string{
			"/api/v1/investor/auth/signup",
			"/api/v1/investor/auth/signin",
		}

		if slices.Contains(bypassedPaths, c.Request().URL.Path) {
			return next(c)
		}

		cookie, err := c.Cookie(consts.CookieInvestorSession)
		if err != nil {
			return consts.ErrUnauthorized
		}

		investorSession, err := m.repository.InvestorSession.FindByToken(c.Request().Context(), cookie.Value)
		if err != nil {
			return consts.ErrUnauthorized
		}

		investor, err := m.repository.Investor.FindById(c.Request().Context(), investorSession.InvestorId.String())
		if err != nil {
			return consts.ErrUnauthorized
		}

		ctx := current.SetInvestor(c.Request().Context(), investor)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
