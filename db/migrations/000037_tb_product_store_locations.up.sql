CREATE TABLE "product_store_locations" (
  "id" SERIAL PRIMARY KEY,
  "product_id" INT DEFAULT NULL,
  "store_location_id" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_store_locations_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_store_locations_store_location FOREIGN KEY (store_location_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

