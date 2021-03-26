CREATE USER gorm@'%';

SELECT Host, User, plugin, authentication_string
FROM mysql.user;

UPDATE mysql.user
SET authentication_string=password('gorm'), Host='%'
WHERE User = 'gorm';



CREATE DATABASE gormdb
	CHAR SET 'utf8mb4'
	CHARACTER SET 'utf8mb4';

USE gormdb;

GRANT ALL PRIVILEGES ON gormdb.* TO gorm@'%';

FLUSH PRIVILEGES;