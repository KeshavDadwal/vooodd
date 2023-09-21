CREATE TABLE "store_languages" (
  "id" SERIAL PRIMARY KEY,
  "language_id" INT NOT NULL,
  "store_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_store_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_store_languages_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

