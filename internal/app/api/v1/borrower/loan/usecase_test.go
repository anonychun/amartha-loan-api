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

func TestUsecase_Create(t *testing.T) {
	repo := do.MustInvoke[*repository.Repository](bootstrap.Injector)

	usecase := &Usecase{
		repository: repo,
	}

	t.Run("should create loan successfully", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower",
			EmailAddress: "test@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		ctx = current.SetBorrower(ctx, testBorrower)

		principalAmount := int64(1000000) // 1 million
		req := CreateRequest{
			PrincipalAmount: principalAmount,
		}

		result, err := usecase.Create(ctx, req)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result should not be nil")
		}

		if result.BorrowerId != testBorrower.Id.String() {
			t.Errorf("Expected BorrowerId %s, got %s", testBorrower.Id.String(), result.BorrowerId)
		}

		if result.PrincipalAmount != principalAmount {
			t.Errorf("Expected PrincipalAmount %d, got %d", principalAmount, result.PrincipalAmount)
		}

		if result.Status != string(entity.LoanStatusProposed) {
			t.Errorf("Expected Status %s, got %s", string(entity.LoanStatusProposed), result.Status)
		}

		savedLoan, err := repo.Loan.FindById(ctx, result.Id)
		if err != nil {
			t.Fatalf("Failed to fetch saved loan: %v", err)
		}

		if savedLoan.Id.String() != result.Id {
			t.Errorf("Expected saved loan Id %s, got %s", result.Id, savedLoan.Id.String())
		}

		if savedLoan.BorrowerId != testBorrower.Id {
			t.Errorf("Expected saved loan BorrowerId %s, got %s", testBorrower.Id.String(), savedLoan.BorrowerId.String())
		}

		if savedLoan.PrincipalAmount != principalAmount {
			t.Errorf("Expected saved loan PrincipalAmount %d, got %d", principalAmount, savedLoan.PrincipalAmount)
		}

		if savedLoan.Status != entity.LoanStatusProposed {
			t.Errorf("Expected saved loan Status %s, got %s", string(entity.LoanStatusProposed), string(savedLoan.Status))
		}

		if savedLoan.BorrowerInterestRate != consts.LoanBorrowerInterestRate {
			t.Errorf("Expected saved loan BorrowerInterestRate %f, got %f", consts.LoanBorrowerInterestRate, savedLoan.BorrowerInterestRate)
		}

		if savedLoan.InvestorRoiRate != consts.LoanInvestorRoiRate {
			t.Errorf("Expected saved loan InvestorRoiRate %f, got %f", consts.LoanInvestorRoiRate, savedLoan.InvestorRoiRate)
		}

		if savedLoan.Id == uuid.Nil {
			t.Error("Saved loan should have a valid UUID")
		}

		if savedLoan.CreatedAt.IsZero() {
			t.Error("Saved loan should have CreatedAt timestamp")
		}

		if savedLoan.UpdatedAt.IsZero() {
			t.Error("Saved loan should have UpdatedAt timestamp")
		}
	})

	t.Run("should create multiple loans for same borrower", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower Multiple",
			EmailAddress: "testmultiple@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		ctx = current.SetBorrower(ctx, testBorrower)

		req1 := CreateRequest{PrincipalAmount: 1000000}
		result1, err := usecase.Create(ctx, req1)
		if err != nil {
			t.Fatalf("Create first loan failed: %v", err)
		}

		req2 := CreateRequest{PrincipalAmount: 2000000}
		result2, err := usecase.Create(ctx, req2)
		if err != nil {
			t.Fatalf("Create second loan failed: %v", err)
		}

		if result1.Id == result2.Id {
			t.Error("Two loans should have different IDs")
		}

		if result1.BorrowerId != result2.BorrowerId {
			t.Error("Both loans should belong to the same borrower")
		}

		if result1.BorrowerId != testBorrower.Id.String() {
			t.Errorf("Expected BorrowerId %s, got %s", testBorrower.Id.String(), result1.BorrowerId)
		}

		if result1.PrincipalAmount == result2.PrincipalAmount {
			t.Error("Loans should have different principal amounts")
		}
	})
}
