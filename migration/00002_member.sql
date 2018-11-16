-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."members" (
  "id" uuid NOT NULL,
  "name" text,
  "phone" text,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "updated_at" timestamptz,
  CONSTRAINT "members_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."members";