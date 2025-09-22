package loan

import (
	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/repository"
	"github.com/anonychun/amartha-loan-api/internal/storage"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	repository *repository.Repository
	s3         *storage.S3
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		repository: do.MustInvoke[*repository.Repository](i),
		s3:         do.MustInvoke[*storage.S3](i),
	}, nil
}

type Handler struct {
	usecase *Usecase
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		usecase: do.MustInvoke[*Usecase](i),
	}, nil
}
