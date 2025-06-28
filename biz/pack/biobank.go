package pack

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/model"
	"strconv"
)

const timeFormat = "2006-01-02 15:04:05" // 统一时间格式

func BioSample(data *db.BioSample) *model.BioSample {
	if data == nil {
		return nil
	}

	return &model.BioSample{
		SampleID:         strconv.FormatInt(data.SampleID, 10), // int64转string
		PatientID:        strconv.FormatInt(data.PatientID, 10),
		SampleType:       data.SampleType,
		CollectionTime:   data.CollectionTime.Format(timeFormat),
		CollectionSite:   data.CollectionSite,
		CollectorID:      strconv.FormatInt(data.CollectorID, 10),
		ProcessingMethod: data.ProcessingMethod,
		StorageCondition: data.StorageCondition,
		StorageLocation:  data.StorageLocation,
		Notes:            data.Notes,
		CreatedAt:        data.CreatedAt.Format(timeFormat),
		UpdatedAt:        data.UpdatedAt.Format(timeFormat),
	}
}

func BioSampleList(data []*db.BioSample, count int64) *model.SampleList {
	v := make([]*model.BioSample, 0, len(data))
	for _, d := range data {
		v = append(v, BioSample(d))
	}
	return &model.SampleList{
		Items: v,
		Total: count,
	}
}
