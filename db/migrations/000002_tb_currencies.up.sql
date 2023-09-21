CREATE TABLE "currencies" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "symbol" varchar(10) NOT NULL,
  "isdefault" boolean DEFAULT 'false',
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
