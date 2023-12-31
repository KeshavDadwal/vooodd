CREATE TYPE gender_enum AS ENUM ('Male', 'Female');

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "gender" gender_enum NOT NULL CHECK (gender IN ('Male', 'Female')),
  "dob" DATE NOT NULL,
  "address" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "country_id" int,
  "zip" varchar NOT NULL,
  "mobile_no" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "user_type" varchar DEFAULT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "status" boolean DEFAULT 'false',
  "latitude" DECIMAL(10, 8),
  "longitude" DECIMAL(11, 8),
  "photo" varchar DEFAULT NULL,
  "facebook_id" varchar DEFAULT NULL,
  "google_id" varchar DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login_time" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_users_country FOREIGN KEY (country_id) REFERENCES countries(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
