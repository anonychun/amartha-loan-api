package consts

import (
	"net/http"

	"github.com/anonychun/amartha-loan-api/internal/api"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound

	ErrUnauthorized                  = &api.Error{Status: http.StatusUnauthorized, Errors: "You are not allowed to perform this action"}
	ErrInvalidCredentials            = &api.Error{Status: http.StatusUnauthorized, Errors: "Invalid email or password"}
	ErrEmailAddressAlreadyRegistered = &api.Error{Status: http.StatusConflict, Errors: "Email address already registered"}

	ErrAdminNotFound = &api.Error{Status: http.StatusNotFound, Errors: "Admin not found"}

	ErrLoanNotFound                           = &api.Error{Status: http.StatusNotFound, Errors: "Loan not found"}
	ErrInvalidLoanState                       = &api.Error{Status: http.StatusUnprocessableEntity, Errors: "Invalid loan state"}
	ErrLoanNotAvailableForInvestment          = &api.Error{Status: http.StatusUnprocessableEntity, Errors: "Loan is not available for investment"}
	ErrInvestmentAmountExceedsAvailableAmount = &api.Error{Status: http.StatusUnprocessableEntity, Errors: "Investment amount exceeds available amount"}
)
