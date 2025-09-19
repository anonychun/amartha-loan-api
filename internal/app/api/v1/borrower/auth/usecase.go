package auth

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
)

func (u *Usecase) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	emailAddressExists, err := u.repository.Borrower.ExistsByEmailAddress(ctx, req.EmailAddress)
	if err != nil {
		return nil, err
	}

	if emailAddressExists {
		return nil, consts.ErrEmailAddressAlreadyRegistered
	}

	borrower := &entity.Borrower{
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	}

	err = borrower.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	res := &SignUpResponse{}
	err = u.repository.Transaction(ctx, func(ctx context.Context) error {
		err := u.repository.Borrower.Create(ctx, borrower)
		if err != nil {
			return err
		}

		borrowerSession := &entity.BorrowerSession{
			BorrowerId: borrower.Id,
			IpAddress:  req.IpAddress,
			UserAgent:  req.UserAgent,
		}
		borrowerSession.GenerateToken()

		err = u.repository.BorrowerSession.Create(ctx, borrowerSession)
		if err != nil {
			return err
		}

		res.Token = borrowerSession.Token
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	borrower, err := u.repository.Borrower.FindByEmailAddress(ctx, req.EmailAddress)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = borrower.ComparePassword(req.Password)
	if err != nil {
		return nil, consts.ErrInvalidCredentials
	}

	borrowerSession := &entity.BorrowerSession{
		BorrowerId: borrower.Id,
		IpAddress:  req.IpAddress,
		UserAgent:  req.UserAgent,
	}
	borrowerSession.GenerateToken()

	err = u.repository.BorrowerSession.Create(ctx, borrowerSession)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{Token: borrowerSession.Token}, nil
}

func (u *Usecase) SignOut(ctx context.Context, req SignOutRequest) error {
	borrowerSession, err := u.repository.BorrowerSession.FindByToken(ctx, req.Token)
	if err == consts.ErrRecordNotFound {
		return consts.ErrUnauthorized
	} else if err != nil {
		return err
	}

	err = u.repository.BorrowerSession.DeleteById(ctx, borrowerSession.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) Me(ctx context.Context) (*MeResponse, error) {
	borrower := current.Borrower(ctx)
	if borrower == nil {
		return nil, consts.ErrUnauthorized
	}

	res := &MeResponse{}
	res.Borrower.Id = borrower.Id.String()
	res.Borrower.EmailAddress = borrower.EmailAddress

	return res, nil
}
