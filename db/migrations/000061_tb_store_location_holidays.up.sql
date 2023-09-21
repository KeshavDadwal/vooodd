CREATE TABLE "store_location_holidays" (
  "id" SERIAL PRIMARY KEY,
  "store_location_id" int NOT NULL,
  "name" varchar NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_store_location_holidays_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 