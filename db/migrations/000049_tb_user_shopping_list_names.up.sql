CREATE TABLE "user_shopping_list_names" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "how_obtain" varchar DEFAULT NULL,
  "frequency" boolean DEFAULT 'false',
  "type" varchar DEFAULT NULL,
  "isnotify" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_shopping_list_name_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

