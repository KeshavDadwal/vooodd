CREATE TABLE "content_management_mobiles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "title" varchar NOT NULL,
  "link" varchar DEFAULT NULL,
  "link_text" varchar DEFAULT NULL,
  "content" varchar NOT NULL,
  "isblocked" boolean DEFAULT 'false',
  "isdeleted" boolean DEFAULT 'false',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

