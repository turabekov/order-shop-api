CREATE TABLE IF NOT EXISTS "customers" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone" VARCHAR,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);