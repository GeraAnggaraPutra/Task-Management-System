CREATE TABLE IF NOT EXISTS "sessions" (
  "guid" varchar UNIQUE PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "user_guid" varchar NOT NULL,
  "access_token" text NOT NULL,
  "access_token_expires_at" timestamp NOT NULL,
  "refresh_token" text NOT NULL,
  "refresh_token_expires_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL
);