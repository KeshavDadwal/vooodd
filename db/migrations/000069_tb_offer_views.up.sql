CREATE TABLE "offer_views" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int NOT NULL,
  "offer_id" int NOT NULL,
  "mall_id" int NOT NULL,
  "store_id" int NOT NULL,
  "store_location_id" int NOT NULL,
  "user_id" int NOT NULL,
  "viewed" int Default NULL,
  "view_date" date NOT NULL,
  "is_clicked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_offer_views_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_offer_views_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_offer_views_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_offer_views_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_offer_views_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_offer_views_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
)