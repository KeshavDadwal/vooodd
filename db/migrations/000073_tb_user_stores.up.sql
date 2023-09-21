CREATE TABLE "user_stores" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "store_location_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_stores_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_stores_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT
)