package service

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/clinical"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type ClinicalService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewClinicalService(ctx context.Context, c *app.RequestContext) *ClinicalService {
	return &ClinicalService{ctx: ctx, c: c}
}

func (s *ClinicalService) CreatePatient(req *clinical.AddPatientRequest) error {
	return db.CreatePatient(s.ctx, &db.Patient{
		PatientId:    req.PatientID,
		Name:         req.Name,
		Gender:       req.Gender,
		Age:          req.Age,
		ContactPhone: req.ContactPhone,
	})
}
func (s *ClinicalService) CreateDiagnose(req *clinical.AddDiagnoseRequest) error {
	return db.CreateDiagnose(s.ctx, &db.Diagnose{
		DiagnosisId:   req.DiagnosisID,
		PatientId:     req.PatientID,
		DoctorId:      req.DoctorID,
		DiseaseName:   req.DiseaseName,
		DiagnosisTime: time.Unix(req.DiagnosisTime, 0),
		CreatedAT:     time.Now(),
		Notes:         req.Notes,
	})
}
func (s *ClinicalService) CreateSymptom(req *clinical.AddSymptomRequest) error {
	return db.CreateSymptom(s.ctx, &db.Symptom{
		DiagnosisId:      req.DiagnosisID,
		Description:      req.Description,
		StartTime:        time.Unix(req.StartTime, 0),
		SignsDescription: req.SignsDescription,
		SymptomId:        req.SymptomID,
	})
}
func (s *ClinicalService) QueryMedicalCase(req *clinical.QueryCaseRequest) (*db.Case, error) {
	return db.QueryMedicalCase(s.ctx, req.PatientID)
}
