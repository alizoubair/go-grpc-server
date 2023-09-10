CREATE TABLE "users" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "address" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
)