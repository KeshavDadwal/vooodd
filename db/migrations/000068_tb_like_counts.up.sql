CREATE TABLE "like_counts" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "offer_id" int NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_like_counts_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_like_counts_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 