CREATE TABLE "user_favourite_stores" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "store_id" INT NOT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_favourite_stores_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_user_favourite_stores_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

