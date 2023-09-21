CREATE TYPE type_enum AS ENUM ('Department','ProductCategory','Product','Brand','Filter','More_FIlter','Private','Sku');

CREATE TABLE "coupon_conditions" (
  "id" SERIAL PRIMARY KEY,
  "coupon_id" int NOT NULL,
  "idx" int NOT NULL,
  "type" type_enum DEFAULT 'Product',
  "type_id" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_coupon_conditions_coupon FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 