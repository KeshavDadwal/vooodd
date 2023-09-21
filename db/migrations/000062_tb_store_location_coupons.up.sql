CREATE TABLE "store_location_coupons" (
  "id" SERIAL PRIMARY KEY,
  "store_location_id" int NOT NULL,
  "product_id" int NOT NULL,
  "coupon_id" int NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "coupon_price" decimal(10,2) NOT NULL,
  "is_percent" boolean DEFAULT 'false',
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "sequence" int DEFAULT NULL,
  "isblocked" boolean DEFAULT 'false',
  "isdeleted" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_store_location_coupons_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_store_location_coupons_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_store_location_coupons_coupon FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 