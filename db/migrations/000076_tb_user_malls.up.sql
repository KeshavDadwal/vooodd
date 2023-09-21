CREATE TABLE "user_malls" (
   "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "mall_id" int NOT NULL,
  "device_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_malls_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_malls_mall FOREIGN KEY (mall_id) REFERENCES malls(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 