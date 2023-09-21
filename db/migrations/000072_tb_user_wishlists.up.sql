CREATE TABLE "user_wishlists" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "product_id" int NOT NULL,
  "store_location_id" int DEFAULT NULL,
  "wish_price" float NOT NULL,
  "max_price" float NOT NULL,
  "immediately" boolean Default 'false',
  "date_till" date DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_wishlists_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_wishlists_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_wishlists_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT  
)