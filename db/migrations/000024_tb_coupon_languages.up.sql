CREATE TABLE "coupon_languages" (
  "id" SERIAL PRIMARY KEY,
  "coupon_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "disclaimer" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_coupon_languages_coupon FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_coupon_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

