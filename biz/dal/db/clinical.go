package db

import (
	"Hospital/pkg/constants"
	"Hospital/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
)

func CreatePatient(ctx context.Context, patient *Patient) error {
	var p *Patient
	err := DB.WithContext(ctx).
		Table(constants.TablePatient).
		Where("patient_id = ?", patient.PatientId).
		First(&p).
		Error
	if err == nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Patient Exist")
	}

	err = DB.WithContext(ctx).
		Table(constants.TablePatient).
		Create(patient).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "Patient.Create error"+err.Error())
	}
	return nil
}
func CreateDiagnose(ctx context.Context, diagnose *Diagnose) error {
	var d *Diagnose
	err := DB.WithContext(ctx).
		Table(constants.TableDiagnose).
		Where("diagnosis_id = ?", diagnose.DiagnosisId).
		First(&d).
		Error
	if err == nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Diagnose Exist error")
	}
	err = DB.WithContext(ctx).
		Table(constants.TableDiagnose).
		Create(diagnose).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "Diagnose.Create error"+err.Error())
	}
	return nil
}

func QueryMedicalCase(ctx context.Context, patientId int64) ([]*Case, int64, error) {
	var c []*Case
	var count int64
	err := DB.WithContext(ctx).
		Table(constants.ViewMedicalCase).
		Where("patient_id=?", patientId).
		Count(&count).
		Find(&c).
		Error
	if err != nil {
		return nil, -1, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryMedicalCase: "+err.Error())
	}
	return c, count, nil
}
func QueryAllMedicalCase(ctx context.Context, patientId int64) ([]*Case, int64, error) {
	var c []*Case
	var count int64
	err := DB.WithContext(ctx).
		Table(constants.ViewMedicalCase).
		Count(&count).
		Find(&c).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, -1, errno.NewErrNo(errno.InternalServiceErrorCode, "the medical case not exist")
		}
		return nil, -1, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryMedicalCase: "+err.Error())
	}
	return c, count, nil
}
