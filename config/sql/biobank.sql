-- 3. 创建表结构
CREATE TABLE biological_samples (
                                    sample_id bigint PRIMARY KEY AUTO_INCREMENT,
                                    patient_id VARCHAR(20) NOT NULL,
                                    sample_type VARCHAR(50) NOT NULL,
                                    collection_time DATETIME NOT NULL,
                                    collection_site VARCHAR(100),
                                    collector_id VARCHAR(20) NOT NULL,
                                    processing_method VARCHAR(100),
                                    storage_condition VARCHAR(50) NOT NULL,
                                    storage_location VARCHAR(100) NOT NULL,
                                    notes TEXT,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) AUTO_INCREMENT = 90000;
