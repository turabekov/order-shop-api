CREATE TABLE IF NOT EXISTS "products" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "price" NUMERIC CHECK(price >= 0) DEFAULT 0,
    "category_id" UUID NOT NULL REFERENCES categories(id),
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);