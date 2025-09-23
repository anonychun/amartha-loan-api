package loan

import (
	"testing"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/google/uuid"
	"github.com/samber/do/v2"
)

func TestUsecase_Invest(t *testing.T) {
	repo := do.MustInvoke[*repository.Repository](bootstrap.Injector)

	usecase := &Usecase{
		repository: repo,
	}

	t.Run("should invest in loan successfully", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower",
			EmailAddress: "borrower_invest@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash borrower password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		testInvestor := &entity.Investor{
			Name:         "Test Investor",
			EmailAddress: "investor@example.com",
		}
		err = testInvestor.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor password: %v", err)
		}

		err = repo.Investor.Create(ctx, testInvestor)
		if err != nil {
			t.Fatalf("Failed to create test investor: %v", err)
		}

		ctx = current.SetInvestor(ctx, testInvestor)

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      1000000, // 1 million
			Status:               entity.LoanStatusApproved,
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		investmentAmount := int64(500000) // 500k, half of the loan
		req := InvestRequest{
			Id:     testLoan.Id.String(),
			Amount: investmentAmount,
		}

		result, err := usecase.Invest(ctx, req)
		if err != nil {
			t.Fatalf("Invest failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result should not be nil")
		}

		if result.Id != testLoan.Id.String() {
			t.Errorf("Expected loan Id %s, got %s", testLoan.Id.String(), result.Id)
		}

		if result.PrincipalAmount != testLoan.PrincipalAmount {
			t.Errorf("Expected PrincipalAmount %d, got %d", testLoan.PrincipalAmount, result.PrincipalAmount)
		}

		if result.Status != string(entity.LoanStatusApproved) {
			t.Errorf("Expected Status %s, got %s", string(entity.LoanStatusApproved), result.Status)
		}

		savedLoan, err := repo.Loan.FindById(ctx, testLoan.Id.String())
		if err != nil {
			t.Fatalf("Failed to fetch saved loan: %v", err)
		}

		if savedLoan.Status != entity.LoanStatusApproved {
			t.Errorf("Expected saved loan Status %s, got %s", string(entity.LoanStatusApproved), string(savedLoan.Status))
		}

		totalInvestment, err := repo.Investment.SumOfAmountsByLoanId(ctx, testLoan.Id)
		if err != nil {
			t.Fatalf("Failed to get total investment: %v", err)
		}

		if totalInvestment != investmentAmount {
			t.Errorf("Expected total investment %d, got %d", investmentAmount, totalInvestment)
		}
	})

	t.Run("should fully fund loan and update status to invested", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower Full Fund",
			EmailAddress: "borrower_full@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash borrower password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		testInvestor := &entity.Investor{
			Name:         "Test Investor Full",
			EmailAddress: "investor_full@example.com",
		}
		err = testInvestor.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor password: %v", err)
		}

		err = repo.Investor.Create(ctx, testInvestor)
		if err != nil {
			t.Fatalf("Failed to create test investor: %v", err)
		}

		ctx = current.SetInvestor(ctx, testInvestor)

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      1000000, // 1 million
			Status:               entity.LoanStatusApproved,
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		req := InvestRequest{
			Id:     testLoan.Id.String(),
			Amount: testLoan.PrincipalAmount, // Full amount
		}

		_, err = usecase.Invest(ctx, req)
		if err != nil {
			t.Fatalf("Invest failed: %v", err)
		}

		savedLoan, err := repo.Loan.FindById(ctx, testLoan.Id.String())
		if err != nil {
			t.Fatalf("Failed to fetch saved loan: %v", err)
		}

		if savedLoan.Status != entity.LoanStatusInvested {
			t.Errorf("Expected saved loan Status %s, got %s", string(entity.LoanStatusInvested), string(savedLoan.Status))
		}

		totalInvestment, err := repo.Investment.SumOfAmountsByLoanId(ctx, testLoan.Id)
		if err != nil {
			t.Fatalf("Failed to get total investment: %v", err)
		}

		if totalInvestment != testLoan.PrincipalAmount {
			t.Errorf("Expected total investment %d, got %d", testLoan.PrincipalAmount, totalInvestment)
		}
	})

	t.Run("should handle loan not found", func(t *testing.T) {
		ctx := t.Context()
		testInvestor := &entity.Investor{
			Name:         "Test Investor Not Found",
			EmailAddress: "investor_notfound@example.com",
		}
		err := testInvestor.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor password: %v", err)
		}

		err = repo.Investor.Create(ctx, testInvestor)
		if err != nil {
			t.Fatalf("Failed to create test investor: %v", err)
		}

		ctx = current.SetInvestor(ctx, testInvestor)

		nonExistentId := uuid.New().String()

		req := InvestRequest{
			Id:     nonExistentId,
			Amount: 100000,
		}

		_, err = usecase.Invest(ctx, req)

		if err != consts.ErrLoanNotFound {
			t.Errorf("Expected ErrLoanNotFound, got: %v", err)
		}
	})

	t.Run("should handle loan not available for investment", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower Not Available",
			EmailAddress: "borrower_notavailable@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash borrower password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		testInvestor := &entity.Investor{
			Name:         "Test Investor Not Available",
			EmailAddress: "investor_notavailable@example.com",
		}
		err = testInvestor.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor password: %v", err)
		}

		err = repo.Investor.Create(ctx, testInvestor)
		if err != nil {
			t.Fatalf("Failed to create test investor: %v", err)
		}

		ctx = current.SetInvestor(ctx, testInvestor)

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      1000000,
			Status:               entity.LoanStatusProposed, // Not approved
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		req := InvestRequest{
			Id:     testLoan.Id.String(),
			Amount: 100000,
		}

		_, err = usecase.Invest(ctx, req)

		if err != consts.ErrLoanNotAvailableForInvestment {
			t.Errorf("Expected ErrLoanNotAvailableForInvestment, got: %v", err)
		}
	})

	t.Run("should handle investment amount exceeds available amount", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower Exceeds",
			EmailAddress: "borrower_exceeds@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash borrower password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		investor1 := &entity.Investor{
			Name:         "Test Investor 1",
			EmailAddress: "investor1_exceeds@example.com",
		}
		err = investor1.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor1 password: %v", err)
		}

		err = repo.Investor.Create(ctx, investor1)
		if err != nil {
			t.Fatalf("Failed to create investor1: %v", err)
		}

		investor2 := &entity.Investor{
			Name:         "Test Investor 2",
			EmailAddress: "investor2_exceeds@example.com",
		}
		err = investor2.HashPassword("investor123")
		if err != nil {
			t.Fatalf("Failed to hash investor2 password: %v", err)
		}

		err = repo.Investor.Create(ctx, investor2)
		if err != nil {
			t.Fatalf("Failed to create investor2: %v", err)
		}

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      1000000, // 1 million
			Status:               entity.LoanStatusApproved,
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		existingInvestment := &entity.Investment{
			LoanId:     testLoan.Id,
			InvestorId: investor1.Id,
			Amount:     900000, // 900k out of 1M
		}

		err = repo.Investment.Create(ctx, existingInvestment)
		if err != nil {
			t.Fatalf("Failed to create existing investment: %v", err)
		}

		ctx = current.SetInvestor(ctx, investor2)

		req := InvestRequest{
			Id:     testLoan.Id.String(),
			Amount: 200000, // 200k, but only 100k is available
		}

		_, err = usecase.Invest(ctx, req)

		if err != consts.ErrInvestmentAmountExceedsAvailableAmount {
			t.Errorf("Expected ErrInvestmentAmountExceedsAvailableAmount, got: %v", err)
		}
	})
}
