CREATE TABLE "content_management_banners" (
  "id" SERIAL PRIMARY KEY,
  "content_management_id" INT ,
  "banner" varchar NOT NULL,
  "banner_text" varchar NOT NULL,
  "banner_link" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_content_management_banners_content_management FOREIGN KEY (content_management_id) REFERENCES content_managements(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

