CREATE TABLE "user_events" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_user_events_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

