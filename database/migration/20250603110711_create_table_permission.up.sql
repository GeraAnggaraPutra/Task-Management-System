CREATE TABLE IF NOT EXISTS "permissions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);

INSERT INTO "permissions" 
    ("name")
VALUES
    ('User Management'),
    ('Role Management'),
    ('Project Management'),
    ('Task Management'),
    ('Audit Logs')
ON CONFLICT DO NOTHING; 