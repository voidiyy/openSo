CREATE TABLE "users"
(
    "id"                 serial PRIMARY KEY,
    "username"           VARCHAR(50) UNIQUE  NOT NULL,
    "email"              VARCHAR(100) UNIQUE NOT NULL,
    "password_hash"      VARCHAR(255)        NOT NULL,
    "donation_sum"       float8              NOT NULL default 0,
    "supported_projects" bigint[]            not null default '{}',
    "supported_organizations" bigint[]            not null default '{}',
    "profile_image_url"  VARCHAR(255)        NOT NULL DEFAULT '',
    "created_at"         timestamptz         NOT NULL DEFAULT (now()),
    "last_login"         timestamptz         NOT NULL DEFAULT (now()),
    "updated_at"         timestamptz         NOT NULL DEFAULT (now())
);

CREATE TABLE "authors"
(
    "id"                serial PRIMARY KEY,
    "nick_name"         varchar(50) UNIQUE  NOT NULL,
    "email"             varchar(50) UNIQUE  NOT NULL,
    "password_hash"     varchar(255)        NOT NULL,
    "payments"          varchar(100) not null default '',
    "projects"          text[]              not null default '{}',
    "bio"               text                NOT NULL DEFAULT '',
    "link"              varchar(100)        NOT NULL DEFAULT '',
    "profile_image_url" varchar(255) not null default '',
    "additional_info"   text[]              not null default '{}',
    "created_at"        timestamptz         NOT NULL DEFAULT (now()),
    "last-login"        timestamptz         not null default (now()),
    "updated_at"        timestamptz         not null default (now())
);

CREATE TABLE "projects"
(
    "id"           serial PRIMARY KEY,
    "author_id"    INTEGER     NOT NULL,
    "title"        text        NOT NULL DEFAULT '',
    "category"     text        not null,
    "description"  text        NOT NULL DEFAULT '',
    "link"         text        not null unique,
    "details"      text[]      not null,
    "payments"     text        not null,
    "status"       VARCHAR(50) NOT NULL DEFAULT 'active',
    "supporters"   bigint[]    not null default '{}',
    "funding_goal" float8               DEFAULT 0,
    "funds_raised" float8               DEFAULT 0,
    "created_at"   timestamptz NOT NULL DEFAULT (now()),
    "updated_at"   timestamptz not null default (now())
);

CREATE TABLE "organizations"
(
    "id"              serial UNIQUE PRIMARY KEY,
    "category"        varchar(100)        not null,
    "name"            VARCHAR(255) UNIQUE NOT NULL,
    "description"     text                not null,
    "website"         VARCHAR(100)        NOT NULL,
    "contact_email"   VARCHAR(100) UNIQUE NOT NULL,
    "logo_url"        VARCHAR(255)        not null,
    "additional_info" text[]              not null,
    "supporters"      bigint[]            not null,
    "created_at"      TIMESTAMP DEFAULT (now()),
    "updated_at"      TIMESTAMP DEFAULT (now())
);

CREATE TABLE "transactions_card"
(
    "id"               serial PRIMARY KEY,
    "user_id"          bigint,
    "author_id"        bigint,
    "sender_addr"      varchar(255)   NOT NULL,
    "receiver_addr"    varchar(255)   NOT NULL,
    "project_id"       INTEGER,
    "amount"           DECIMAL(10, 2) NOT NULL,
    "transaction_date" timestamp DEFAULT (now()),
    "payment_method"   VARCHAR(50)
);

CREATE TABLE "transactions_crypto"
(
    "id"               serial PRIMARY KEY,
    "user_id"          bigint,
    "author_id"        bigint,
    "sender_addr"      text           NOT NULL,
    "receiver_addr"    text           NOT NULL,
    "network"          varchar(255)   NOT NULL,
    "tax"              decimal(10, 4) NOT NULL,
    "project_id"       INTEGER,
    "amount"           DECIMAL(10, 2) NOT NULL,
    "transaction_date" timestamp DEFAULT (now()),
    "payment_method"   VARCHAR(50)
);

CREATE TABLE "comments"
(
    "id"         serial PRIMARY KEY,
    "user_id"    bigint,
    "author_id"  bigint,
    "post_id"    bigint not null ,
    "created_at" timestamp DEFAULT (now()),
    "message"    text NOT NULL
);
