package db

import (
	"time"
)

type VisitMessage struct {
	FollowUpID          int       `gorm:"primaryKey;autoIncrement;column:follow_up_id;comment:随访ID，主键"`
	PatientID           string    `gorm:"type:varchar(20);not null;column:patient_id;comment:患者ID，关联临床数据库"`
	FollowUpType        string    `gorm:"type:varchar(50);not null;column:follow_up_type;comment:随访类型：术后/用药/慢性病管理等"`
	FollowUpTime        time.Time `gorm:"column:follow_up_time;comment:随访时间"`
	ResponsiblePerson   string    `gorm:"type:varchar(20);not null;column:responsible_person;comment:随访负责人(医生/护士ID)"`
	SymptomChanges      string    `gorm:"type:text;column:symptom_changes;comment:症状变化描述"`
	VitalSigns          string    `gorm:"type:varchar(255);column:vital_signs;comment:体征记录"`
	MedicationAdherence int8      `gorm:"type:tinyint;column:medication_adherence;comment:用药依从性：0-不依从，1-部分依从，2-完全依从"`
	AdverseEvents       string    `gorm:"type:text;column:adverse_events;comment:不良事件描述"`
	CreatedAt           time.Time `gorm:"autoCreateTime;column:created_at;comment:记录创建时间"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime;column:updated_at;comment:记录更新时间"`
}
