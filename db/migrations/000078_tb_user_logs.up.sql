CREATE TABLE "user_logs" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "mall_id" int NOT NULL,
  "store_id" int NOT NULL,
  "offer_id" int NOT NULL,
  "name" varchar(255) NOT NULL,
  "text" varchar(255) NOT NULL,
  "isblocked" boolean Default 'false',
  "isdeleted" boolean Default 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_logs_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_user_logs_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_user_logs_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_user_logs_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT
)