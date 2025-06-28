namespace go visit

include "model.thrift"

struct AddVisitMessageRequest{
        2: required string patient_id,             // 患者ID
        3: required string follow_up_type,         // 随访类型：术后/用药/慢性病管理等
        4: required i64 follow_up_time,         // 随访时间
        5: required string responsible_person,     // 随访负责人(医生/护士ID)
        6: required string symptom_changes,        // 症状变化描述
        7: required string vital_signs,            // 体征记录
        8: required i8 medication_adherence,       // 用药依从性：0-不依从，1-部分依从，2-完全依从
        9: required string adverse_events,         // 不良事件描述
}
struct AddVisitMessageResponse{
    1: required model.BaseResp base,
}
struct QueryVisitMessageRequest{
    1: required i64 patient_id,
}
struct QueryVisitMessageResponse{
        1:  required model.BaseResp base,
        2: required model.VisitMessageList data,
}

service VisitService {
    AddVisitMessageResponse AddVisitMessage(1:AddVisitMessageRequest req)(api.post = "/visit/add"),
    QueryVisitMessageResponse QueryVisitMessage(1: QueryVisitMessageRequest req)(api.get ="visit/info"),
}
