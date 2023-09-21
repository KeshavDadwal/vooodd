CREATE TABLE "stores" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "store_no" varchar DEFAULT NULL,
  "description" varchar NOT NULL,
  "logo" varchar DEFAULT NULL,
  "phone_no" varchar DEFAULT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "user_id" INT DEFAULT NULL,
  "icon" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_stores_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

