CREATE TABLE "offers_categories" (
  "id" SERIAL PRIMARY KEY,
  "deal_id" INT DEFAULT NULL,
  "offer_id" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_offers_categories_deal FOREIGN KEY (deal_id) REFERENCES deals(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_offers_categories_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

