CREATE TABLE "user_actions" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "save_type" varchar(255) NOT NULL,
  "tablename" varchar(255) NOT NULL,
  "autoincrement_id" int DEFAULT NULL,
  "fieldname" varchar(255) NOT NULL,
  "value" text NOT NULL,
  "old_value" text DEFAULT NULL,
  "note" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_actions_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 