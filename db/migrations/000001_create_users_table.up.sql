CREATE TABLE if not exists users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    document_number BIGINT NOT NULL,
    email VARCHAR(255) NULL
);
