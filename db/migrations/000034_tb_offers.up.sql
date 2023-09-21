CREATE TABLE "offers" (
  "id" SERIAL PRIMARY KEY,
  "deal_no" varchar DEFAULT NULL,
  "store_id" INT DEFAULT NULL,
  "product_id" INT DEFAULT NULL,
  "name" varchar NOT NULL,
  "path" varchar DEFAULT NULL,
  "description" varchar DEFAULT NULL,
  "isblocked" boolean DEFAULT 'false',
  "isdeleted" boolean DEFAULT 'false',
  "user_id" INT DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_offers_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_offers_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_offers_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

