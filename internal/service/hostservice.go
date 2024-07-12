package service

import (
	"context"

	pb "common-inv/api/inventory/v1"
)

type HostService struct {
	pb.UnimplementedHostServiceServer
}

func NewHostServiceService() *HostService {
	return &HostService{}
}

func (s *HostService) ListHosts(ctx context.Context, req *pb.ListHostsRequest) (*pb.ListHostsResponse, error) {
	return &pb.ListHostsResponse{}, nil
}
func (s *HostService) GetHost(ctx context.Context, req *pb.GetHostRequest) (*pb.GetHostResponse, error) {
	return &pb.GetHostResponse{}, nil
}
func (s *HostService) CreateHost(ctx context.Context, req *pb.CreateHostRequest) (*pb.CreateHostResponse, error) {
	return &pb.CreateHostResponse{}, nil
}
func (s *HostService) UpdateHost(ctx context.Context, req *pb.UpdateHostRequest) (*pb.UpdateHostResponse, error) {
	return &pb.UpdateHostResponse{}, nil
}
func (s *HostService) DeleteHost(ctx context.Context, req *pb.DeleteHostRequest) (*pb.DeleteHostResponse, error) {
	return &pb.DeleteHostResponse{}, nil
}
