CREATE TABLE harga_pupuk (    
    id SERIAL PRIMARY KEY,
    nama_pupuk VARCHAR UNIQUE,
    harga INTEGER,
    client_type TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);