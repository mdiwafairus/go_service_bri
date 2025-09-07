CREATE TABLE clients (    
    id SERIAL PRIMARY KEY,
    client_id VARCHAR UNIQUE,
    client_secret VARCHAR,
    name VARCHAR,
    client_type TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);