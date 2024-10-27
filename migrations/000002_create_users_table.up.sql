CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    iin VARCHAR(50) NOT NULL UNIQUE,
    phone VARCHAR(50) NOT NULL UNIQUE,
    fullname VARCHAR(50) NOT NULL, 
    total_money int DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
