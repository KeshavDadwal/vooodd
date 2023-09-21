CREATE TABLE "product_competitors" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int NOT NULL,
  "competitor_id" int NOT NULL,
  "store_location_id" int NOT NULL,
  "price" decimal(10,2) NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_competitors_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_competitors_competitor FOREIGN KEY (competitor_id) REFERENCES competitors(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_competitors_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 