CREATE TABLE "user_mall_permissions" (
 "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "mall_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_mall_permissions_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_mall_permissions_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 