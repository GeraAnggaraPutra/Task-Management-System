CREATE TABLE IF NOT EXISTS "users" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role_guid" varchar NOT NULL REFERENCES roles(guid) ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

INSERT INTO "users" 
    ("username", "email", "password", "role_guid")
VALUES
    ('Super Admin', 'superadmin@gmail.com', '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC', (SELECT guid FROM roles WHERE name = 'Super Admin')),
    ('Admin', 'admin@gmail.com', '$2a$12$LQi1CpKB/dUNMKko2sHd/.umM9hdOYSoMRF7b8JbgiV3ZvSWIEqQC', (SELECT guid FROM roles WHERE name = 'Admin'))
ON CONFLICT DO NOTHING;