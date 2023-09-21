CREATE TABLE "import_flows" (
  "id" SERIAL PRIMARY KEY,
  "file_name" varchar(255) NOT NULL,
  "no_rows" int DEFAULT NULL,
  "created" timestamptz NOT NULL DEFAULT (now()),
  "modified" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar(255) NOT NULL DEFAULT 'Uploaded,In Progress,Completed',
  "no_error" int DEFAULT NULL,
  "note" text NOT NULL,
  "log_file" varchar(255) DEFAULT NULL,
  "completed_file" varchar(255) DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
)