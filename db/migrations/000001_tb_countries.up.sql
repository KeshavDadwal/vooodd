CREATE TABLE "countries" (
  "id" SERIAL PRIMARY KEY,
  "iso2" varchar(2) NOT NULL,
  "short_name" varchar NOT NULL,
  "long_name" varchar NOT NULL,
  "iso3" varchar(3) NOT NULL,
  "numcode" varchar(6) NOT NULL,
  "un_member" varchar(12) NOT NULL,
  "calling_code" varchar(8) NOT NULL,
  "cctld" varchar(5) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);