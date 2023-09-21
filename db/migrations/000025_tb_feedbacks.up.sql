CREATE TABLE "feedbacks" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar NOT NULL,
  "message" varchar NOT NULL,
  "user_id" INT DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_feedbacks_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
