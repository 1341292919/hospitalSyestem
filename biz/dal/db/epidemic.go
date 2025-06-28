package db

import (
	"Hospital/pkg/constants"
	"Hospital/pkg/errno"
	"context"
	"database/sql"
	"errors"
)

func CreatEpidemicCase(ctx context.Context, epidemicCase *EpidemicCase) error {
	var e *EpidemicCase
	err := DB.WithContext(ctx).
		Table(constants.TableEpidemicCase).
		Where("case_id = ?", epidemicCase.CaseID).
		First(&e).
		Error
	if err == nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "case exist")
	}
	err = DB.WithContext(ctx).
		Table(constants.TableEpidemicCase).
		Create(epidemicCase).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "add EpidemicCase failed:"+err.Error())
	}
	return nil
}
func QueryEpidemicCase(ctx context.Context, patientId int64) (*EpidemicCase, error) {
	var s *EpidemicCase
	err := DB.WithContext(ctx).
		Table(constants.TableEpidemicCase).
		Where("patient_id = ?", patientId).
		First(&s).
		Error
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "patient‘s case not exist")
		}
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryEpidemicCase"+err.Error())
	}
	return s, nil
}

func QueryEpidemicCaseById(ctx context.Context, caseId int64) (*EpidemicCase, error) {
	var s *EpidemicCase
	err := DB.WithContext(ctx).
		Table(constants.TableEpidemicCase).
		Where("case_id = ?", caseId).
		First(&s).
		Error
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "case not exist")
		}
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryEpidemicCase"+err.Error())
	}
	return s, nil
}
