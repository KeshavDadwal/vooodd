CREATE TABLE "user_coupons" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "coupon_id" INT NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_coupons_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_user_coupons_coupon FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

