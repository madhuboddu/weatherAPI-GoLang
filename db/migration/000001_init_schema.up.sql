CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "zipcode" varchar NOT NULL,
  "location" varchar NOT NULL,
  "temperature" FLOAT NOT NULL,
  "feels_like" FLOAT NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);