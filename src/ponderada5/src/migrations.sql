CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Criação da tabela estacao
CREATE TABLE estacao (
    id_estacao UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT,
    latitude FLOAT8,
    longitude FLOAT8,
    timestamp TIMESTAMP
);

-- Inserção de dados na tabela estacao
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 2', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 3', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 4', -23.5718, -46.708);
INSERT INTO estacao (name, latitude, longitude) VALUES ('Sensor Name 5', -23.5718, -46.708);

-- Criação da tabela Gas
CREATE TABLE Gas (
    id_gas SERIAL,
    id_estacao UUID REFERENCES estacao(id_estacao),
    co FLOAT8,
    co2 FLOAT8,
    no2 FLOAT8,
    mp10 FLOAT8,
    mp25 FLOAT8,
    timestamp TIMESTAMP
);
