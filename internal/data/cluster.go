package data

import (
	"common-inv/internal/biz"
	"common-inv/internal/models"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ClusterRepo = (*clusterRepo)(nil)

type clusterRepo struct {
	data *Data
	log  *log.Helper
}

func (c clusterRepo) List(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {

	panic("implement me")
}

func (c clusterRepo) Get(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func (c clusterRepo) Create(ctx context.Context, input *models.ClusterIn) error {
	cluster := models.ConvertClusterInToCluster(input)

	log.Infof("clustermodel: %v", cluster)

	connfactory := c.data.DbFactory.New()
	if err := connfactory.Create(&cluster).Error; err != nil {
		return err
	}

	return nil
}

func (c clusterRepo) Update(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func (c clusterRepo) Delete(ctx context.Context, in *models.ClusterIn) (*models.ClusterOut, error) {
	//TODO implement me
	panic("implement me")
}

func NewClusterRepo(data *Data, logger log.Logger) biz.ClusterRepo {
	return &clusterRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/beer")),
	}
}
