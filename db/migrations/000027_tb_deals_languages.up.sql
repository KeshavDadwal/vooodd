CREATE TABLE "deals_languages" (
  "id" SERIAL PRIMARY KEY,
  "deals_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_deals_languages_deals FOREIGN KEY (deals_id) REFERENCES deals(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_deals_languages_languages FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

