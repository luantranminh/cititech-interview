-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."assignments" (
  "member_id" uuid references "public"."members",
  "project_id" uuid references "public"."projects",
  CONSTRAINT "assignments_pkey" PRIMARY KEY ("member_id", "project_id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."assignments";