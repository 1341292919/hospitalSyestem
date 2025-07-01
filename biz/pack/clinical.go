package pack

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/model"
	"time"
)

func Case(data *db.Case) *model.MedicalCase {
	return &model.MedicalCase{
		PatientID:          data.PatientID,
		PatientName:        data.PatientName,
		PatientGender:      data.PatientGender,
		Age:                int32(data.Age),
		ContactPhone:       data.ContactPhone,
		DiagnosisID:        data.DiagnosisID,
		DoctorID:           data.DoctorID,
		DiseaseName:        data.DiseaseName,
		DiagnosisTime:      data.DiagnosisTime.Format(time.RFC3339),
		DiagnosisNotes:     data.DiagnosisNotes,
		DiagnosisCreatedAt: data.CreatedAt.Format(time.RFC3339),
		SymptomID:          data.SymptomID,
		SymptomDescription: data.Description, //diagnosis_description
		SymptomStartTime:   data.StartTime.Format(time.RFC3339),
		SignsDescription:   data.Signs,
		SymptomCreatedAt:   data.CreatedAt.Format(time.RFC3339),
	}
}
func CaseList(data []*db.Case, count int64) *model.MedicalCaseList {
	k := make([]*model.MedicalCase, 0)
	for _, v := range data {
		k = append(k, Case(v))
	}
	return &model.MedicalCaseList{
		Items: k,
		Total: count,
	}
}

// Helper function to convert time pointer to string
func timePtrToStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
