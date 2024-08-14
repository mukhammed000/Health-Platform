package storage

import (
	pb "analytics/genproto/analytics"
)

type InitRoot interface {
	Analytics() AnalyticsService
}

type AnalyticsService interface {
	// Medical Records
	AddMedicalRecord(req *pb.AddMedicalRecordReq) (*pb.AddMedicalRecordRes, error)
	GetMedicalRecord(req *pb.GetMedicalRecordsReq) (*pb.GetMedicalRecordsRes, error)
	GetMedicalRecordsById(req *pb.GetMedicalRecordsByIdReq) (*pb.GetMedicalRecordsByIdRes, error)
	UpdateMedicalRecord(req *pb.UpdateMedicalRecordReq) (*pb.Empty, error)
	DeleteMedicalRecord(req *pb.DeleteMedicalRecordReq) (*pb.Empty, error)
	ListMedicalRecords(req *pb.ListMedicalRecordsReq) (*pb.ListMedicalRecordsRes, error)

	// Lifestyle Data
	AddLifestyleData(req *pb.AddLifestyleDataReq) (*pb.AddLifestyleDataRes, error)
	GetLifestyleData(req *pb.GetLifestyleDataReq) (*pb.GetLifestyleDataRes, error)
	GetLifestyleDataById(req *pb.GetLifestyleDataByIdReq) (*pb.GetLifestyleDataByIdRes, error)
	UpdateLifestyleData(req *pb.UpdateLifestyleDataReq) (*pb.Empty, error)
	DeleteLifestyleData(req *pb.DeleteLifestyleDataReq) (*pb.Empty, error)

	// Wearable Data
	AddWearableData(req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error)
	GetWearableData(req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error)
	GetWearableDataById(req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error)
	UpdateWearableData(req *pb.UpdateWearableDataReq) (*pb.Empty, error)
	DeleteWearableData(req *pb.DeleteWearableDataReq) (*pb.Empty, error)

	// Health Recommendations
	GenerateHealthRecommendations(req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error)
	GetHealthRecommendationsById(req *pb.GetHealthRecommendationsByIdReq) (*pb.GetHealthRecommendationsByIdRes, error)

	// Health Monitoring
	GetRealtimeHealthMonitoring(req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error)
	GetDailyHealthSummary(req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error)
	GetWeeklyHealthSummary(req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error)
}
