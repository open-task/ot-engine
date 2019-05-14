CREATE TABLE IF NOT EXISTS users (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  addr VARCHAR(43) UNIQUE,
  email VARCHAR(80),
  updatetime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS skills (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `tag` VARCHAR(200) UNIQUE,
  claim INT DEFAULT 0 COMMENT 'claim number of this skill',
  submit INT DEFAULT 0 COMMENT 'submit number of this skill',
  confirm INT DEFAULT 0 COMMENT 'confirm number of this skill',
  updatetime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS statements (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  skill_id INT,
  status INT DEFAULT 0 COMMENT '0 default(claimed), 1 claimed, 2 submitted, 3 confirmed',
  submit INT DEFAULT 0 COMMENT 'submit number of this user',
  confirm INT DEFAULT 0 COMMENT 'confirm number of this user',
  filter TINYINT(1) DEFAULT 0 COMMENT '0 normal, 1 filter',
  updatetime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY(user_id, skill_id),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE NO action ON UPDATE NO action,
  FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE NO action ON UPDATE NO action
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;