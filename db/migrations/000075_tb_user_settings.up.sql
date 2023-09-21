CREATE TABLE "user_settings" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "mall_radius" int DEFAULT NULL,
  "store_radius" int DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_settings_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 