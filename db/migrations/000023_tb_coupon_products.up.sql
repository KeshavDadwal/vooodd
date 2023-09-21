CREATE TABLE "coupon_products" (
  "id" SERIAL PRIMARY KEY,
  "coupon_id" INT DEFAULT NULL,
  "product_id" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_coupon_products_coupon FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupon_products_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

