CREATE TABLE "mall_languages" (
 "id" SERIAL PRIMARY KEY,
  "mall_id" int NOT NULL,
  "language_id" int NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "zip" varchar NOT NULL,
  "phone_no" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_mall_languages_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_malls_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 