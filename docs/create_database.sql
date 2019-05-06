CREATE DATABASE rinkeby3 CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'engine'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON rinkeby3.* To 'engine'@'localhost';