CREATE TABLE scan_results (
  scan_result_id INT PRIMARY KEY NOT NULL,
  scan_ip    VARCHAR(39) NOT NULL,
  host       VARCHAR(40) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  query_data JSON NOT NULL
);
