CREATE TABLE "product_category_maps" (
  "id" SERIAL PRIMARY KEY,
  "product_id" INT NOT NULL,
  "product_category_id" INT NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_category_maps_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_category_maps_product_category FOREIGN KEY (product_category_id) REFERENCES product_categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

