CREATE TABLE "zones" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "state" varchar NOT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

