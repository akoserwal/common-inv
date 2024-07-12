package service

import (
	"context"

	pb "common-inv/api/inventory/v1"
)

type ResourceService struct {
	pb.UnimplementedResourceServiceServer
}

func NewResourceServiceService() *ResourceService {
	return &ResourceService{}
}

func (s *ResourceService) ListResources(ctx context.Context, req *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	return &pb.ListResourcesResponse{}, nil
}
func (s *ResourceService) GetResource(ctx context.Context, req *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	return &pb.GetResourceResponse{}, nil
}
