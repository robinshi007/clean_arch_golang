-- https://www.cybertec-postgresql.com/en/1-to-1-relationship-in-postgresql-for-real/
-- https://www.vertabelo.com/blog/how-to-store-authentication-data-in-a-database-part-3-logging-in-with-external-services/
-- mobile, language, country, time_zome
BEGIN;
  CREATE TABLE IF NOT EXISTS user_accounts (
    uid SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(254) UNIQUE NOT NULL,

    password VARCHAR(200) NOT NULL,
    password_salt VARCHAR(50),
    password_hash_argorithm VARCHAR(50),

    email_confirmation_token VARCHAR(200),
    email_confirmation_expire TIMESTAMPTZ,
    status INTEGER,

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
    );
  CREATE TABLE IF NOT EXISTS user_profiles (
    uid INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(254) NOT NULL,

    first_name VARCHAR(100),
    last_name VARCHAR(100),

    language VARCHAR(100),
    timezone VARCHAR(100),

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
    );
COMMIT;
