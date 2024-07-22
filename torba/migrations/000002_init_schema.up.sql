CREATE TABLE "users"
(
    "user_id"                 bigserial PRIMARY KEY,
    "username"           VARCHAR(50) UNIQUE  NOT NULL,
    "email"              VARCHAR(100) UNIQUE NOT NULL,
    "password_hash"      VARCHAR(255)        NOT NULL,
    "created_at"         timestamptz         NOT NULL DEFAULT 'epoch'::timestamp,
    "last_login"         timestamptz         NOT NULL DEFAULT 'epoch'::timestamp,
    "updated_at"         timestamptz         NOT NULL DEFAULT 'epoch'::timestamp
);

CREATE TABLE "authors"
(
    "author_id"                bigserial PRIMARY KEY,
    "nick_name"         varchar(50) UNIQUE  NOT NULL,
    "email"             varchar(50) UNIQUE  NOT NULL,
    "password_hash"     varchar(255)        NOT NULL,
    "payments"          text not null default '',
    "bio"               text                NOT NULL DEFAULT '',
    "link"              text        NOT NULL DEFAULT '',
    "additional_info"   text              not null default '',
    "created_at"        timestamptz         NOT NULL DEFAULT 'epoch'::timestamp,
    "last_login"        timestamptz         not null default 'epoch'::timestamp,
    "updated_at"        timestamptz         not null default 'epoch'::timestamp
);

CREATE TABLE "projects"
(
    "project_id"   serial      PRIMARY KEY,
    "author_id"    bigint      NOT NULL,
    "title"        text        NOT NULL DEFAULT '',
    "category"     text        not null,
    "sub_category" text        not null,
    "description"  text        NOT NULL DEFAULT '',
    "link"         text        not null unique,
    "details"      text        not null,
    "payments"     text        not null,
    "status"       boolean     NOT NULL DEFAULT true,
    "funding_goal" decimal(10,2)      DEFAULT 0.00,
    "funds_raised" decimal(10,2)    DEFAULT 0.00,
    "created_at"   timestamptz NOT NULL DEFAULT 'epoch'::timestamp,
    "updated_at"   timestamptz not null default 'epoch'::timestamp,
    foreign key (author_id) references authors (author_id) on delete cascade
);

CREATE TABLE "organizations"
(
    "org_id"          serial UNIQUE PRIMARY KEY,
    "author_id"       bigint not null ,
    "category"        varchar(100)        not null,
    "sub_category" text        not null,
    "name"            VARCHAR(50) UNIQUE NOT NULL,
    "description"     text                not null,
    "website"         text        NOT NULL,
    "contact_email"   VARCHAR(50) UNIQUE NOT NULL,
    "logo_url"        text        not null,
    "additional_info" text              not null,
    "created_at"      timestamptz DEFAULT 'epoch'::timestamp,
    "updated_at"      timestamptz DEFAULT 'epoch'::timestamp,
    foreign key (author_id) references authors (author_id) on delete cascade
);

CREATE TABLE project_supporters
(
    entity_type        VARCHAR(20) NOT NULL,     -- Тип сутності: 'user' або 'author'
    entity_id          BIGINT      NOT NULL,     -- ID користувача або автора
    project_id         BIGINT      NOT NULL,
    donation_amount    decimal(10,2) DEFAULT 0.00, -- Сума донату
    donation_date timestamptz not null default 'epoch'::timestamp,                -- Дата останнього донату
    primary key (entity_type, entity_id),
    FOREIGN KEY (project_id) REFERENCES projects (project_id) on delete cascade,
    CONSTRAINT fk_p_supporter_entity_type CHECK (entity_type IN ('user', 'author'))
);

CREATE TABLE org_supporters
(
    entity_type        VARCHAR(20) NOT NULL,     -- Тип сутності: 'user' або 'author'
    entity_id          BIGINT      NOT NULL,     -- ID користувача або автора
    org_id         BIGINT      NOT NULL,
    donation_amount    decimal(10,2) DEFAULT 0.00, -- Сума донату
    donation_date timestamptz not null default 'epoch'::timestamp,                -- Дата останнього донату
    primary key (entity_type, entity_id),
    FOREIGN KEY (org_id) REFERENCES organizations (org_id) on delete cascade ,
    CONSTRAINT fk_o_supporter_entity_type CHECK (entity_type IN ('user', 'author'))
);

CREATE TABLE "transactions_card"
(
    "id"               serial PRIMARY KEY,
    "user_id"          bigint,
    "author_id"        bigint,
    "sender_addr"      varchar(255)   NOT NULL,
    "receiver_addr"    varchar(255)   NOT NULL,
    "project_id"       bigint,
    "amount"           decimal(10,2) NOT NULL,
    "transaction_date" timestamptz DEFAULT 'epoch'::timestamp,
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
    "tax"              decimal(10,2) NOT NULL,
    "project_id"       bigint,
    "amount"           decimal(10,2) NOT NULL,
    "transaction_date" timestamptz DEFAULT 'epoch'::timestamp,
    "payment_method"   VARCHAR(50)
);

CREATE TABLE comments
(
    commentator   VARCHAR(20) NOT NULL,
    post_type    varchar(20)      NOT NULL,
    commentator_id      BIGINT not null ,
    post_id bigint not null ,
    comment text   not null ,
    comment_date timestamptz not null default 'epoch'::timestamp,                -- Дата останнього донату
    primary key (commentator_id, post_id),
    CONSTRAINT fk_commentator_entity_type CHECK (commentator IN ('user', 'author')),
    CONSTRAINT fk_post_entity_type CHECK (post_type IN ('organization', 'project')),
    FOREIGN KEY (post_id) REFERENCES organizations (org_id) on delete cascade
);
