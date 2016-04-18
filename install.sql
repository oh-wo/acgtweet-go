-- Drop any existing tables.
DROP DATABASE IF EXISTS acgtweet;
CREATE DATABASE acgtweet;
USE acgtweet;

CREATE TABLE users (
  id         INT NOT NULL AUTO_INCREMENT,
  first_name TEXT,
  last_name  TEXT,
  PRIMARY KEY (id)
);

INSERT INTO users (first_name,last_name) VALUES ('owen', 'bodley');