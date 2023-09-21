CREATE TABLE "home_banner_languages" (
  "id" SERIAL PRIMARY KEY,
  "home_banner_id" INT DEFAULT NULL,
  "language_id" INT DEFAULT NULL,
  "name" varchar DEFAULT NULL,
  "description" varchar DEFAULT NULL,
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_home_banner_languages_home_banner FOREIGN KEY (home_banner_id) REFERENCES home_banners(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT fk_home_banner_languages_languages FOREIGN KEY (language_id) REFERENCES languages(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

