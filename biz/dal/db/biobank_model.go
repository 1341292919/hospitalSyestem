package db

import "time"

type BioSample struct {
	SampleID         int64     `gorm:"primaryKey;column:sample_id;type:varchar(20);comment:样本ID（主键）"`
	PatientID        int64     `gorm:"index;column:patient_id;type:varchar(20);not null;comment:患者ID"`
	SampleType       string    `gorm:"index;column:sample_type;type:varchar(50);not null;comment:样本类型"`
	CollectionTime   time.Time `gorm:"index;column:collection_time;type:datetime;not null;comment:采集时间"`
	CollectionSite   string    `gorm:"column:collection_site;type:varchar(100);comment:采集部位"`
	CollectorID      int64     `gorm:"column:collector_id;type:varchar(20);not null;comment:采集人ID"`
	ProcessingMethod string    `gorm:"column:processing_method;type:varchar(100);comment:处理方法"`
	StorageCondition string    `gorm:"column:storage_condition;type:varchar(50);not null;comment:存储条件"`
	StorageLocation  string    `gorm:"index;column:storage_location;type:varchar(100);not null;comment:存储位置"`
	Notes            string    `gorm:"column:notes;type:text;comment:备注信息"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at;comment:创建时间"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at;comment:更新时间"`
}
