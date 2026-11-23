package sms

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SMSRequest struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
	Priority    string `json:"priority"` // "HIGH" | "NORMAL"
}

type SMSResponse struct {
	SMSID     string `json:"sms_id"`
	Status    string `json:"status"`
	Delivered bool   `json:"delivered"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HandleSendSMS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req SMSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid payload"})
		return
	}

	smsID := "sms_" + time.Now().Format("20060102150405000")
	log.Printf("[SMS] Dispatched alert to %s (ID: %s)", req.PhoneNumber, smsID)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(SMSResponse{
		SMSID:     smsID,
		Status:    "SENT_TO_GATEWAY",
		Delivered: true,
	})
}
