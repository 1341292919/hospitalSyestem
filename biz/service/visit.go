package service

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/visit"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type VisitService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewVisitService(ctx context.Context, c *app.RequestContext) *VisitService {
	return &VisitService{ctx: ctx, c: c}
}

func (s *VisitService) NewVisitMessage(req *visit.AddVisitMessageRequest) error {
	return db.CreateVisitMessage(s.ctx, &db.VisitMessage{
		PatientID:           req.PatientID,
		FollowUpType:        req.FollowUpType,
		FollowUpTime:        time.Unix(req.FollowUpTime, 0),
		ResponsiblePerson:   req.ResponsiblePerson,
		VitalSigns:          req.VitalSigns,
		MedicationAdherence: req.MedicationAdherence,
		SymptomChanges:      req.SymptomChanges,
		AdverseEvents:       req.AdverseEvents,
		CreatedAt:           time.Now(),
	})
}
func (s *VisitService) QueryVisitMessage(req *visit.QueryVisitMessageRequest) ([]*db.VisitMessage, int64, error) {
	return db.QueryVisitMessage(s.ctx, req.PatientID)
}
