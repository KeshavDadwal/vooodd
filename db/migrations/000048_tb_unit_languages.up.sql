CREATE TABLE "unit_languages" (
  "id" SERIAL PRIMARY KEY,
  "unit_id" INT NOT NULL,
  "language_id" INT NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_unit_languages_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_unit_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

