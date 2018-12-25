CREATE TABLE scan_results (
  scan_result_id BIGSERIAL PRIMARY KEY,
  scan_ip    VARCHAR(39) NOT NULL,
  host       VARCHAR(40) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  query_data JSON NOT NULL
);
