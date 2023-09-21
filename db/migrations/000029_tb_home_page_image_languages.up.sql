CREATE TABLE "home_page_image_languages" (
  "id" SERIAL PRIMARY KEY,
  "home_page_image_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_home_page_image_languages_home_page FOREIGN KEY (home_page_image_id) REFERENCES home_page_images(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_home_page_image_languages_languages FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

