CREATE TABLE "store_categories" (
  "id" SERIAL PRIMARY KEY,
  "deal_id" INT NOT NULL,
  "store_id" INT NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_store_categories_deal FOREIGN KEY (deal_id) REFERENCES deals(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_store_categories_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

