CREATE TABLE google_users (
    id SERIAL PRIMARY KEY,
    google_id VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    verified_email BOOLEAN NOT NULL,
    name VARCHAR(255),
    given_name VARCHAR(255),
    family_name VARCHAR(255),
    picture_url VARCHAR(255)
);

CREATE TABLE oauth_tokens (
    id SERIAL PRIMARY KEY,
    google_user_id INTEGER REFERENCES google_users(id),
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255),
    scope VARCHAR(255),
    token_type VARCHAR(50),
    expiry_date TIMESTAMP
);