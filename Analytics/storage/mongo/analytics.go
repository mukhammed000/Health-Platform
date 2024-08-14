package mongo

import (
	pb "analytics/genproto/analytics"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalyticsStorage struct {
	db *mongo.Database
}

func NewAnalyticsStorage(db *mongo.Database) *AnalyticsStorage {
	return &AnalyticsStorage{db: db}
}

// Medical Records
func (s *AnalyticsStorage) AddMedicalRecord(req *pb.AddMedicalRecordReq) (*pb.AddMedicalRecordRes, error) {
	collection := s.db.Collection("medical_records")

	record := bson.M{
		"_id":         req.Id,
		"user_id":     req.UserId,
		"record_type": req.RecordType,
		"record_date": req.RecordDate,
		"description": req.Description,
		"doctor_id":   req.DoctorId,
		"attachments": req.Attachments,
		"created_at":  req.CreatedAt,
		"updated_at":  req.UpdatedAt,
	}

	_, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return nil, fmt.Errorf("failed to insert medical record: %v", err)
	}

	res := &pb.AddMedicalRecordRes{
		Id:          req.Id,
		UserId:      req.UserId,
		RecordType:  req.RecordType,
		RecordDate:  req.RecordDate,
		Description: req.Description,
		DoctorId:    req.DoctorId,
		Attachments: req.Attachments,
		CreatedAt:   req.CreatedAt,
	}

	return res, nil
}

func (s *AnalyticsStorage) GetMedicalRecord(req *pb.GetMedicalRecordsReq) (*pb.GetMedicalRecordsRes, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"user_id": req.UserId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find medical records: %v", err)
	}
	defer cursor.Close(context.TODO())

	var records []*pb.AddMedicalRecordRes

	for cursor.Next(context.TODO()) {
		var record bson.M
		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("failed to decode medical record: %v", err)
		}

		records = append(records, &pb.AddMedicalRecordRes{
			Id:          record["_id"].(string),
			UserId:      record["user_id"].(string),
			RecordType:  record["record_type"].(string),
			RecordDate:  record["record_date"].(string),
			Description: record["description"].(string),
			DoctorId:    record["doctor_id"].(string),
			Attachments: toStringSlice(record["attachments"].([]interface{})),
			CreatedAt:   record["created_at"].(string),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during cursor iteration: %v", err)
	}

	res := &pb.GetMedicalRecordsRes{
		Records: records,
	}

	return res, nil
}

func toStringSlice(interfaces []interface{}) []string {
	var strings []string
	for _, i := range interfaces {
		if s, ok := i.(string); ok {
			strings = append(strings, s)
		}
	}
	return strings
}

func (s *AnalyticsStorage) GetMedicalRecordsById(req *pb.GetMedicalRecordsByIdReq) (*pb.GetMedicalRecordsByIdRes, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"user_id": req.UserId, "_id": req.RecordId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("medical record not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find medical record: %v", err)
	}

	res := &pb.GetMedicalRecordsByIdRes{
		Info: &pb.AddMedicalRecordRes{
			Id:          record["_id"].(string),
			UserId:      record["user_id"].(string),
			RecordType:  record["record_type"].(string),
			RecordDate:  record["record_date"].(string),
			Description: record["description"].(string),
			DoctorId:    record["doctor_id"].(string),
			Attachments: toStringSlice(record["attachments"].([]interface{})),
			CreatedAt:   record["created_at"].(string),
		},
	}

	return res, nil
}

func (s *AnalyticsStorage) UpdateMedicalRecord(req *pb.UpdateMedicalRecordReq) (*pb.Empty, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"_id": req.Id}

	update := bson.M{
		"$set": bson.M{
			"user_id":     req.UserId,
			"record_type": req.RecordType,
			"record_date": req.RecordDate,
			"description": req.Description,
			"doctor_id":   req.DoctorId,
			"attachments": req.Attachments,
			"updated_at":  req.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update medical record: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *AnalyticsStorage) DeleteMedicalRecord(req *pb.DeleteMedicalRecordReq) (*pb.Empty, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"_id": req.Id, "user_id": req.UserId}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete medical record: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *AnalyticsStorage) ListMedicalRecords(req *pb.ListMedicalRecordsReq) (*pb.ListMedicalRecordsRes, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"user_id": req.UserId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find medical records: %v", err)
	}
	defer cursor.Close(context.TODO())

	var records []*pb.AddMedicalRecordRes

	for cursor.Next(context.TODO()) {
		var record bson.M
		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("failed to decode medical record: %v", err)
		}

		records = append(records, &pb.AddMedicalRecordRes{
			Id:          record["_id"].(string),
			UserId:      record["user_id"].(string),
			RecordType:  record["record_type"].(string),
			RecordDate:  record["record_date"].(string),
			Description: record["description"].(string),
			DoctorId:    record["doctor_id"].(string),
			Attachments: toStringSlice(record["attachments"].([]interface{})),
			CreatedAt:   record["created_at"].(string),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during cursor iteration: %v", err)
	}

	return &pb.ListMedicalRecordsRes{
		Records: records,
	}, nil
}

// Lifestyle Data
func (s *AnalyticsStorage) AddLifestyleData(req *pb.AddLifestyleDataReq) (*pb.AddLifestyleDataRes, error) {
	collection := s.db.Collection("lifestyle_data")

	data := bson.M{
		"_id":           req.Id,
		"user_id":       req.UserId,
		"data_type":     req.DataType,
		"data_value":    req.DataValue,
		"recorded_date": req.RecordedDate,
		"created_at":    req.CreatedAt,
		"updated_at":    req.UpdatedAt,
	}

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to insert lifestyle data: %v", err)
	}

	res := &pb.AddLifestyleDataRes{
		Id:           req.Id,
		UserId:       req.UserId,
		DataType:     req.DataType,
		DataValue:    req.DataValue,
		RecordedDate: req.RecordedDate,
		CreatedAt:    req.CreatedAt,
	}

	return res, nil
}

func (s *AnalyticsStorage) GetLifestyleData(req *pb.GetLifestyleDataReq) (*pb.GetLifestyleDataRes, error) {
	collection := s.db.Collection("lifestyle_data")

	filter := bson.M{"user_id": req.UserId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find lifestyle data: %v", err)
	}
	defer cursor.Close(context.TODO())

	var data []*pb.AddLifestyleDataRes

	for cursor.Next(context.TODO()) {
		var record bson.M
		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("failed to decode lifestyle data: %v", err)
		}

		data = append(data, &pb.AddLifestyleDataRes{
			Id:           record["_id"].(string),
			UserId:       record["user_id"].(string),
			DataType:     record["data_type"].(string),
			DataValue:    record["data_value"].(string),
			RecordedDate: record["recorded_date"].(string),
			CreatedAt:    record["created_at"].(string),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during cursor iteration: %v", err)
	}

	return &pb.GetLifestyleDataRes{
		Data: data,
	}, nil
}

func (s *AnalyticsStorage) GetLifestyleDataById(req *pb.GetLifestyleDataByIdReq) (*pb.GetLifestyleDataByIdRes, error) {
	collection := s.db.Collection("lifestyle_data")

	filter := bson.M{"user_id": req.UserId, "_id": req.DataId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("lifestyle data not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find lifestyle data: %v", err)
	}

	res := &pb.GetLifestyleDataByIdRes{
		Info: &pb.AddLifestyleDataRes{
			Id:           record["_id"].(string),
			UserId:       record["user_id"].(string),
			DataType:     record["data_type"].(string),
			DataValue:    record["data_value"].(string),
			RecordedDate: record["recorded_date"].(string),
			CreatedAt:    record["created_at"].(string),
		},
	}

	return res, nil
}

func (s *AnalyticsStorage) UpdateLifestyleData(req *pb.UpdateLifestyleDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("lifestyle_data")

	filter := bson.M{"_id": req.Id, "user_id": req.UserId}

	update := bson.M{
		"$set": bson.M{
			"data_type":     req.DataType,
			"data_value":    req.DataValue,
			"recorded_date": req.RecordedDate,
			"updated_at":    req.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update lifestyle data: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *AnalyticsStorage) DeleteLifestyleData(req *pb.DeleteLifestyleDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("lifestyle_data")

	filter := bson.M{"_id": req.Id, "user_id": req.UserId}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete lifestyle data: %v", err)
	}

	return &pb.Empty{}, nil
}

// Wearable Data
func (s *AnalyticsStorage) AddWearableData(req *pb.AddWearableDataReq) (*pb.AddWearableDataRes, error) {
	collection := s.db.Collection("wearable_data")

	data := bson.M{
		"_id":                req.Id,
		"user_id":            req.UserId,
		"device_type":        req.DeviceType,
		"data_type":          req.DataType,
		"data_value":         req.DataValue,
		"recorded_timestamp": req.RecordedTimestamp,
		"created_at":         req.CreatedAt,
		"updated_at":         req.UpdatedAt,
	}

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to insert wearable data: %v", err)
	}

	res := &pb.AddWearableDataRes{
		Id:                req.Id,
		UserId:            req.UserId,
		DeviceType:        req.DeviceType,
		DataType:          req.DataType,
		DataValue:         req.DataValue,
		RecordedTimestamp: req.RecordedTimestamp,
		CreatedAt:         req.CreatedAt,
	}

	return res, nil
}

func (s *AnalyticsStorage) GetWearableData(req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error) {
	collection := s.db.Collection("wearable_data")

	filter := bson.M{"user_id": req.UserId}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find wearable data: %v", err)
	}
	defer cursor.Close(context.TODO())

	var data []*pb.AddWearableDataRes

	for cursor.Next(context.TODO()) {
		var record bson.M
		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("failed to decode wearable data: %v", err)
		}

		data = append(data, &pb.AddWearableDataRes{
			Id:                record["_id"].(string),
			UserId:            record["user_id"].(string),
			DeviceType:        record["device_type"].(string),
			DataType:          record["data_type"].(string),
			DataValue:         record["data_value"].(string),
			RecordedTimestamp: record["recorded_timestamp"].(string),
			CreatedAt:         record["created_at"].(string),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during cursor iteration: %v", err)
	}

	return &pb.GetWearableDataRes{
		Data: data,
	}, nil
}

func (s *AnalyticsStorage) GetWearableDataById(req *pb.GetWearableDataByIdReq) (*pb.GetWearableDataByIdRes, error) {
	collection := s.db.Collection("wearable_data")

	filter := bson.M{"user_id": req.UserId, "_id": req.DataId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("wearable data not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find wearable data: %v", err)
	}

	res := &pb.GetWearableDataByIdRes{
		Info: &pb.AddWearableDataRes{
			Id:                record["_id"].(string),
			UserId:            record["user_id"].(string),
			DeviceType:        record["device_type"].(string),
			DataType:          record["data_type"].(string),
			DataValue:         record["data_value"].(string),
			RecordedTimestamp: record["recorded_timestamp"].(string),
			CreatedAt:         record["created_at"].(string),
		},
	}

	return res, nil
}

func (s *AnalyticsStorage) UpdateWearableData(req *pb.UpdateWearableDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("wearable_data")

	filter := bson.M{"_id": req.Id, "user_id": req.UserId}

	update := bson.M{
		"$set": bson.M{
			"device_type":        req.DeviceType,
			"data_type":          req.DataType,
			"data_value":         req.DataValue,
			"recorded_timestamp": req.RecordedTimestamp,
			"updated_at":         req.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update wearable data: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *AnalyticsStorage) DeleteWearableData(req *pb.DeleteWearableDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("wearable_data")

	filter := bson.M{"_id": req.Id, "user_id": req.UserId}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete wearable data: %v", err)
	}

	return &pb.Empty{}, nil
}

// GenerateHealthRecommendations generates health recommendations based on the provided data.
// This method uses Gemini AI for generating personalized recommendations.
//
// Parameters:
// - req: Contains user_id, recommendation_type, description, priority, created_at, and updated_at.
//
// Returns:
//   - A response with generated health recommendations including id, user_id, recommendation_type,
//     description, priority, created_at, and updated_at.
//
// Note:
//   - Integration with Gemini AI should be implemented here. The AI will process the provided data and
//     generate appropriate health recommendations based on user-specific information and requirements.
func (s *AnalyticsStorage) GenerateHealthRecommendations(req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error) {
	// TODO: Implement integration with Gemini AI to generate health recommendations
	// Example: Call Gemini AI service with the provided data and receive generated recommendations

	// Placeholder code, to be replaced with actual implementation
	return nil, nil
}

func (s *AnalyticsStorage) GetHealthRecommendationsById(req *pb.GetHealthRecommendationsByIdReq) (*pb.GetHealthRecommendationsByIdRes, error) {
	collection := s.db.Collection("health_recommendations")

	filter := bson.M{"user_id": req.UserId, "_id": req.RecommendationId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("health recommendation not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find health recommendation: %v", err)
	}

	res := &pb.GetHealthRecommendationsByIdRes{
		Info: &pb.GenerateHealthRecommendationsRes{
			Id:                 record["_id"].(string),
			UserId:             record["user_id"].(string),
			RecommendationType: record["recommendation_type"].(string),
			Description:        record["description"].(string),
			Priority:           int32(record["priority"].(int)),
			CreatedAt:          record["created_at"].(string),
			UpdatedAt:          record["updated_at"].(string),
		},
	}

	return res, nil
}

// Health Monitoring
func (s *AnalyticsStorage) GetRealtimeHealthMonitoring(req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error) {
	return nil, nil
}

func (s *AnalyticsStorage) GetDailyHealthSummary(req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error) {
	return nil, nil
}

func (s *AnalyticsStorage) GetWeeklyHealthSummary(req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error) {
	return nil, nil
}
