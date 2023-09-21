CREATE TABLE "home_banners" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "url_type" varchar NOT NULL,
  "description" varchar NOT NULL,
  "param" varchar DEFAULT NULL,
  "image" varchar NOT NULL,
  "isdefault" boolean DEFAULT 'false',
  "sequence" INT DEFAULT NULL,
  "start_date" DATE NOT NULL,
  "end_date" DATE NOT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

