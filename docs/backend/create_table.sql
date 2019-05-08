CREATE TABLE IF NOT EXISTS skill (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  addr VARCHAR(43),
  skill VARCHAR(200),
  status INT DEFAULT 0 COMMENT '0 default(claimed), 1 claimed, 2 submitted, 3 confirmed',
  submit_num INT DEFAULT 0,
  confirm_num INT DEFAULT 0,
  filter TINYINT(1) DEFAULT 0 COMMENT '0 normal, 1 filter',
  updatetime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY(addr, skill)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;