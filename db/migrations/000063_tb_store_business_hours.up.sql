CREATE TABLE "store_business_hours" (
  "id" SERIAL PRIMARY KEY,
  "store_location_id" int NOT NULL,
  "day" varchar(255) NOT NULL,
  "start_time" date NOT NULL,
  "end_time" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_store_business_hours_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 