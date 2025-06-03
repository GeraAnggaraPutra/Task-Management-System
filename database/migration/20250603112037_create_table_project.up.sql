CREATE TABLE IF NOT EXISTS "projects" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar(255) NOT NULL,
  "description" text,
  "start_date" date,
  "end_date" date,
  "status" varchar(50) NOT NULL DEFAULT 'Active',
  "pic_user_guid" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'system',
  "updated_at" timestamp,
  "updated_by" varchar
);