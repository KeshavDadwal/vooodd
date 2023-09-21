CREATE TABLE "mall_views" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "mall_id" int NOT NULL,
  "store_id" int NOT NULL,
  "visited" date NOT NULL,
  "device_id" varchar(255) DEFAULT NULL,
  "from_device" boolean DEFAULT 'false',
  "is_physical" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_mall_views_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_mall_views_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_mall_views_store FOREIGN KEY (store_id) REFERENCES stores(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 