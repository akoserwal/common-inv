package biz

import (
	"common-inv/internal/models"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ ClusterRepo = (*ClusterUsecase)(nil)

type ClusterRepo interface {
	List(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error)
	Get(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error)
	Create(ctx context.Context, in *models.ClusterIn) error
	Update(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error)
	Delete(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error)
}

type ClusterUsecase struct {
	repo ClusterRepo
}

func (c ClusterUsecase) List(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterUsecase) Get(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterUsecase) Create(ctx context.Context, in *models.ClusterIn) error {
	return c.repo.Create(ctx, in)
}

func (c ClusterUsecase) Update(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClusterUsecase) Delete(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func NewClusterUsecase(repo ClusterRepo, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{
		repo: repo,
	}
}
