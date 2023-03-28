CREATE TABLE IF NOT EXISTS "couriers" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone_number" VARCHAR,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);