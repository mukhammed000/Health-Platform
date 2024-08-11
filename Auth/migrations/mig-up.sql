CREATE TYPE gender_enum AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS users (
    user_id UUID UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth VARCHAR(20) NOT NULL,
    gender gender_enum NOT NULL,
    role VARCHAR(50) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at VARCHAR(250) NOT NULL,
    updated_at VARCHAR(250) NOT NULL,
    deleted_at INT DEFAULT 0,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS tokens (
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    token_id UUID UNIQUE NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    expires_at VARCHAR(250) NOT NULL,
    created_at VARCHAR(250) NOT NULL,
    PRIMARY KEY (token_id)
);