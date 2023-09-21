CREATE TABLE "brands_languages" (
  "id" SERIAL PRIMARY KEY,
  "brand_id" INT,
  "language_id" INT,
  "name" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_brands_languages_brand FOREIGN KEY (brand_id) REFERENCES brands(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_brands_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

