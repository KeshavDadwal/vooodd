CREATE TABLE "share_counts" (
  "id" SERIAL PRIMARY KEY,
  "language_id" INT NOT NULL,
  "product_category_id" INT NOT NULL,
  "share_type" varchar DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_share_counts_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_share_counts_product_category FOREIGN KEY (product_category_id) REFERENCES product_categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

