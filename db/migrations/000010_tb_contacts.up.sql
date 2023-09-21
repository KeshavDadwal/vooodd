CREATE TABLE "contacts" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar DEFAULT NULL,
  "email" varchar NOT NULL,
  "user_id" INT DEFAULT NULL,
  "subject" varchar DEFAULT NULL,
  "message" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_contacts_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
