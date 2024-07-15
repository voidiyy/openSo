CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "password_hash" VARCHAR(255) NOT NULL,
  "donation_sum" float8 NOT NULL ,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamp NOT NULL,
  "profile_image_url" VARCHAR(255)
);

CREATE TABLE "supported_projects" (
  "supporter_id" bigint NOT NULL,
  "project_id" bigint,
  "support_date" timestamp DEFAULT (now()),
  "support_sum" decimal(10,2) NOT NULL,
  PRIMARY KEY (supporter_id,project_id)
);

CREATE TABLE "supported_organizations" (
  "supporter_id" bigint NOT NULL,
  "organization_id" bigint,
  "support_date" timestamp DEFAULT (now()),
  "support_sum" decimal(10,2) NOT NULL,
  PRIMARY KEY (supporter_id,organization_id)
);

CREATE TABLE "authors" (
  "id" serial PRIMARY KEY,
  "author_name" varchar(50) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "payments" varchar(100) NOT NULL,
  "bio" text NOT NULL DEFAULT '',
  "link" varchar(100) NOT NULL DEFAULT '',
  "profile_image_url" varchar(255),
  "additional_info" jsonb,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "projects" (
  "id" serial PRIMARY KEY,
  "author_id" INTEGER NOT NULL,
  "title" VARCHAR(255) NOT NULL DEFAULT '',
  "category" varchar(100),
  "description" text NOT NULL DEFAULT '',
  "details" jsonb,
  "payments" varchar(100),
  "status" VARCHAR(50) NOT NULL DEFAULT 'active',
  "funding_goal" decimal(10,2) DEFAULT 0,
  "funds_raised" decimal(10,2) DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "organizations" (
  "id" serial UNIQUE PRIMARY KEY,
  "category" varchar(100),
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "description" text,
  "website" VARCHAR(100) NOT NULL,
  "contact_email" VARCHAR(100) UNIQUE NOT NULL,
  "logo_url" VARCHAR(255),
  "additional_info" jsonb,
  "created_at" TIMESTAMP DEFAULT (now()),
  "updated_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "categories_p" (
  "id" serial UNIQUE NOT NULL,
  "name" VARCHAR(100) PRIMARY KEY,
  "description" text
);

CREATE TABLE "categories_o" (
  "id" serial UNIQUE NOT NULL,
  "name" VARCHAR(100) PRIMARY KEY,
  "description" text
);

CREATE TABLE "subcategories_p" (
  "id" serial UNIQUE NOT NULL,
  "name" varchar(100) PRIMARY KEY,
  "parent_name" varchar(100),
  "parent_id" bigint,
  "description" text NOT NULL
);

CREATE TABLE "subcategories_o" (
  "id" serial UNIQUE NOT NULL,
  "name" varchar(100) PRIMARY KEY,
  "parent_name" varchar(100),
  "parent_id" bigint,
  "description" text NOT NULL
);

CREATE TABLE "transactions_card" (
  "id" serial PRIMARY KEY,
  "user_id" bigint,
  "author_id" bigint,
  "sender_addr" varchar(255) NOT NULL,
  "receiver_addr" varchar(255) NOT NULL,
  "project_id" INTEGER,
  "amount" DECIMAL(10,2) NOT NULL,
  "transaction_date" timestamp DEFAULT (now()),
  "payment_method" VARCHAR(50)
);

CREATE TABLE "transactions_crypto" (
  "id" serial PRIMARY KEY,
  "user_id" bigint,
  "author_id" bigint,
  "sender_addr" text NOT NULL,
  "receiver_addr" text NOT NULL,
  "network" varchar(255) NOT NULL,
  "tax" decimal(10,4) NOT NULL,
  "project_id" INTEGER,
  "amount" DECIMAL(10,2) NOT NULL,
  "transaction_date" timestamp DEFAULT (now()),
  "payment_method" VARCHAR(50)
);

CREATE TABLE "comments" (
  "id" serial PRIMARY KEY,
  "user_id" bigint,
  "author_id" bigint,
  "post_id" bigint,
  "created_at" timestamp DEFAULT (now()),
  "message" text NOT NULL
);

ALTER TABLE "supported_projects" ADD FOREIGN KEY ("supporter_id") REFERENCES "users" ("id");

ALTER TABLE "supported_projects" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "supported_organizations" ADD FOREIGN KEY ("supporter_id") REFERENCES "users" ("id");

ALTER TABLE "supported_organizations" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("category") REFERENCES "categories_p" ("name");

ALTER TABLE "projects" ADD FOREIGN KEY ("payments") REFERENCES "authors" ("payments");

ALTER TABLE "organizations" ADD FOREIGN KEY ("category") REFERENCES "categories_o" ("name");

ALTER TABLE "subcategories_p" ADD FOREIGN KEY ("parent_name") REFERENCES "categories_p" ("name");

ALTER TABLE "subcategories_p" ADD FOREIGN KEY ("parent_id") REFERENCES "categories_p" ("id");

ALTER TABLE "subcategories_o" ADD FOREIGN KEY ("parent_name") REFERENCES "categories_o" ("name");

ALTER TABLE "subcategories_o" ADD FOREIGN KEY ("parent_id") REFERENCES "categories_o" ("id");

ALTER TABLE "transactions_card" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions_card" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "transactions_card" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "transactions_crypto" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions_crypto" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "transactions_crypto" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "projects" ("id");

