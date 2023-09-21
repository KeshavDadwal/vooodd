CREATE TABLE "units" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

