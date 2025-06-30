package db

import (
	"Hospital/pkg/constants"
	"Hospital/pkg/errno"
	"context"
)

func CreateVisitMessage(ctx context.Context, v *VisitMessage) error {
	err := DB.WithContext(ctx).
		Table(constants.TableVisit).
		Create(v).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "CreateVisitMessage: create visit message failed")
	}
	return nil
}

func QueryVisitMessage(ctx context.Context, patientId int64) ([]*VisitMessage, int64, error) {
	var v []*VisitMessage
	var count int64
	var err error

	if patientId == 0 {
		err = DB.WithContext(ctx).
			Table(constants.TableVisit).
			Count(&count).
			Find(&v).
			Error
	} else {
		err = DB.WithContext(ctx).
			Table(constants.TableVisit).
			Where("patient_id = ?", patientId).
			Count(&count).
			Find(&v).
			Error
	}
	if err != nil {
		return nil, -1, errno.NewErrNo(errno.InternalDatabaseErrorCode, "Failed to find visitMessage")
	}
	if count == 0 {
		return nil, 0, errno.NewErrNo(errno.InternalServiceErrorCode, "patient no message")
	}
	return v, count, nil
}
