-- DDL comum para PostgreSQL e SQLite.
-- Em SQLite, INTEGER PRIMARY KEY recebe o valor rowid automaticamente.
-- Em PostgreSQL, configurar a coluna id como IDENTITY na migration de provisionamento.
-- TIMESTAMP é usado por compatibilidade com SQLite; a aplicação grava os
-- instantes com o fuso America/Sao_Paulo.

CREATE TABLE IF NOT EXISTS sensor (
  id INTEGER PRIMARY KEY,
  unique_id VARCHAR NOT NULL,
  class VARCHAR NULL,
  icon VARCHAR NULL,
  manufacturer VARCHAR NULL,
  model VARCHAR NULL,
  name VARCHAR NOT NULL,
  picture VARCHAR NULL,
  registered_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  unit_of_measurement VARCHAR NULL,
  CONSTRAINT uq_sensor_unique_id UNIQUE (unique_id)
);

CREATE TABLE IF NOT EXISTS measurement (
  id INTEGER PRIMARY KEY,
  id_sensor INTEGER NOT NULL,
  collected_at TIMESTAMP NOT NULL,
  value DOUBLE PRECISION NOT NULL,
  class VARCHAR NOT NULL,
  unit_of_measurement VARCHAR NOT NULL,
  CONSTRAINT ck_measurement_class
    CHECK (class IN ('temperature', 'atmospheric_pressure', 'humidity')),
  CONSTRAINT ck_measurement_unit_of_measurement
    CHECK (unit_of_measurement IN ('hPa', 'Pa', 'kPa', 'bar', 'psi', 'C', 'F', 'K', '%', 'mmHg')),
  CONSTRAINT fk_measurement_sensor
    FOREIGN KEY (id_sensor)
    REFERENCES sensor (id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS publication (
  id INTEGER PRIMARY KEY,
  id_measurement INTEGER NOT NULL,
  sent_at TIMESTAMP NOT NULL,
  destination VARCHAR NOT NULL,
  CONSTRAINT uq_publication_measurement_destination UNIQUE (id_measurement, destination),
  CONSTRAINT ck_publication_destination
    CHECK (destination IN ('mqtt', 'postgres', 'influxdb')),
  CONSTRAINT fk_publication_measurement
    FOREIGN KEY (id_measurement)
    REFERENCES measurement (id)
    ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS ix_measurement_time ON measurement (collected_at, id);
CREATE INDEX IF NOT EXISTS ix_measurement_sensor ON measurement (id_sensor);
