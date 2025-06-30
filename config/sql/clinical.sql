-- 1. 患者基本信息表
CREATE TABLE IF NOT EXISTS  patients (
                                         patient_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '患者ID，主键，自动递增',
                                         name VARCHAR(100) NOT NULL COMMENT '患者姓名',
                                         gender ENUM('男', '女', '其他') NOT NULL COMMENT '性别',
                                         age INT NOT NULL COMMENT '年龄',
                                         contact_phone  VARCHAR(20) COMMENT '联系方式',
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

                                         INDEX idx_patient_name (name) COMMENT '姓名索引',
                                         INDEX idx_patient_gender (gender) COMMENT '性别索引',
                                         INDEX idx_patient_age (age) COMMENT '年龄索引'
) COMMENT '患者基本信息表';

-- 2. 诊断记录表
CREATE TABLE diagnoses (
                           diagnosis_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '诊断ID',
                           patient_id INT NOT NULL COMMENT '患者ID',
                           doctor_id VARCHAR(20) NOT NULL COMMENT '主治医生ID（无外键约束）',
                           disease_name VARCHAR(100) NOT NULL COMMENT '疾病名称',
                           description TEXT NOT NULL COMMENT '症状描述',
                           start_time DATETIME COMMENT '开始时间',
                           signs_description TEXT COMMENT '体征描述',
                           diagnosis_time DATETIME NOT NULL COMMENT '诊断时间',
                           notes TEXT COMMENT '备注信息',
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',

                           FOREIGN KEY (patient_id) REFERENCES patients(patient_id)
                               ON DELETE CASCADE
                               ON UPDATE CASCADE,

                           INDEX idx_diagnosis_patient (patient_id) COMMENT '患者ID索引',
                           INDEX idx_diagnosis_doctor (doctor_id) COMMENT '医生ID索引',
                           INDEX idx_diagnosis_disease (disease_name) COMMENT '疾病名称索引',
                           INDEX idx_diagnosis_time (diagnosis_time) COMMENT '诊断时间索引'
) COMMENT '诊断记录表';


-- 3. 创建视图
CREATE OR REPLACE VIEW patient_medical_records
AS
SELECT
    p.patient_id,
    p.name AS patient_name,
    p.gender AS patient_gender,
    p.age,
    p.contact_phone,
    d.diagnosis_id,
    d.doctor_id,
    d.disease_name,
    d.diagnosis_time,
    d.notes AS diagnosis_notes,
    d.created_at AS diagnosis_created_at,
    s.symptom_id,
    s.description AS symptom_description,
    s.start_time AS symptom_start_time,
    s.signs_description,
    s.created_at AS symptom_created_at
FROM
    diagnoses d
        JOIN patients p ON d.patient_id = p.patient_id
        LEFT JOIN symptoms s ON d.diagnosis_id = s.diagnosis_id;