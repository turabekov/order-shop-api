CREATE TABLE IF NOT EXISTS "categories" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR ,
    "parent_id" UUID REFERENCES categories(id) DEFAULT NULL,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);