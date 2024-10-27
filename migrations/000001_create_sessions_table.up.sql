CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(50) NOT NULL UNIQUE,
    status VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
