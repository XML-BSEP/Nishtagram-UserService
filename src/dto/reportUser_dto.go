package dto

import (
	"time"
	"user-service/domain/enum"
)

type ReportUserDTO struct {
	ID string `bson:"_id" json:"id"`
	Reported string `bson:"reported" json:"reported"`
	ReportType enum.UserReportType `bson:"report_type" json:"report_type"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	ReportedStatus enum.ReportStatus `bson:"report_status" json:"report_status"`
}
