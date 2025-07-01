package db

import "time"

type Patient struct {
	PatientId    int64
	Name         string
	Gender       string
	Age          int64
	ContactPhone string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Diagnose struct {
	DiagnosisId      int64
	PatientId        int64
	DoctorId         int64
	DiseaseName      string
	DiagnosisTime    time.Time
	Description      string
	StartTime        time.Time
	SignsDescription string
	Notes            string
	CreatedAT        time.Time
}

type Case struct {
	// 患者信息
	PatientID     int64  `gorm:"column:patient_id" json:"patient_id"`
	PatientName   string `gorm:"column:patient_name" json:"patient_name"`
	PatientGender string `gorm:"column:patient_gender" json:"patient_gender"` // 枚举值：男/女/其他
	Age           int    `gorm:"column:age" json:"age"`
	ContactPhone  string `gorm:"column:contact_phone" json:"contact_phone"`

	// 诊断信息
	DiagnosisID    int64     `gorm:"column:diagnosis_id" json:"diagnosis_id"`
	DoctorID       string    `gorm:"column:doctor_id" json:"doctor_id"`
	DiseaseName    string    `gorm:"column:disease_name" json:"disease_name"`
	Description    string    `gorm:"column:diagnosis_description" json:"description,omitempty"`
	DiagnosisTime  time.Time `gorm:"column:diagnosis_time" json:"diagnosis_time"`
	DiagnosisNotes string    `gorm:"column:diagnosis_notes" json:"diagnosis_notes"`
	CreatedAt      time.Time `gorm:"column:diagnosis_created_at" json:"created_at"`

	// 症状信息（可能为空）
	SymptomID      int64     `gorm:"column:symptom_id" json:"symptom_id,omitempty"`
	StartTime      time.Time `gorm:"column:symptom_start_time" json:"start_time,omitempty"`
	Signs          string    `gorm:"column:signs_description" json:"signs,omitempty"`
	SymptomCreated time.Time `gorm:"column:symptom_created_at" json:"symptom_created_at,omitempty"`
}
