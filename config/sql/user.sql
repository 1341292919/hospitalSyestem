CREATE TABLE `doctor` (
                          doctor_id VARCHAR(20) PRIMARY KEY COMMENT '医生唯一标识ID',
                          name VARCHAR(50) NOT NULL COMMENT '医生姓名',
                          specialty VARCHAR(100) COMMENT '专业领域',
                          title VARCHAR(50) COMMENT '职称',
                          department VARCHAR(50) NOT NULL COMMENT '所属科室',
                          password VARCHAR(255) NOT NULL COMMENT '登录密码',
                          contact_phone VARCHAR(20) COMMENT '联系电话',
                          create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)
    ENGINE=InnoDB
    DEFAULT CHARSET=utf8mb4
    COMMENT='医生基本信息表';

CREATE TABLE `nurse` (
                         nurse_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '护士唯一标识ID',
                         name VARCHAR(50) NOT NULL COMMENT '护士姓名',
                         contact_phone VARCHAR(20) COMMENT '联系电话',
                         department VARCHAR(50) NOT NULL COMMENT '所属科室',
                         password VARCHAR(255) NOT NULL COMMENT '登录密码',
                         position VARCHAR(50) COMMENT '职位(护士长/主管护师/护师/护士等)',
                         create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)
    ENGINE=InnoDB
    DEFAULT CHARSET=utf8mb4
    COMMENT='护士基本信息表';

-- 创建管理员表
CREATE TABLE `admin` (
                         admin_id INT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员唯一标识ID',
                         username VARCHAR(50) NOT NULL UNIQUE COMMENT '管理员登录用户名，必须唯一',
                         password VARCHAR(255) NOT NULL COMMENT '管理员登录密码',
                         contact_phone VARCHAR(20) COMMENT '联系电话',
                         create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '账号创建时间',
                         status TINYINT DEFAULT 1 COMMENT '账号状态：0-禁用，1-启用'
)
    ENGINE=InnoDB
    DEFAULT CHARSET=utf8mb4
    COMMENT='系统管理员基本信息表';

-- 插入默认管理员（密码建议使用加密后的值，这里示例使用明文"admin123"，实际应使用如bcrypt加密）
INSERT INTO `admin` (username, password, contact_phone, status)
VALUES ('admin', 'admin', '13800138000', 1);