package db

import (
	"Hospital/pkg/constants"
	"Hospital/pkg/errno"
	"context"
)

func CreatSample(ctx context.Context, sample *BioSample) error {
	err := DB.WithContext(ctx).
		Table(constants.TableBioBank).
		Create(sample).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "add sample failed:"+err.Error())
	}
	return nil
}
func QuerySample(ctx context.Context, patientId int64) ([]*BioSample, int64, error) {
	var s []*BioSample
	var count int64
	err := DB.WithContext(ctx).
		Table(constants.TableBioBank).
		Where("patient_id = ?", patientId).
		Count(&count).
		Find(&s).
		Error
	if err != nil {
		return nil, 0, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QuerySample"+err.Error())
	}
	if count == 0 {
		return nil, 0, errno.NewErrNo(errno.InternalServiceErrorCode, "Sample not found")
	}
	return s, count, nil
}
func QuerySampleById(ctx context.Context, sampleId int64) ([]*BioSample, int64, error) {
	var s []*BioSample
	var count int64
	err := DB.WithContext(ctx).
		Table(constants.TableBioBank).
		Where("sample_id = ?", sampleId).
		Count(&count).
		Find(&s).
		Error
	if err != nil {
		return nil, 0, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QuerySample"+err.Error())
	}
	if count == 0 {
		return nil, 0, errno.NewErrNo(errno.InternalServiceErrorCode, "Sample not found")
	}
	return s, count, nil
}
