CREATE TABLE "push_notifications" (
  "id" SERIAL PRIMARY KEY,
  "device_id" varchar(255) NOT NULL,
  "store_location_id" int NOT NULL,
  "offer_id" int NOT NULL,
  "is_sent" boolean DEFAULT 'false',
  "is_clicked" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_push_notifications_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_push_notifications_offer FOREIGN KEY (offer_id) REFERENCES offers(id) ON UPDATE CASCADE ON DELETE RESTRICT
) 