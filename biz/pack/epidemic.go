package pack

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/model"
)

func EpidemicCase(data *db.EpidemicCase) *model.EpidemicCase {
	// 日期转换为时间戳（秒级）
	onsetTimestamp := data.OnsetDate.Unix()
	diagnosisTimestamp := data.DiagnosisDate.Unix()

	return &model.EpidemicCase{
		CaseID:            data.CaseID, // int64 直接赋值
		PatientID:         data.PatientID,
		OnsetDate:         onsetTimestamp,
		DiagnosisDate:     diagnosisTimestamp,
		CaseType:          data.CaseType,
		InfectionSource:   data.InfectionSource,
		TransmissionRoute: data.TransmissionRoute,
		Symptoms:          data.Symptoms,
		TravelHistory:     data.TravelHistory,
		RiskLocations:     data.RiskLocations,
		CloseContacts:     int16(data.CloseContacts), // int64 转为 int16
		UpdateTime:        data.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

func EpidemicCaseList(data []*db.EpidemicCase) *model.EpidemicCaseList {
	epidemicCases := make([]*model.EpidemicCase, 0)
	for _, v := range data {
		epidemicCases = append(epidemicCases, EpidemicCase(v))
	}
	return &model.EpidemicCaseList{
		Items: epidemicCases,
		Total: int64(len(epidemicCases)), // 计算总数
	}
}
