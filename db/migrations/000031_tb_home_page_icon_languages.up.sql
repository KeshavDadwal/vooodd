CREATE TABLE "home_page_icon_languages" (
  "id" SERIAL PRIMARY KEY,
  "home_page_icon_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_home_page_icon_languages_home_page_icon FOREIGN KEY (home_page_icon_id) REFERENCES home_page_icons(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_home_page_icon_languages_languages FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

