CREATE TABLE "coupons" (
  "id" SERIAL PRIMARY KEY,
  "store_id" INT DEFAULT NULL,
  "brand_id" INT DEFAULT NULL,
  "product_id" INT DEFAULT NULL,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "description" varchar DEFAULT NULL,
  "path" varchar DEFAULT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "user_id" INT NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "ispercent" INT NOT NULL,
  "start_date" DATE NOT NULL,
  "end_date" DATE NOT NULL,
  "unit_id" INT DEFAULT NULL,
  "qty" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_coupons_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupons_brand FOREIGN KEY (brand_id) REFERENCES brands(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupons_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupons_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupons_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

