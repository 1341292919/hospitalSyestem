CREATE TABLE epidemic_case (
                               case_id VARCHAR(32) PRIMARY KEY COMMENT '病例ID，唯一标识',
                               patient_id VARCHAR(32) COMMENT '患者ID，可匿名',
                               onset_date DATE COMMENT '发病日期',
                               diagnosis_date DATE COMMENT '诊断日期',
                               case_type ENUM('confirmed', 'suspected', 'asymptomatic') COMMENT '病例类型：确诊/疑似/无症状',
                               infection_source ENUM('local', 'imported', 'unknown') COMMENT '感染来源：本地/输入/不明',
                               transmission_route SET('droplet', 'contact', 'airborne', 'other') COMMENT '传播途径：飞沫/接触/空气等',
                               symptoms VARCHAR(1000) COMMENT '症状信息',
                               travel_history VARCHAR(1000) COMMENT '旅行史',
                               risk_locations VARCHAR(1000) COMMENT '高危地点',
                               close_contacts SMALLINT UNSIGNED DEFAULT 0 COMMENT '密接人数',
                               update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) COMMENT '患者流行病学关键信息暴露接触信息表';