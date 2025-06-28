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
func CreateSymptom(ctx context.Context, symptom *Symptom) error {
	var s *Symptom
	err := DB.WithContext(ctx).
		Table(constants.TableSymptom).
		Where("symptom_id = ? ", symptom.SymptomId).
		First(&s).
		Error
	if err == nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Symptom Exist")
	}
	err = DB.WithContext(ctx).
		Table(constants.TableSymptom).
		Create(symptom).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "Symptom.Create error"+err.Error())
	}
	return nil
}

func QueryMedicalCase(ctx context.Context, patientId int64) (*Case, error) {
	var c *Case
	err := DB.WithContext(ctx).
		Table(constants.ViewMedicalCase).
		Where("patient_id=?", patientId).
		First(&c).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "the medical case not exist")
		}
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryMedicalCase: "+err.Error())
	}
	return c, nil
}
