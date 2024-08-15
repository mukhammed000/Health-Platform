package mongo

import (
	pb "analytics/genproto/analytics"
	ai "analytics/storage/Ai"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		"created_at":  time.Now().String(),
		"updated_at":  time.Now().String(),
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
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	return res, nil
}

func (s *AnalyticsStorage) GetMedicalRecord(req *pb.GetMedicalRecordsReq) (*pb.GetMedicalRecordsRes, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{}

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
			UpdatedAt:   record["updated_at"].(string),
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

	filter := bson.M{"_id": req.RecordId}

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
			UpdatedAt:   record["updated_at"].(string),
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
			"updated_at":  time.Now().String(),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return &pb.Empty{Text: "Failed to update medical record", IsDone: false}, fmt.Errorf("failed to update medical record: %v", err)
	}

	if result.MatchedCount == 0 {
		return &pb.Empty{Text: "Record not found", IsDone: false}, nil
	}

	fmt.Println("--------------------->", result)

	return &pb.Empty{Text: "Record updated successfully", IsDone: true}, nil
}

func (s *AnalyticsStorage) DeleteMedicalRecord(req *pb.DeleteMedicalRecordReq) (*pb.Empty, error) {
	collection := s.db.Collection("medical_records")

	filter := bson.M{"_id": req.Id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return &pb.Empty{Text: "Failed to delete medical record", IsDone: false}, fmt.Errorf("failed to delete medical record: %v", err)
	}

	if result.DeletedCount == 0 {
		return &pb.Empty{Text: "Medical record not found", IsDone: false}, nil
	}

	return &pb.Empty{IsDone: true, Text: "Medical record successfully deleted!"}, nil
}

func (s *AnalyticsStorage) ListMedicalRecords(req *pb.ListMedicalRecordsReq) (*pb.ListMedicalRecordsRes, error) {
	collection := s.db.Collection("medical_records")

	// Filter is empty to list all medical records
	filter := bson.M{}

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
			UpdatedAt:   record["updated_at"].(string),
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
		"created_at":    time.Now().String(),
		"updated_at":    time.Now().String(),
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
		CreatedAt:    time.Now().String(),
	}

	return res, nil
}

func (s *AnalyticsStorage) GetLifestyleData(req *pb.GetLifestyleDataReq) (*pb.GetLifestyleDataRes, error) {
	collection := s.db.Collection("lifestyle_data")

	// No filter is applied; fetching all documents from the collection
	cursor, err := collection.Find(context.TODO(), bson.M{})
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

		// Extract and handle _id field
		id, ok := record["_id"].(string) // Check if _id is stored as a string
		if !ok {
			if objectID, ok := record["_id"].(primitive.ObjectID); ok {
				id = objectID.Hex() // Convert ObjectID to string
			} else {
				return nil, fmt.Errorf("failed to convert _id to string or ObjectID")
			}
		}

		// Extract other fields with default values if not present
		userID, _ := record["user_id"].(string)
		dataType, _ := record["data_type"].(string)
		dataValue, _ := record["data_value"].(string)
		recordedDate, _ := record["recorded_date"].(string)
		createdAt, _ := record["created_at"].(string)

		data = append(data, &pb.AddLifestyleDataRes{
			Id:           id,
			UserId:       userID,
			DataType:     dataType,
			DataValue:    dataValue,
			RecordedDate: recordedDate,
			CreatedAt:    createdAt,
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

	// Filter by the "_id" field
	filter := bson.M{"_id": req.DataId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("lifestyle data not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find lifestyle data: %v", err)
	}

	// Extract fields from the record
	id, ok := record["_id"].(string)
	if !ok {
		id = ""
	}

	createdAt, _ := record["created_at"].(string)
	updatedAt, _ := record["updated_at"].(string)
	recordedDate, _ := record["recorded_date"].(string)
	userID, _ := record["user_id"].(string)
	dataType, _ := record["data_type"].(string)
	dataValue, _ := record["data_value"].(string)

	// Create the response object
	res := &pb.GetLifestyleDataByIdRes{
		Info: &pb.AddLifestyleDataRes{
			Id:           id,
			UserId:       userID,
			DataType:     dataType,
			DataValue:    dataValue,
			RecordedDate: recordedDate,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
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
			"updated_at":    time.Now().String(),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return &pb.Empty{Text: "Failed to update lifestyle data", IsDone: false}, fmt.Errorf("failed to update lifestyle data: %v", err)
	}

	if result.MatchedCount == 0 {
		return &pb.Empty{Text: "Lifestyle data not found", IsDone: false}, nil
	}

	return &pb.Empty{IsDone: true, Text: "Lifestyle data updated successfully!"}, nil
}

func (s *AnalyticsStorage) DeleteLifestyleData(req *pb.DeleteLifestyleDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("lifestyle_data")

	filter := bson.M{"_id": req.Id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return &pb.Empty{Text: "Failed to delete lifestyle data", IsDone: false}, fmt.Errorf("failed to delete lifestyle data: %v", err)
	}

	if result.DeletedCount == 0 {
		return &pb.Empty{Text: "Lifestyle data not found", IsDone: false}, nil
	}

	return &pb.Empty{IsDone: true, Text: "Lifestyle data deleted successfully!"}, nil
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
		"created_at":         time.Now().String(),
		"updated_at":         time.Now().String(),
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
		CreatedAt:         data["created_at"].(string),
	}

	return res, nil
}

func (s *AnalyticsStorage) GetWearableData(req *pb.GetWearableDataReq) (*pb.GetWearableDataRes, error) {
	collection := s.db.Collection("wearable_data")

	// Assuming we need to fetch all records or filter by user_id if it's a global request
	filter := bson.M{} // Use an appropriate filter if needed, e.g., {"user_id": req.UserId}

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

	// Construct the filter to find the document by data_id and user_id
	filter := bson.M{"_id": req.DataId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("wearable data not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find wearable data: %v", err)
	}

	// Construct the response
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
			"updated_at":         time.Now().String(),
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return &pb.Empty{Text: "Failed to update wearable data", IsDone: false}, fmt.Errorf("failed to update wearable data: %v", err)
	}

	if result.MatchedCount == 0 {
		return &pb.Empty{Text: "Wearable data not found", IsDone: false}, nil
	}

	return &pb.Empty{IsDone: true, Text: "Wearable data is successfully updated!"}, nil
}

func (s *AnalyticsStorage) DeleteWearableData(req *pb.DeleteWearableDataReq) (*pb.Empty, error) {
	collection := s.db.Collection("wearable_data")

	filter := bson.M{"_id": req.Id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return &pb.Empty{Text: "Failed to delete wearable data", IsDone: false}, fmt.Errorf("failed to delete wearable data: %v", err)
	}

	if result.DeletedCount == 0 {
		return &pb.Empty{Text: "Wearable data not found", IsDone: false}, nil
	}

	return &pb.Empty{IsDone: true, Text: "Wearable data successfully deleted"}, nil
}


func (s *AnalyticsStorage) GenerateHealthRecommendations(req *pb.GenerateHealthRecommendationsReq) (*pb.GenerateHealthRecommendationsRes, error) {

	collection := s.db.Collection("health_recommendations")

	rsp, err := ai.WorkWithAI(req)
	if err != nil {
		log.Fatalln("Error while working with ai", err)
	}

	res := &pb.GenerateHealthRecommendationsRes{
		Id:                 primitive.NewObjectID().Hex(),
		UserId:             rsp.UserId,
		RecommendationType: rsp.RecommendationType,
		Description:        rsp.Description,
		CreatedAt:          time.Now().String(),
		UpdatedAt:          time.Now().String(),
	}

	_, err = collection.InsertOne(context.Background(), res)
	if err != nil {
		log.Println("Error inserting into database:", err)
		return nil, err
	}

	return res, nil
}

func (s *AnalyticsStorage) GetHealthRecommendationsById(req *pb.GetHealthRecommendationsByIdReq) (*pb.GetHealthRecommendationsByIdRes, error) {
	collection := s.db.Collection("health_recommendations")

	filter := bson.M{"id": req.RecommendationId}

	var record bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("health recommendation not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find health recommendation: %v", err)
	}

	log.Printf("Retrieved record: %+v", record)

	id, ok := record["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("id field is missing or invalid")
	}
	userId, ok := record["user_id"].(string)
	if !ok {
		userId = ""
	}
	recommendationType, ok := record["recommendationtype"].(string)
	if !ok || recommendationType == "" {
		return nil, fmt.Errorf("recommendation_type field is missing or invalid")
	}
	description, ok := record["description"].(string)
	if !ok || description == "" {
		return nil, fmt.Errorf("description field is missing or invalid")
	}
	createdAt, ok := record["createdat"].(string)
	if !ok || createdAt == "" {
		return nil, fmt.Errorf("created_at field is missing or invalid")
	}
	updatedAt, ok := record["updatedat"].(string)
	if !ok || updatedAt == "" {
		return nil, fmt.Errorf("updated_at field is missing or invalid")
	}

	res := &pb.GetHealthRecommendationsByIdRes{
		Info: &pb.GenerateHealthRecommendationsRes{
			Id:                 id,
			UserId:             userId,
			RecommendationType: recommendationType,
			Description:        description,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
		},
	}

	return res, nil
}

// Health Monitoring
func (s *AnalyticsStorage) GetRealtimeHealthMonitoring(req *pb.GetRealtimeHealthMonitoringReq) (*pb.GetRealtimeHealthMonitoringRes, error) {

	collection := s.db.Collection("health_monitoring")

	filter := bson.M{"user_id": req.UserId}

	var result struct {
		UserId         string            `bson:"user_id"`
		MonitoringData map[string]string `bson:"monitoring_data"`
	}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	res := &pb.GetRealtimeHealthMonitoringRes{
		UserId:         result.UserId,
		MonitoringData: result.MonitoringData,
	}

	return res, nil
}

func (s *AnalyticsStorage) GetDailyHealthSummary(req *pb.GetDailyHealthSummaryReq) (*pb.GetDailyHealthSummaryRes, error) {

	collection := s.db.Collection("daily_health_summaries")

	filter := bson.M{"user_id": req.UserId}

	var result struct {
		UserId  string `bson:"user_id"`
		Summary string `bson:"summary"`
	}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	res := &pb.GetDailyHealthSummaryRes{
		UserId:  result.UserId,
		Summary: result.Summary,
	}

	return res, nil
}

func (s *AnalyticsStorage) GetWeeklyHealthSummary(req *pb.GetWeeklyHealthSummaryReq) (*pb.GetWeeklyHealthSummaryRes, error) {

	collection := s.db.Collection("weekly_health_summaries")

	filter := bson.M{"user_id": req.UserId}

	var result struct {
		UserId  string `bson:"user_id"`
		Summary string `bson:"summary"`
	}

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	res := &pb.GetWeeklyHealthSummaryRes{
		UserId:  result.UserId,
		Summary: result.Summary,
	}

	return res, nil
}
