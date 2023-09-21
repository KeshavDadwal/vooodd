CREATE TABLE "product_languages" (
  "id" SERIAL PRIMARY KEY,
  "product_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar NOT NULL,
  "sku" varchar NOT NULL,
  "description" varchar DEFAULT NULL,
  "short_description" varchar DEFAULT NULL,
  "weight" decimal(10,2) DEFAULT NULL,
  "meta_title" varchar NOT NULL,
  "meta_keywords" varchar NOT NULL,
  "meta_description" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_languages_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

