package service

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/biobank"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type BioBankService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewBioBankService(ctx context.Context, c *app.RequestContext) *BioBankService {
	return &BioBankService{ctx: ctx, c: c}
}

func (s *BioBankService) NewSample(req *biobank.AddSampleRequest) error {
	return db.CreatSample(s.ctx, &db.BioSample{
		PatientID:        req.PatientID,
		SampleType:       req.SampleType,
		CollectionTime:   time.Unix(req.CollectionTime, 0),
		CollectionSite:   req.CollectionSite,
		CollectorID:      req.CollectorID,
		ProcessingMethod: req.ProcessingMethod,
		StorageCondition: req.StorageCondition,
		StorageLocation:  req.StorageLocation,
		Notes:            req.Notes,
		CreatedAt:        time.Now(),
	})
}

func (s *BioBankService) QuerySampleByPatient(req *biobank.QuerySampleByPatientRequest) ([]*db.BioSample, int64, error) {
	return db.QuerySample(s.ctx, req.PatientID)
}

func (s *BioBankService) QuerySampleById(req *biobank.QuerySampleByIdRequest) ([]*db.BioSample, int64, error) {
	return db.QuerySampleById(s.ctx, req.SampleID)
}
