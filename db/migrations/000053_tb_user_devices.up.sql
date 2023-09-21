CREATE TYPE deviceType_enum AS ENUM ('ANDROID', 'IOS');

CREATE TABLE "user_devices" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "device_id" varchar NOT NULL,
  "device_type" deviceType_enum NOT NULL,
  "wuid" varchar DEFAULT NULL,
  "device_token" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_devices_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

