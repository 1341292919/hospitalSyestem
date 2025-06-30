package service

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/epidemic"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type EpidemicService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewEpidemicService(ctx context.Context, c *app.RequestContext) *EpidemicService {
	return &EpidemicService{ctx: ctx, c: c}
}

func (s *EpidemicService) CreatEpidemicCase(req *epidemic.AddEpidemicCaseRequest) error {
	return db.CreatEpidemicCase(s.ctx, &db.EpidemicCase{
		CaseID:            req.CaseID,
		PatientID:         req.PatientID,
		OnsetDate:         time.Unix(req.OnsetDate, 0),
		DiagnosisDate:     time.Unix(req.DiagnosisDate, 0),
		CaseType:          req.CaseType,
		InfectionSource:   req.InfectionSource,
		TransmissionRoute: req.TransmissionRoute,
		Symptoms:          req.Symptoms,
		TravelHistory:     req.TravelHistory,
		RiskLocations:     req.RiskLocations,
		CloseContacts:     req.CloseContacts,
		UpdateTime:        time.Now(),
	})
}

func (s *EpidemicService) QueryEpidemicCaseByPatient(req *epidemic.QueryEpidemicCaseByPatientRequest) ([]*db.EpidemicCase, error) {
	return db.QueryEpidemicCase(s.ctx, req.PatientID)
}

func (s *EpidemicService) QueryEpidemicCaseById(req *epidemic.QueryEpidemicCaseByIdRequest) (*db.EpidemicCase, error) {
	return db.QueryEpidemicCaseById(s.ctx, req.CaseID)
}
