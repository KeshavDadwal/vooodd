CREATE TABLE "product_images" (
  "id" SERIAL PRIMARY KEY,
  "product_id" INT NOT NULL,
  "image" varchar NOT NULL,
  "isdeleted" boolean DEFAULT 'false',
  "isblocked" boolean DEFAULT 'false',
  "isdefault" boolean DEFAULT 'false',
   "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT fk_product_images_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

