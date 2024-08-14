package service

import (
	"context"
	"log"

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

// Medical Records
func (s *AnalyticsService) AddMedicalRecord(ctx context.Context, req *pb.AddMedicalRecordReq) (*pb.AddMedicalRecordRes, error) {
	resp, err := s.stg.Analytics().AddMedicalRecord(req)
	if err != nil {
		log.Println("Error adding medical record: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetMedicalRecord(ctx context.Context, req *pb.GetMedicalRecordsReq) (*pb.GetMedicalRecordsRes, error) {
	resp, err := s.stg.Analytics().GetMedicalRecord(req)
	if err != nil {
		log.Println("Error getting medical record: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetMedicalRecordsById(ctx context.Context, req *pb.GetMedicalRecordsByIdReq) (*pb.GetMedicalRecordsByIdRes, error) {
	resp, err := s.stg.Analytics().GetMedicalRecordsById(req)
	if err != nil {
		log.Println("Error getting medical record by ID: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) UpdateMedicalRecord(ctx context.Context, req *pb.UpdateMedicalRecordReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().UpdateMedicalRecord(req)
	if err != nil {
		log.Println("Error updating medical record: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) DeleteMedicalRecord(ctx context.Context, req *pb.DeleteMedicalRecordReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().DeleteMedicalRecord(req)
	if err != nil {
		log.Println("Error deleting medical record: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) ListMedicalRecords(ctx context.Context, req *pb.ListMedicalRecordsReq) (*pb.ListMedicalRecordsRes, error) {
	resp, err := s.stg.Analytics().ListMedicalRecords(req)
	if err != nil {
		log.Println("Error listing medical records: ", err)
		return nil, err
	}
	return resp, nil
}

// Lifestyle Data
func (s *AnalyticsService) AddLifestyleData(ctx context.Context, req *pb.AddLifestyleDataReq) (*pb.AddLifestyleDataRes, error) {
	resp, err := s.stg.Analytics().AddLifestyleData(req)
	if err != nil {
		log.Println("Error adding lifestyle data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetLifestyleData(ctx context.Context, req *pb.GetLifestyleDataReq) (*pb.GetLifestyleDataRes, error) {
	resp, err := s.stg.Analytics().GetLifestyleData(req)
	if err != nil {
		log.Println("Error getting lifestyle data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetLifestyleDataById(ctx context.Context, req *pb.GetLifestyleDataByIdReq) (*pb.GetLifestyleDataByIdRes, error) {
	resp, err := s.stg.Analytics().GetLifestyleDataById(req)
	if err != nil {
		log.Println("Error getting lifestyle data by ID: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) UpdateLifestyleData(ctx context.Context, req *pb.UpdateLifestyleDataReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().UpdateLifestyleData(req)
	if err != nil {
		log.Println("Error updating lifestyle data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) DeleteLifestyleData(ctx context.Context, req *pb.DeleteLifestyleDataReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().DeleteLifestyleData(req)
	if err != nil {
		log.Println("Error deleting lifestyle data: ", err)
		return nil, err
	}
	return resp, nil
}

// Wearable Data
func (s *AnalyticsService) AddWearableData(ctx context.Context, req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error) {
	resp, err := s.stg.Analytics().AddWearableData(req)
	if err != nil {
		log.Println("Error adding wearable data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetWearableData(ctx context.Context, req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error) {
	resp, err := s.stg.Analytics().GetWearableData(req)
	if err != nil {
		log.Println("Error getting wearable data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetWearableDataById(ctx context.Context, req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error) {
	resp, err := s.stg.Analytics().GetWearableDataById(req)
	if err != nil {
		log.Println("Error getting wearable data by ID: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) UpdateWearableData(ctx context.Context, req *pb.UpdateWearableDataReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().UpdateWearableData(req)
	if err != nil {
		log.Println("Error updating wearable data: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) DeleteWearableData(ctx context.Context, req *pb.DeleteWearableDataReq) (*pb.Empty, error) {
	resp, err := s.stg.Analytics().DeleteWearableData(req)
	if err != nil {
		log.Println("Error deleting wearable data: ", err)
		return nil, err
	}
	return resp, nil
}

// Health Recommendations
func (s *AnalyticsService) GenerateHealthRecommendations(ctx context.Context, req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error) {
	resp, err := s.stg.Analytics().GenerateHealthRecommendations(req)
	if err != nil {
		log.Println("Error generating health recommendations: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetHealthRecommendationsById(ctx context.Context, req *pb.GetHealthRecommendationsByIdReq) (*pb.GetHealthRecommendationsByIdRes, error) {
	resp, err := s.stg.Analytics().GetHealthRecommendationsById(req)
	if err != nil {
		log.Println("Error getting health recommendations by ID: ", err)
		return nil, err
	}
	return resp, nil
}

// Health Monitoring
func (s *AnalyticsService) GetRealtimeHealthMonitoring(ctx context.Context, req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error) {
	resp, err := s.stg.Analytics().GetRealtimeHealthMonitoring(req)
	if err != nil {
		log.Println("Error getting real-time health monitoring: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetDailyHealthSummary(ctx context.Context, req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error) {
	resp, err := s.stg.Analytics().GetDailyHealthSummary(req)
	if err != nil {
		log.Println("Error getting daily health summary: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AnalyticsService) GetWeeklyHealthSummary(ctx context.Context, req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error) {
	resp, err := s.stg.Analytics().GetWeeklyHealthSummary(req)
	if err != nil {
		log.Println("Error getting weekly health summary: ", err)
		return nil, err
	}
	return resp, nil
}
