-- Database schema (PostgreSQL)

DROP TABLE IF EXISTS family_list;
DROP TABLE IF EXISTS customer;
DROP TABLE IF EXISTS nationality;

CREATE TABLE nationality (
  nationality_id   SERIAL       PRIMARY KEY,
  nationality_name VARCHAR(50)  NOT NULL,
  nationality_code CHAR(2)      NOT NULL
);

CREATE TABLE customer (
  cst_id         SERIAL       PRIMARY KEY,
  nationality_id INT          NOT NULL,
  cst_name       CHAR(50)     NOT NULL,
  cst_dob        DATE         NOT NULL,
  "cst_phoneNum" VARCHAR(20)  NOT NULL,
  cst_email      VARCHAR(50)  NOT NULL,
  CONSTRAINT fk_customer_nationality
  FOREIGN KEY (nationality_id)
  REFERENCES nationality (nationality_id)
);

CREATE TABLE family_list (
  fl_id       SERIAL       PRIMARY KEY,
  cst_id      INT          NOT NULL,
  fl_relation VARCHAR(50)  NOT NULL,
  fl_name     VARCHAR(50)  NOT NULL,
  fl_dob      VARCHAR(50)  NOT NULL,
  CONSTRAINT fk_family_customer
  FOREIGN KEY (cst_id)
  REFERENCES customer (cst_id)
  ON DELETE CASCADE
);

CREATE INDEX idx_family_cst_id ON family_list (cst_id);
CREATE INDEX idx_customer_nationality_id ON customer (nationality_id);
