namespace go biobank

include "model.thrift" 

struct AddSampleRequest{
    2: required i64 patient_id,              // 患者ID
    3: required string sample_type,             // 样本类型
    4: required i64 collection_time,         // 采集时间
    5: required string collection_site,         // 采集部位
    6: required i64 collector_id,            // 采集人ID
    7: required string processing_method,       // 处理方法
    8: required string storage_condition,       // 存储条件
    9: required string storage_location,        // 存储位置
    10:required string notes,                  // 备注信息
}

struct AddSampleResponse{
    1: model.BaseResp base,
}

struct QuerySampleByPatientRequest{
    1: required i64 patient_id
}

struct QuerySampleByPatientResponse{
         1: model.BaseResp base,
         2: model.SampleList data,
}

struct QuerySampleByIdRequest{
    1: required i64 sample_id
}

struct QuerySampleByIdResponse{
         1: model.BaseResp base,
         2: model.SampleList data,
}


service BioBankService {
    AddSampleResponse AddSample(1:AddSampleRequest req)(api.post = "/sample/add"),
    QuerySampleByIdResponse QuerySampleById(1:QuerySampleByIdRequest req)(api.get = "/sample/info"),
    QuerySampleByPatientResponse QuerySampleByPatient(1:QuerySampleByPatientRequest req)(api.get = "/sample/info/patient"),
}