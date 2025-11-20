CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  name varchar(255),
  created_at timestamptz DEFAULT now()
);
