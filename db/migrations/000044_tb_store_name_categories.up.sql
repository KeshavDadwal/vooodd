CREATE TABLE "store_name_categories" (
  "id" SERIAL PRIMARY KEY,
  "parent_id" INT DEFAULT NULL,
  "lft" INT DEFAULT NULL,
  "rght" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_deals_parent FOREIGN KEY (parent_id) REFERENCES store_name_categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

