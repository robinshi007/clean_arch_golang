CREATE TABLE IF NOT EXISTS redirects(
  id SERIAL PRIMARY KEY NOT NULL,
  code VARCHAR (50) UNIQUE NOT NULL,
  url VARCHAR (255),

  created_by_id INTEGER REFERENCES user_profiles(uid),
  created_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
);
