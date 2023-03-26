CREATE TABLE "customers" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone" VARCHAR,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE "users" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone_number" VARCHAR, 
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "couriers" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "phone_number" VARCHAR,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "products" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR, 
    "price" NUMERIC CHECK(price >= 0) DEFAULT 0,
    "category_id" UUID NOT NULL REFERENCES categories(id),
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "categories" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR ,
    "parent_id" UUID REFERENCES categories(id) DEFAULT NULL,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "orders" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" VARCHAR ,
    "price" NUMERIC CHECK(quantity >= 0) DEFAULT 0,
    "phone_number" VARCHAR,
    "latitude" NUMERIC,
    "longtitude" NUMERIC, 
    "user_id" UUID NOT NULL REFERENCES users(id),
    "customer_id" UUID NOT NULL REFERENCES customers(id),
    "courier_id" UUID REFERENCES couriers(id),
    "product_id" UUID NOT NULL REFERENCES products(id),
    "quantity" NUMERIC CHECK(quantity >= 0) DEFAULT 0,
    "updated_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);