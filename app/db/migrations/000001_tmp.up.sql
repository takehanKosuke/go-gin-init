CREATE TABLE users (
  id bigint(20) AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) UNIQUE,
  content text,
  is_accept tinyint(1) DEFAULT 0 NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
