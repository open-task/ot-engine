-- dev
CREATE DATABASE backend_dev CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'backend'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON backend_dev.* To 'backend'@'localhost';

-- test
CREATE DATABASE backend_test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'backend'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON backend_test.* To 'backend'@'localhost';

-- prod
CREATE DATABASE backend CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'backend'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON backend.* To 'backend'@'localhost';