CREATE TABLE "product_category_languages" (
  "id" SERIAL PRIMARY KEY,
  "language_id" INT NOT NULL,
  "product_category_id" INT NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_category_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_category_languages_product_category FOREIGN KEY (product_category_id) REFERENCES product_categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

