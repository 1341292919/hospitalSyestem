namespace go model

struct User{
    1: required i64 id,           //用户id
    2: required string name,    //用户名
    3 :required string contact_phone,
    4: required string created_at,
    5: required string position
    6: required string specialty
    7: required string department
    8: required string title
    9: required string identity
}
struct UserList{
        1: required list <User> items,
        2: required i64 total,                      // 总数量
}

struct BaseResp {
    1: required i64 code,          //请求返回的状态码
    2: required string msg,        //返回的消息
}
struct MedicalCase {
    1: required i64 patient_id,          // 患者ID
    2: required string patient_name,     // 患者姓名
    3: required string patient_gender,   // 患者性别(男/女/其他)
    4: required i32 age,                 // 患者年龄
    5: required  string contact_phone,    // 联系方式

    6: required i64 diagnosis_id,        // 诊断ID
    7: required string doctor_id,        // 主治医生ID
    8: required string disease_name,     // 疾病名称
    9: required string diagnosis_time,   // 诊断时间
    10: required  string diagnosis_notes, // 诊断备注
    11: required string diagnosis_created_at, // 诊断创建时间

    12: required  i64 symptom_id,         // 症状ID(可能为空)
    13: required  string symptom_description, // 症状描述
    14: required  string symptom_start_time,  // 症状开始时间
    15: required  string signs_description,   // 体征描述
    16: required  string symptom_created_at   // 症状记录时间
}

struct VisitMessage {
    1: required string follow_up_id,           // 随访ID
    2: required string patient_id,             // 患者ID
    3: required string follow_up_type,         // 随访类型：术后/用药/慢性病管理等
    4: required string follow_up_time,         // 随访时间
    5: required string responsible_person,     // 随访负责人(医生/护士ID)
    
    6: required string symptom_changes,        // 症状变化描述
    7: required string vital_signs,            // 体征记录
    
    8: required i8 medication_adherence,       // 用药依从性：0-不依从，1-部分依从，2-完全依从
    9: required string adverse_events,         // 不良事件描述
    
    10: required string created_at,            // 记录创建时间
    11: required string updated_at             // 记录更新时间
}

struct VisitMessageList {
    1: required list<VisitMessage> items,      // 随访记录列表
    2: required i64 total,                     // 总记录数
}


struct BioSample {
    1: required string sample_id,               // 样本ID（主键）
    2: required string patient_id,              // 患者ID
    3: required string sample_type,             // 样本类型
    4: required string collection_time,         // 采集时间
    5: required string collection_site,         // 采集部位
    6: required string collector_id,            // 采集人ID
    7: required string processing_method,       // 处理方法
    8: required string storage_condition,       // 存储条件
    9: required string storage_location,        // 存储位置
    10: required string notes,                  // 备注信息
    11: required string created_at,             // 创建时间
    12: required string updated_at              // 更新时间
}

struct SampleList {
    1: required list <BioSample> items,    // 样本列表
    2: required i64 total,                      // 总数量
}


struct EpidemicCase {
    1: required i64 case_id,                 // 病例ID，唯一标识
    2: required i64 patient_id,              // 患者ID
    3: required i64 onset_date,              // 发病日期（格式：YYYY-MM-DD）
    4: required i64 diagnosis_date,          // 诊断日期（格式：YYYY-MM-DD）
    5: required string case_type,               // 病例类型：confirmed/suspected/asymptomatic
    6: required string infection_source,        // 感染来源：local/imported/unknown
    7: required string transmission_route,      // 传播途径：逗号分隔的字符串如"droplet,contact"
    8: required string symptoms,                // 症状信息
    9: required string travel_history,          // 旅行史
    10: required string risk_locations,         // 高危地点
    11: required i16 close_contacts,            // 密接人数（smallint 对应 i16）
    12: required string update_time
}

struct EpidemicCaseList{
    1: required list <EpidemicCase> items,    // 病例列表
    2: required i64 total,                      // 总数量
}