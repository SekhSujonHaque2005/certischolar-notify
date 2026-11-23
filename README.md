# CertiScholar Communications & Notification Service (Golang)

High-concurrency Email, SMS, and Webhook dispatch engine for CertiScholar.
Governed by **CS-SRS-002 v3.2 (FR-10.1, §9.3, ADR-005)**.

## Responsibilities
- Asynchronous student notification delivery (application status changes, correction alerts, deadline reminders).
- SMS gateway integration (OTP, critical alerts) with idempotency keys.
- Signed webhook event delivery to university partners & external scholarship agencies.
- Isolated from transactional core database to protect performance.

## Endpoints
- `GET /healthz` - Liveness probe
- `GET /readyz` - Readiness probe
- `POST /v1/notifications/email` - Send student email notification
- `POST /v1/notifications/sms` - Send instant SMS alert
- `POST /v1/notifications/webhook` - Deliver signed partner webhook

## Local Run
```bash
go run cmd/server/main.go
# Server listens on port 8081
```
