CREATE TABLE "content_managements" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "title" varchar DEFAULT NULL,
  "content" varchar DEFAULT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now())
);

