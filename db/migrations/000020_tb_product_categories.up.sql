CREATE TABLE "product_categories" (
  "id" SERIAL PRIMARY KEY,
  "parent_id" INT DEFAULT NULL,
  "lft" INT DEFAULT NULL,
  "rght" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
  "logo" varchar DEFAULT NULL,
  "store_id" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_categories_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_product_categories_parent FOREIGN KEY (parent_id) REFERENCES product_categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

