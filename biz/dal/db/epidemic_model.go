package db

import (
	"time"
)

type EpidemicCase struct {
	CaseID            int64     `gorm:"primaryKey;column:case_id;type:varchar(32);comment:病例ID，唯一标识"`
	PatientID         int64     `gorm:"column:patient_id;type:varchar(32);comment:患者ID，可匿名"`
	OnsetDate         time.Time `gorm:"column:onset_date;type:date;comment:发病日期"`
	DiagnosisDate     time.Time `gorm:"column:diagnosis_date;type:date;comment:诊断日期"`
	CaseType          string    `gorm:"index;column:case_type;type:enum('confirmed','suspected','asymptomatic');comment:病例类型：确诊/疑似/无症状"`
	InfectionSource   string    `gorm:"column:infection_source;type:enum('local','imported','unknown');comment:感染来源：本地/输入/不明"`
	TransmissionRoute string    `gorm:"column:transmission_route;type:SET('droplet','contact','airborne','other');comment:传播途径：飞沫/接触/空气等"`
	Symptoms          string    `gorm:"column:symptoms;type:varchar(1000);comment:症状信息"`
	TravelHistory     string    `gorm:"column:travel_history;type:varchar(1000);comment:旅行史"`
	RiskLocations     string    `gorm:"column:risk_locations;type:varchar(1000);comment:高危地点"`
	CloseContacts     int64     `gorm:"column:close_contacts;type:smallint unsigned;default:0;comment:密接人数"`
	UpdateTime        time.Time `gorm:"autoUpdateTime;column:update_time;comment:最后更新时间"`
}
