package auth

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/repository"
)

func (u *Usecase) SignUp(ctx context.Context, req SignUpRequest) (*SignUpResponse, error) {
	emailAddressExists, err := u.repository.Investor.ExistsByEmailAddress(ctx, req.EmailAddress)
	if err != nil {
		return nil, err
	}

	if emailAddressExists {
		return nil, consts.ErrEmailAddressAlreadyRegistered
	}

	investor := &entity.Investor{
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	}

	err = investor.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	res := &SignUpResponse{}
	err = repository.Transaction(ctx, func(ctx context.Context) error {
		err := u.repository.Investor.Create(ctx, investor)
		if err != nil {
			return err
		}

		investorSession := &entity.InvestorSession{
			InvestorId: investor.Id,
			IpAddress:  req.IpAddress,
			UserAgent:  req.UserAgent,
		}
		investorSession.GenerateToken()

		err = u.repository.InvestorSession.Create(ctx, investorSession)
		if err != nil {
			return err
		}

		res.Token = investorSession.Token
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *Usecase) SignIn(ctx context.Context, req SignInRequest) (*SignInResponse, error) {
	investor, err := u.repository.Investor.FindByEmailAddress(ctx, req.EmailAddress)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = investor.ComparePassword(req.Password)
	if err != nil {
		return nil, consts.ErrInvalidCredentials
	}

	adminSession := &entity.AdminSession{
		AdminId:   investor.Id,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}
	adminSession.GenerateToken()

	err = u.repository.AdminSession.Create(ctx, adminSession)
	if err != nil {
		return nil, err
	}

	return &SignInResponse{Token: adminSession.Token}, nil
}

func (u *Usecase) SignOut(ctx context.Context, req SignOutRequest) error {
	investorSession, err := u.repository.InvestorSession.FindByToken(ctx, req.Token)
	if err == consts.ErrRecordNotFound {
		return consts.ErrUnauthorized
	} else if err != nil {
		return err
	}

	err = u.repository.InvestorSession.DeleteById(ctx, investorSession.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) Me(ctx context.Context) (*MeResponse, error) {
	investor := current.Investor(ctx)
	if investor == nil {
		return nil, consts.ErrUnauthorized
	}

	res := &MeResponse{}
	res.Investor.Id = investor.Id.String()
	res.Investor.EmailAddress = investor.EmailAddress

	return res, nil
}
