CREATE TABLE "users" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone_number" VARCHAR, 
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);