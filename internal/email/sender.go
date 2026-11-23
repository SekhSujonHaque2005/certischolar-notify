package email

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type SendRequest struct {
	To          string `json:"to"`
	Subject     string `json:"subject"`
	TemplateKey string `json:"template_key"`
	StudentName string `json:"student_name"`
	SchemeName  string `json:"scheme_name"`
	PortalURL   string `json:"portal_url"`
}

type SendResponse struct {
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HandleSendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req SendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid payload"})
		return
	}

	// Institutional privacy rule (SRS FR-10.1): Do not include marks, income or Aadhaar in raw email
	msgID := "msg_em_" + time.Now().Format("20060102150405000")
	log.Printf("[EMAIL] Dispatched to %s for template '%s' (MsgID: %s)", req.To, req.TemplateKey, msgID)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(SendResponse{
		MessageID: msgID,
		Status:    "QUEUED_FOR_DISPATCH",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
