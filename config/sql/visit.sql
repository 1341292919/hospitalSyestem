CREATE TABLE follow_up_record (
    -- 随访ID作为主键
                                  follow_up_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '随访ID，主键',
    -- 患者信息
                                  patient_id VARCHAR(20) NOT NULL COMMENT '患者ID，关联临床数据库',
    -- 随访基本信息
                                  follow_up_type VARCHAR(50) NOT NULL COMMENT '随访类型：术后/用药/慢性病管理等',
                                  follow_up_time DATETIME NOT NULL COMMENT '随访时间',
                                  responsible_person VARCHAR(20) NOT NULL COMMENT '随访负责人(医生/护士ID)',
    -- 症状与体征
                                  symptom_changes TEXT COMMENT '症状变化描述',
                                  vital_signs VARCHAR(255) COMMENT '体征记录',
    -- 用药情况
                                  medication_adherence TINYINT COMMENT '用药依从性：0-不依从，1-部分依从，2-完全依从',
    -- 不良事件
                                  adverse_events TEXT COMMENT '不良事件描述',
    -- 系统记录
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    -- 创建索引
                                  INDEX idx_patient_id (patient_id) COMMENT '患者ID索引，便于按患者查询',
                                  INDEX idx_follow_up_time (follow_up_time) COMMENT '随访时间索引',
                                  INDEX idx_responsible_person (responsible_person) COMMENT '负责人索引',
                                  INDEX idx_type_time (follow_up_type, follow_up_time) COMMENT '联合索引：类型+时间',
                                  INDEX idx_medication_adherence (medication_adherence) COMMENT '用药依从性索引'
) ENGINE=InnoDB AUTO_INCREMENT=50000 DEFAULT CHARSET=utf8mb4 COMMENT='综合随访记录表';