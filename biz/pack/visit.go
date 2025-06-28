package pack

import (
	"Hospital/biz/dal/db"
	"Hospital/biz/model/model"
	"fmt"
	"strconv"
)

func VisitMessage(data *db.VisitMessage) *model.VisitMessage {
	if data == nil {
		return nil
	}

	// 时间格式常量
	const timeFormat = "2006-01-02 15:04:05"

	return &model.VisitMessage{
		FollowUpID:          toString(data.FollowUpID), // int转string
		PatientID:           data.PatientID,
		FollowUpType:        data.FollowUpType,
		FollowUpTime:        data.FollowUpTime.Format(timeFormat),
		ResponsiblePerson:   data.ResponsiblePerson,
		SymptomChanges:      data.SymptomChanges,
		VitalSigns:          data.VitalSigns,
		MedicationAdherence: data.MedicationAdherence,
		AdverseEvents:       data.AdverseEvents,
		CreatedAt:           data.CreatedAt.Format(timeFormat),
		UpdatedAt:           data.UpdatedAt.Format(timeFormat),
	}
}

func VisitMessageList(data []*db.VisitMessage, count int64) *model.VisitMessageList {
	v := make([]*model.VisitMessage, 0, len(data))
	for _, d := range data {
		v = append(v, VisitMessage(d))
	}
	return &model.VisitMessageList{
		Items: v,
		Total: count,
	}
}

// 辅助函数：任意类型转string
func toString(v interface{}) string {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
