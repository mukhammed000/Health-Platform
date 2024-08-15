package handler

import (
	"context"
	"net/http"

	"api/genproto/analytics"

	"github.com/gin-gonic/gin"
)

// AddMedicalRecord godoc
// @Summary Add new medical record
// @Description Add a new medical record for a user
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.AddMedicalRecordReq true "Medical record details"
// @Success 200 {object} analytics.AddMedicalRecordRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-record [post]
func (h *Handler) AddMedicalRecord(c *gin.Context) {
	var req analytics.AddMedicalRecordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.AddMedicalRecord(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMedicalRecords godoc
// @Summary Get all medical records
// @Description Retrieve all medical records
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetMedicalRecordsRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-records [get]
func (h *Handler) GetMedicalRecords(c *gin.Context) {
	req := &analytics.GetMedicalRecordsReq{} // Customize if needed
	res, err := h.Analytics.GetMedicalRecord(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMedicalRecordById godoc
// @Summary Get medical record by ID
// @Description Retrieve a medical record by its ID
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Medical record ID"
// @Success 200 {object} analytics.GetMedicalRecordsByIdRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-record/{id} [get]
func (h *Handler) GetMedicalRecordById(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.GetMedicalRecordsByIdReq{RecordId: id}
	res, err := h.Analytics.GetMedicalRecordsById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateMedicalRecord godoc
// @Summary Update a medical record
// @Description Update an existing medical record
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.UpdateMedicalRecordReq true "Updated medical record details"
// @Success 200 {object} analytics.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-record [put]
func (h *Handler) UpdateMedicalRecord(c *gin.Context) {
	var req analytics.UpdateMedicalRecordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.UpdateMedicalRecord(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteMedicalRecord godoc
// @Summary Delete a medical record
// @Description Delete a medical record by its ID
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Medical record ID"
// @Success 200 {object} analytics.Empty
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-record/{id} [delete]
func (h *Handler) DeleteMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.DeleteMedicalRecordReq{Id: id}
	res, err := h.Analytics.DeleteMedicalRecord(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListMedicalRecords godoc
// @Summary List all medical records
// @Description List all medical records with optional filters
// @Tags Medical Records
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.ListMedicalRecordsRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/medical-records/list [get]
func (h *Handler) ListMedicalRecords(c *gin.Context) {
	req := &analytics.ListMedicalRecordsReq{} // Customize if needed
	res, err := h.Analytics.ListMedicalRecords(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AddLifestyleData godoc
// @Summary Add new lifestyle data
// @Description Add new lifestyle data for a user
// @Tags Life Style Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.AddLifestyleDataReq true "Lifestyle data details"
// @Success 200 {object} analytics.AddLifestyleDataRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/lifestyle-data [post]
func (h *Handler) AddLifestyleData(c *gin.Context) {
	var req analytics.AddLifestyleDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.AddLifestyleData(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLifestyleData godoc
// @Summary Get all lifestyle data
// @Description Retrieve all lifestyle data
// @Tags Life Style Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetLifestyleDataRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/lifestyle-data [get]
func (h *Handler) GetLifestyleData(c *gin.Context) {
	req := &analytics.GetLifestyleDataReq{} // Customize if needed
	res, err := h.Analytics.GetLifestyleData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLifestyleDataById godoc
// @Summary Get lifestyle data by ID
// @Description Retrieve lifestyle data by its ID
// @Tags Life Style Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Lifestyle data ID"
// @Success 200 {object} analytics.GetLifestyleDataByIdRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/lifestyle-data/{id} [get]
func (h *Handler) GetLifestyleDataById(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.GetLifestyleDataByIdReq{DataId: id}

	res, err := h.Analytics.GetLifestyleDataById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateLifestyleData godoc
// @Summary Update lifestyle data
// @Description Update existing lifestyle data
// @Tags Life Style Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.UpdateLifestyleDataReq true "Updated lifestyle data details"
// @Success 200 {object} analytics.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/lifestyle-data [put]
func (h *Handler) UpdateLifestyleData(c *gin.Context) {
	var req analytics.UpdateLifestyleDataReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.UpdateLifestyleData(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteLifestyleData godoc
// @Summary Delete lifestyle data
// @Description Delete lifestyle data by its ID
// @Tags Life Style Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Lifestyle data ID"
// @Success 200 {object} analytics.Empty
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/lifestyle-data/{id} [delete]
func (h *Handler) DeleteLifestyleData(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.DeleteLifestyleDataReq{Id: id}
	res, err := h.Analytics.DeleteLifestyleData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AddWearableData godoc
// @Summary Add new wearable data
// @Description Add new wearable data for a user
// @Tags Wearable Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.AddWearableDataReq true "Wearable data details"
// @Success 200 {object} analytics.AddWearableDataRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/wearable-data [post]
func (h *Handler) AddWearableData(c *gin.Context) {
	var req analytics.AddWearableDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.AddWearableData(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetWearableData godoc
// @Summary Get all wearable data
// @Description Retrieve all wearable data
// @Tags Wearable Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetWearableDataRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/wearable-data [get]
func (h *Handler) GetWearableData(c *gin.Context) {
	req := &analytics.GetWearableDataReq{} // Customize if needed
	res, err := h.Analytics.GetWearableData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetWearableDataById godoc
// @Summary Get wearable data by ID
// @Description Retrieve wearable data by its ID
// @Tags Wearable Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Wearable data ID"
// @Success 200 {object} analytics.GetWearableDataByIdRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/wearable-data/{id} [get]
func (h *Handler) GetWearableDataById(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.GetWearableDataByIdReq{DataId: id}
	res, err := h.Analytics.GetWearableDataById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateWearableData godoc
// @Summary Update wearable data
// @Description Update existing wearable data
// @Tags Wearable Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.UpdateWearableDataReq true "Updated wearable data details"
// @Success 200 {object} analytics.Empty
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/wearable-data [put]
func (h *Handler) UpdateWearableData(c *gin.Context) {
	var req analytics.UpdateWearableDataReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.UpdateWearableData(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteWearableData godoc
// @Summary Delete wearable data
// @Description Delete wearable data by its ID
// @Tags Wearable Data
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Wearable data ID"
// @Success 200 {object} analytics.Empty
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/wearable-data/{id} [delete]
func (h *Handler) DeleteWearableData(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.DeleteWearableDataReq{Id: id}
	res, err := h.Analytics.DeleteWearableData(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GenerateHealthRecommendations godoc
// @Summary Generate health recommendations
// @Description Generate health recommendations based on user data
// @Tags Health Recommendations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body analytics.GenerateHealthRecommendationsReq true "Health data details"
// @Success 200 {object} analytics.GenerateHealthRecommendationsRes
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/health-recommendations [post]
func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
	var req analytics.GenerateHealthRecommendationsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Analytics.GenerateHealthRecommendations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetHealthRecommendationsById godoc
// @Summary Get health recommendations by ID
// @Description Retrieve health recommendations by its ID
// @Tags Health Recommendations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Recommendation ID"
// @Success 200 {object} analytics.GetHealthRecommendationsByIdRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/health-recommendations/{id} [get]
func (h *Handler) GetHealthRecommendationsById(c *gin.Context) {
	id := c.Param("id")
	req := &analytics.GetHealthRecommendationsByIdReq{RecommendationId: id}

	res, err := h.Analytics.GetHealthRecommendationsById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetRealtimeHealthMonitoring godoc
// @Summary Get real-time health monitoring data
// @Description Retrieve real-time health monitoring data
// @Tags Health Monitoring
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetRealtimeHealthMonitoringRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/health-monitoring/realtime [get]
func (h *Handler) GetRealtimeHealthMonitoring(c *gin.Context) {
	req := &analytics.GetRealtimeHealthMonitoringReq{} // Customize if needed
	res, err := h.Analytics.GetRealtimeHealthMonitoring(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetDailyHealthSummary godoc
// @Summary Get daily health summary
// @Description Retrieve daily health summary data
// @Tags Health Monitoring
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetDailyHealthSummaryRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/health-monitoring/daily-summary [get]
func (h *Handler) GetDailyHealthSummary(c *gin.Context) {
	req := &analytics.GetDailyHealthSummaryReq{} // Customize if needed
	res, err := h.Analytics.GetDailyHealthSummary(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetWeeklyHealthSummary godoc
// @Summary Get weekly health summary
// @Description Retrieve weekly health summary data
// @Tags Health Monitoring
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} analytics.GetWeeklyHealthSummaryRes
// @Failure 500 {string} string "Internal Server Error"
// @Router /analytics/health-monitoring/weekly-summary [get]
func (h *Handler) GetWeeklyHealthSummary(c *gin.Context) {
	req := &analytics.GetWeeklyHealthSummaryReq{} // Customize if needed
	res, err := h.Analytics.GetWeeklyHealthSummary(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
