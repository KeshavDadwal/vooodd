CREATE TABLE "store_location_offers" (
  "id" SERIAL PRIMARY KEY,
  "store_location_id" int NOT NULL,
  "product_id" int NOT NULL,
  "offer_id" int NOT NULL,
  "price" decimal(10,2) DEFAULT NULL,
  "offered_price" decimal(10,2) DEFAULT NULL,
  "is_percent" boolean DEFAULT 'false',
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "sequence" int DEFAULT NULL,
  "isblocked" boolean DEFAULT 'false',
  "isdeleted" boolean DEFAULT 'false',
  "priority" boolean DEFAULT 'false',
  "weight" varchar(255) DEFAULT NULL,
  "unit_id" int NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_store_location_offers_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_store_location_offers_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_store_location_offers_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_store_location_offers_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 