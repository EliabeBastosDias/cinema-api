BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  token UUID PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(20) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  last_login TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  active BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO users (token, username, email, password, role, created_at, updated_at, active)
VALUES (
    uuid_generate_v4(),
    'Admin',
    'admin@admin.com',
    '$2a$10$CwTycUXWue0Thq9StjUM0uJ8G89Q/Bi1i.98HrGFHgO/jtp4k8zwW', 
    'admin',
    NOW(),
    NOW(),
    TRUE
);


COMMIT;