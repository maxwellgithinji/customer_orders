BEGIN;
CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar,
  "code" varchar,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
COMMIT;

BEGIN;
CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "item" varchar NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
COMMIT;

BEGIN;
CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "item_id" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "order_date" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
COMMIT;


BEGIN;
ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
COMMIT;

BEGIN;
ALTER TABLE "orders" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");
COMMIT;

