CREATE TABLE "user_shoping_lists" (
 "id" SERIAL PRIMARY KEY,
  "user_shoping_list_name_id" int NOT NULL,
  "product_id" int NOT NULL,
  "user_id" int NOT NULL,
  "store_location_id" int NOT NULL,
  "qty" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
   CONSTRAINT fk_user_shoping_lists_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_shoping_lists_store_location FOREIGN KEY (store_location_id) REFERENCES store_locations(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_shoping_lists_store_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT,
   CONSTRAINT fk_user_shoping_lists_user_shoping_list FOREIGN KEY (user_shoping_list_name_id) REFERENCES user_shopping_list_names(id) ON UPDATE CASCADE ON DELETE RESTRICT 
)