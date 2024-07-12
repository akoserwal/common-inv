package models

import (
	pb "common-inv/api/inventory/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Convert pb.CreateClusterRequest to models.ClusterIn
func ConvertRequestToClusterIn(req *pb.CreateClusterRequest) ClusterIn {
	tt := req.Cluster.Metadata.Tags
	return ClusterIn{
		Metadata: ResourceIn{
			DisplayName:     req.Cluster.Metadata.DisplayName,
			ReporterTime:    req.Cluster.Metadata.ReporterTime.AsTime(),
			ConsoleHref:     req.Cluster.Metadata.ConsoleHref,
			ApiHref:         req.Cluster.Metadata.ApiHref,
			ResourceIdAlias: req.Cluster.Metadata.ResourceIdAlias,
			Tags:            convertResourceTagsOut(tt),
		},
		ClusterCommon: ClusterCommon{
			ExternalClusterId: req.Cluster.ClusterCommon.ExternalClusterId,
			ApiServer:         req.Cluster.ClusterCommon.ApiServer,
			CloudProviderId:   req.Cluster.ClusterCommon.CloudProviderId,
		},
	}
}
func convertResourceTagsOut(pbTags []*pb.ResourceTag) []ResourceTag {
	var tags []ResourceTag
	for _, pbTag := range pbTags {
		tags = append(tags, ResourceTag{
			ID:         IDType(pbTag.Id),
			ResourceID: IDType(pbTag.ResourceId),
			Namespace:  pbTag.Namespace,
			Key:        pbTag.Key,
			Value:      pbTag.Value,
		})
	}
	return tags
}

func ConvertClusterInToCluster(input *ClusterIn) Cluster {
	return Cluster{
		Metadata: Resource{
			DisplayName:  input.Metadata.DisplayName,
			ResourceType: "k8s-cluster", // Assuming ResourceType is "Cluster"
			Tags:         input.Metadata.Tags,
			Reporters: []Reporter{
				{
					Name:    "",
					Type:    "",
					URL:     "",
					Created: input.Metadata.ReporterTime,
					Updated: input.Metadata.ReporterTime,
				},
			},
		},
		ClusterCommon: input.ClusterCommon,
	}
}

// Convert models.ClusterIn to pb.CreateClusterRequest
func ConvertClusterInToRequest(clusterIn ClusterIn) *pb.CreateClusterRequest {
	return &pb.CreateClusterRequest{
		Cluster: &pb.ClusterIn{
			Metadata: &pb.ResourceIn{
				DisplayName:     clusterIn.Metadata.DisplayName,
				ReporterTime:    timestamppb.New(clusterIn.Metadata.ReporterTime),
				ConsoleHref:     clusterIn.Metadata.ConsoleHref,
				ApiHref:         clusterIn.Metadata.ApiHref,
				ResourceIdAlias: clusterIn.Metadata.ResourceIdAlias,
				Tags:            convertResourceTagsIn(clusterIn.Metadata.Tags),
			},
			ClusterCommon: &pb.ClusterCommon{
				ExternalClusterId: clusterIn.ExternalClusterId,
				ApiServer:         clusterIn.ApiServer,
				CloudProviderId:   clusterIn.CloudProviderId,
			},
		},
	}
}

// Helper function to convert []models.ResourceTag to []*pb.ResourceTag
func convertResourceTagsIn(tags []ResourceTag) []*pb.ResourceTag {
	var pbTags []*pb.ResourceTag
	for _, tag := range tags {
		pbTags = append(pbTags, &pb.ResourceTag{
			Id:         int64(tag.ID),
			ResourceId: int64(tag.ResourceID),
			Namespace:  tag.Namespace,
			Key:        tag.Key,
			Value:      tag.Value,
		})
	}
	return pbTags
}

func convertReporters(reporters []Reporter) []*pb.Reporter {
	var pbReporters []*pb.Reporter
	for _, reporter := range reporters {
		pbReporters = append(pbReporters, &pb.Reporter{
			Name:            reporter.Name,
			CreatedAt:       timestamppb.New(reporter.Created),
			LastUpdatedAt:   timestamppb.New(reporter.Updated),
			Type:            reporter.Type,
			Url:             reporter.URL,
			ConsoleHref:     reporter.ConsoleHref,
			ApiHref:         reporter.ApiHref,
			ResourceIdAlias: reporter.ResourceIdAlias,
		})
	}
	return pbReporters
}

// Helper function to convert []models.ResourceTag to []*pb.ResourceTag
func convertResourceTags(tags []ResourceTag) []*pb.ResourceTag {
	var pbTags []*pb.ResourceTag
	for _, tag := range tags {
		pbTags = append(pbTags, &pb.ResourceTag{
			Id:         int64(tag.ID),
			ResourceId: int64(tag.ResourceID),
			Namespace:  tag.Namespace,
			Key:        tag.Key,
			Value:      tag.Value,
		})
	}
	return pbTags
}

// Helper function to convert []models.Reporter to []*pb.Reporter
// Convert models.ClusterOut to pb.CreateClusterResponse
func ConvertClusterOutToResponse(resp *ClusterOut) *pb.CreateClusterResponse {
	return &pb.CreateClusterResponse{
		Cluster: &pb.ClusterOut{
			Metadata: &pb.ResourceOut{
				Resource: &pb.Resource{
					Id:           int64(resp.Metadata.ID),
					DisplayName:  resp.Metadata.DisplayName,
					Tags:         convertResourceTags(resp.Metadata.Tags),
					ResourceType: resp.Metadata.ResourceType,
					Reporters:    convertReporters(resp.Metadata.Reporters),
					CreatedAt:    timestamppb.New(resp.Metadata.CreatedAt),
					UpdatedAt:    timestamppb.New(resp.Metadata.UpdatedAt),
				},
				Href: resp.Metadata.Href,
			},
			ClusterCommon: &pb.ClusterCommon{
				ExternalClusterId: resp.ExternalClusterId,
				ApiServer:         resp.ApiServer,
				CloudProviderId:   resp.CloudProviderId,
			},
		},
	}
}
