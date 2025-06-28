namespace go clinical 

include "model.thrift" 

struct AddPatientRequest{
    1:required i64  patient_id,
    2:required string name,
    3:required string contact_phone,
    4:required i64 age,
    5:required string gender
}
struct AddPatientResponse{
     1: model.BaseResp base,
}
struct AddDiagnoseRequest{
    1:required i64  patient_id,
    2:required i64 doctor_id,
    3:required i64  diagnosis_time,
    4:required string notes,
    5:required string   disease_name
    6:required i64 diagnosis_id,
}
struct AddDiagnoseResponse{
     1: model.BaseResp base,
}
struct AddSymptomRequest{
    1:required i64 diagnosis_id,
    2:required i64 symptom_id,
    3:required string  description,
    4:required i64 start_time,
    5:required string  signs_description,
}
struct AddSymptomResponse{
     1: model.BaseResp base,
}
struct QueryCaseRequest{
    1:required i64  patient_id,
}
struct QueryCaseResponse{
     1: model.BaseResp base,
     2: model.MedicalCase data,
}

service ClinicalService {
    AddPatientResponse AddPatient(1:AddPatientRequest req)(api.post="/clinical/patient/add"),
    AddDiagnoseResponse AddDiagnose(1:AddDiagnoseRequest req)(api.post="/clinical/diagnose/add"),
    AddSymptomResponse AddSymptom(1:AddSymptomRequest req)(api.post="/clinical/symptom/add"),
    QueryCaseResponse QueryCase(1:QueryCaseRequest req)(api.get="/clinical/case/info"),
}

