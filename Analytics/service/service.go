package service

import (
	pb "analytics/genproto/analytics"
	stg "analytics/storage"
)

type AnalyticsService struct {
	stg stg.InitRoot
	pb.UnimplementedAnalyticsServiceServer
}

func NewAnalyticsService(stg stg.InitRoot) *AnalyticsService {
	return &AnalyticsService{
		stg: stg,
	}
}

