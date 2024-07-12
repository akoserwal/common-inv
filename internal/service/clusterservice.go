package service

import (
	pb "common-inv/api/inventory/v1"
	"common-inv/internal/biz"
	"common-inv/internal/models"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ClusterService struct {
	pb.UnimplementedClusterServiceServer
	clusterUsecase *biz.ClusterUsecase
}

func NewClusterServiceService(usecase *biz.ClusterUsecase) *ClusterService {
	return &ClusterService{clusterUsecase: usecase}
}

func (s *ClusterService) ListClusters(ctx context.Context, req *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	return &pb.ListClustersResponse{}, nil
}
func (s *ClusterService) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.GetClusterResponse, error) {

	return &pb.GetClusterResponse{}, nil
}
func (s *ClusterService) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*pb.CreateClusterResponse, error) {
	log.Info(req.Cluster.ClusterCommon)

	m := models.ConvertRequestToClusterIn(req)

	err := s.clusterUsecase.Create(ctx, &m)

	return &pb.CreateClusterResponse{}, err
}
func TimestampToTime(ts *timestamppb.Timestamp) (time.Time, error) {
	if ts == nil {
		return time.Time{}, nil
	}
	return ts.AsTime(), ts.CheckValid()
}

func (s *ClusterService) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*pb.UpdateClusterResponse, error) {
	return &pb.UpdateClusterResponse{}, nil
}
func (s *ClusterService) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*pb.DeleteClusterResponse, error) {
	return &pb.DeleteClusterResponse{}, nil
}
