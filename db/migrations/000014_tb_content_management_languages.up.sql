CREATE TABLE "content_management_languages" (
  "id" SERIAL PRIMARY KEY,
  "content_management_id" INT NOT NULL,
  "language_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_content_management_languages_content_management FOREIGN KEY (content_management_id) REFERENCES content_managements(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_content_management_languages_languages FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

