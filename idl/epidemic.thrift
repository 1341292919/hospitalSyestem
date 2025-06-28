namespace go epidemic

include "model.thrift"
struct AddEpidemicCaseRequest{
        1: required i64 case_id,                 // 病例ID，唯一标识
        2: required i64 patient_id,              // 患者ID
        3: required i64 onset_date,              // 病发日期（格式：YYYY-MM-DD）
        4: required i64 diagnosis_date,          // 诊断日期（格式：YYYY-MM-DD）
        5: required string case_type,               // 病例类型：confirmed/suspected/asymptomatic
        6: required string infection_source,        // 感染来源：local/imported/unknown
        7: required string transmission_route,      // 传播途径：逗号分隔的字符串如"droplet,contact"
        8: required string symptoms,                // 症状信息
        9: required string travel_history,          // 旅行史
        10: required string risk_locations,         // 高危地点
        11: required i64 close_contacts,            // 密接人数（smallint 对应 i16）
}

struct AddEpidemicCaseResponse{
    1: model.BaseResp base,
}
struct QueryEpidemicCaseByIdRequest{
      1: required i64 case_id,                 // 病例ID，唯一标识
}
struct QueryEpidemicCaseByIdResponse{
    1: model.BaseResp base,
    2:model.EpidemicCase data
}
struct QueryEpidemicCaseByPatientRequest{
        2: required i64 patient_id,              // 患者ID
}
struct QueryEpidemicCaseByPatientResponse{
    1: model.BaseResp base,
    2:model.EpidemicCase data
}
service EpidemicService {
    AddEpidemicCaseResponse AddEpidemicCase(1:AddEpidemicCaseRequest req)(api.post="/epidemic/add"),
    QueryEpidemicCaseByIdResponse QueryEpidemicCaseById(1:QueryEpidemicCaseByIdRequest req)(api.get="/epidemic/info"),
    QueryEpidemicCaseByPatientResponse QueryEpidemicCaseByPatient(1:QueryEpidemicCaseByPatientRequest req)(api.get="/epidemic/info/patient"),
}
