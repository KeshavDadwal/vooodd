CREATE TABLE "content_management_mobile_languages" (
  "id" SERIAL PRIMARY KEY,
  "content_management_mobile_id" INT NOT NULL,
  "language_id" INT NOT NULL,
  "title" varchar NOT NULL,
  "link" varchar NOT NULL,
  "link_text" varchar NOT NULL,
  "content" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_content_management_mobile_languages_content_management_mobile FOREIGN KEY (content_management_mobile_id) REFERENCES content_management_mobiles(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_content_management_mobile_languages_language FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

