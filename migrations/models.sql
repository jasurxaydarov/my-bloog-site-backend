CREATE DATABASE my_bloog_site_backend;

CREATE TABLE owner(
    fulname         VARCHAR(25)     NOT NULL,
    password        VARCHAR(16)     NOT NULL,
    role            VARCHAR(5)      DEFAULT "owner",
    phone_number    VARCHAR(20)     NOT NULL,
    gmail           VARCHAR(64)     NOT NULL,
    linked_in       VARCHAR(50),
    telegram        VARCHAR(50),
    github          VARCHAR(50),
    leetcode        VARCHAR(50),
    about_me        TEXT       

);

CREATE TABLE categories(

    category_id     UUID            PRIMARY KEY,
    name            VARCHAR(128)    NOT NULL,
    created_at      TIMESTAMP       DEFAULT current_timestamp
);

CREATE TABLE sub_categories(

    sub_category_id UUID            PRIMARY KEY,
    name            VARCHAR(124)    NOT NULL,
    created_at      TIMESTAMP       DEFAULT current_timestamp,
    category_id     UUID            REFERENCES  categories(category_id)
);


CREATE TABLE articles(
    article_id      UUID            PRIMARY KEY,
    title           VARCHAR         NOT NULL,
    content         TEXT            NOT NULL,
    created_at      TIMESTAMP       DEFAULT current_timestamp,
    update_at       TIMESTAMP       DEFAULT NULL,
    deleted_at      TIMESTAMP       DEFAULT NULL,
    category_id     UUID,
    sub_category_id UUID,
    FOREIGN KEY (category_id)   REFERENCES categories(category_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (sub_category_id)   REFERENCES sub_categories(sub_category_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE viwers(
    viwer_id        UUID            PRIMARY KEY,
    fullname        VARCHAR(45)     NOT NULL,
    username        VARCHAR(45)     NOT NULL,
    gmail           VARCHAR(45)     NOT NULL,
    password        VARCHAR(20)     NOT NULL
);

CREATE TABLE coments(
    coment_id       UUID            PRIMARY KEY,
    coment          VARCHAR(512)    NOT NULL,
    created_at      TIMESTAMP       DEFAULT current_timestamp,
    viwer_id        UUID            REFERENCES viwers(viwer_id),
    article_id      UUID            REFERENCES articles(article_id)
);