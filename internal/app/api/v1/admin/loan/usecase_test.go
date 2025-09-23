package loan

import (
	"bytes"
	"mime/multipart"
	"testing"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/consts"
	"github.com/anonychun/amartha-loan-api/internal/current"
	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/anonychun/amartha-loan-api/internal/storage"
	"github.com/google/uuid"
	"github.com/samber/do/v2"
)

func createTestFileHeader(filename string, content []byte) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("proofOfVisit", filename)
	part.Write(content)
	writer.Close()

	form, _ := multipart.NewReader(body, writer.Boundary()).ReadForm(int64(len(content)) + 1024)
	return form.File["proofOfVisit"][0]
}

func TestUsecase_Approve(t *testing.T) {
	repo := do.MustInvoke[*repository.Repository](bootstrap.Injector)
	s3 := do.MustInvoke[*storage.S3](bootstrap.Injector)

	usecase := &Usecase{
		repository: repo,
		s3:         s3,
	}

	t.Run("should approve loan successfully", func(t *testing.T) {
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

		testAdmin := &entity.Admin{
			Name:         "Test Admin",
			EmailAddress: "admin@example.com",
		}
		err = testAdmin.HashPassword("admin123")
		if err != nil {
			t.Fatalf("Failed to hash admin password: %v", err)
		}

		err = repo.Admin.Create(ctx, testAdmin)
		if err != nil {
			t.Fatalf("Failed to create test admin: %v", err)
		}

		ctx = current.SetAdmin(ctx, testAdmin)

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      1000000,
			Status:               entity.LoanStatusProposed,
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		testFileContent := []byte("test proof of visit content")
		proofOfVisit := createTestFileHeader("proof_of_visit.pdf", testFileContent)

		req := ApproveRequest{
			Id:           testLoan.Id.String(),
			ProofOfVisit: proofOfVisit,
		}

		result, err := usecase.Approve(ctx, req)
		if err != nil {
			t.Fatalf("Approve failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result should not be nil")
		}

		if result.Status != string(entity.LoanStatusApproved) {
			t.Errorf("Expected Status %s, got %s", string(entity.LoanStatusApproved), result.Status)
		}

		if result.Id != testLoan.Id.String() {
			t.Errorf("Expected loan Id %s, got %s", testLoan.Id.String(), result.Id)
		}

		savedLoan, err := repo.Loan.FindById(ctx, testLoan.Id.String())
		if err != nil {
			t.Fatalf("Failed to fetch saved loan: %v", err)
		}

		if savedLoan.Status != entity.LoanStatusApproved {
			t.Errorf("Expected saved loan Status %s, got %s", string(entity.LoanStatusApproved), string(savedLoan.Status))
		}
	})

	t.Run("should handle loan not found", func(t *testing.T) {
		ctx := t.Context()
		testAdmin := &entity.Admin{
			Name:         "Test Admin Not Found",
			EmailAddress: "admin_notfound@example.com",
		}
		err := testAdmin.HashPassword("admin123")
		if err != nil {
			t.Fatalf("Failed to hash admin password: %v", err)
		}

		err = repo.Admin.Create(ctx, testAdmin)
		if err != nil {
			t.Fatalf("Failed to create test admin: %v", err)
		}

		ctx = current.SetAdmin(ctx, testAdmin)

		nonExistentId := uuid.New().String()
		testFileContent := []byte("test proof content")
		proofOfVisit := createTestFileHeader("proof.pdf", testFileContent)

		req := ApproveRequest{
			Id:           nonExistentId,
			ProofOfVisit: proofOfVisit,
		}

		_, err = usecase.Approve(ctx, req)

		if err != consts.ErrLoanNotFound {
			t.Errorf("Expected ErrLoanNotFound, got: %v", err)
		}
	})

	t.Run("should handle invalid loan state", func(t *testing.T) {
		ctx := t.Context()
		testBorrower := &entity.Borrower{
			Name:         "Test Borrower Invalid State",
			EmailAddress: "invalid_state@example.com",
		}
		err := testBorrower.HashPassword("password123")
		if err != nil {
			t.Fatalf("Failed to hash password: %v", err)
		}

		err = repo.Borrower.Create(ctx, testBorrower)
		if err != nil {
			t.Fatalf("Failed to create test borrower: %v", err)
		}

		testAdmin := &entity.Admin{
			Name:         "Test Admin Invalid State",
			EmailAddress: "admin_invalid@example.com",
		}
		err = testAdmin.HashPassword("admin123")
		if err != nil {
			t.Fatalf("Failed to hash admin password: %v", err)
		}

		err = repo.Admin.Create(ctx, testAdmin)
		if err != nil {
			t.Fatalf("Failed to create test admin: %v", err)
		}

		ctx = current.SetAdmin(ctx, testAdmin)

		testLoan := &entity.Loan{
			BorrowerId:           testBorrower.Id,
			PrincipalAmount:      2000000,
			Status:               entity.LoanStatusApproved, // Already approved
			BorrowerInterestRate: consts.LoanBorrowerInterestRate,
			InvestorRoiRate:      consts.LoanInvestorRoiRate,
		}

		err = repo.Loan.Create(ctx, testLoan)
		if err != nil {
			t.Fatalf("Failed to create test loan: %v", err)
		}

		testFileContent := []byte("test proof content")
		proofOfVisit := createTestFileHeader("proof.pdf", testFileContent)

		req := ApproveRequest{
			Id:           testLoan.Id.String(),
			ProofOfVisit: proofOfVisit,
		}

		_, err = usecase.Approve(ctx, req)

		if err != consts.ErrInvalidLoanState {
			t.Errorf("Expected ErrInvalidLoanState, got: %v", err)
		}
	})
}
