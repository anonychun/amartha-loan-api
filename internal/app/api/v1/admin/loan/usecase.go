package loan

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/dto"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/samber/lo"
)

func (u *Usecase) FindAll(ctx context.Context) ([]*LoanDto, error) {
	loans, err := u.repository.Loan.FindAllOrderByIdDesc(ctx)
	if err != nil {
		return nil, err
	}

	res := lo.Map(loans, func(loan *entity.Loan, _ int) *LoanDto {
		return ToLoanDto(loan)
	})

	return res, nil
}

func (u *Usecase) FindById(ctx context.Context, req FindByIdRequest) (*LoanDto, error) {
	loan, err := u.repository.Loan.FindByIdWithAgreementLetter(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrLoanNotFound
	} else if err != nil {
		return nil, err
	}

	res := ToLoanDto(loan)
	if loan.AgreementLetter != nil {
		agreementLetter, err := dto.NewAttachment(u.s3, loan.AgreementLetter)
		if err != nil {
			return nil, err
		}

		res.AgreementLetter = agreementLetter
	}

	return res, nil
}

func (u *Usecase) Approve(ctx context.Context, req ApproveRequest) (*LoanDto, error) {
	loan, err := u.repository.Loan.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrLoanNotFound
	} else if err != nil {
		return nil, err
	}

	if loan.Status != entity.LoanStatusProposed {
		return nil, consts.ErrInvalidLoanState
	}

	err = repository.Transaction(ctx, func(ctx context.Context) error {
		proofOfVisit, err := entity.NewAttachmentFromMultipartFileHeader(req.ProofOfVisit)
		if err != nil {
			return err
		}

		err = u.s3.Put(ctx, proofOfVisit.ObjectName, proofOfVisit.Content, proofOfVisit.ByteSize)
		if err != nil {
			return err
		}

		err = u.repository.Attachment.Create(ctx, proofOfVisit)
		if err != nil {
			return err
		}

		loan.Status = entity.LoanStatusApproved
		err = u.repository.Loan.Update(ctx, loan)
		if err != nil {
			return err
		}

		approval := &entity.Approval{
			LoanId:         loan.Id,
			AdminId:        current.Admin(ctx).Id,
			ProofOfVisitId: proofOfVisit.Id,
		}

		err = u.repository.Approval.Create(ctx, approval)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ToLoanDto(loan), nil
}

func (u *Usecase) Disburse(ctx context.Context, req DisburseRequest) (*LoanDto, error) {
	loan, err := u.repository.Loan.FindById(ctx, req.Id)
	if err == consts.ErrRecordNotFound {
		return nil, consts.ErrLoanNotFound
	} else if err != nil {
		return nil, err
	}

	if loan.Status != entity.LoanStatusInvested {
		return nil, consts.ErrInvalidLoanState
	}

	repository.Transaction(ctx, func(ctx context.Context) error {
		agreementLetter, err := entity.NewAttachmentFromMultipartFileHeader(req.AgreementLetter)
		if err != nil {
			return err
		}

		err = u.s3.Put(ctx, agreementLetter.ObjectName, agreementLetter.Content, agreementLetter.ByteSize)
		if err != nil {
			return err
		}

		err = u.repository.Attachment.Create(ctx, agreementLetter)
		if err != nil {
			return err
		}

		loan.AgreementLetterId = &agreementLetter.Id
		loan.Status = entity.LoanStatusDisbursed

		err = u.repository.Loan.Update(ctx, loan)
		if err != nil {
			return err
		}

		disbursement := &entity.Disbursement{
			LoanId:  loan.Id,
			AdminId: current.Admin(ctx).Id,
		}

		err = u.repository.Disbursement.Create(ctx, disbursement)
		if err != nil {
			return err
		}

		return nil
	})

	return ToLoanDto(loan), nil
}
