CREATE TABLE "competitor_languages" (
  "id" SERIAL PRIMARY KEY,
  "competitor_id" INT,
  "language_id" INT,
  "name" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_competitor_languages_competitor FOREIGN KEY (competitor_id) REFERENCES competitors(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_competitor_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

