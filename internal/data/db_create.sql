DO $$
    BEGIN
        CREATE TYPE USER_ROLE AS ENUM ('admin', 'manager', 'regular_user');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END $$;

CREATE TABLE IF NOT EXISTS user_profile
(
    user_id SERIAL PRIMARY KEY NOT NULL,
    name          TEXT               NOT NULL,
    pass          TEXT               NOT NULL,
    email         TEXT               NOT NULL,
    user_role     USER_ROLE          NOT NULL,

    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_name UNIQUE (name),
    CONSTRAINT proper_name CHECK ( name ~* '^[A-Za-z0-9_\-]{6,20}$' ),
    CONSTRAINT proper_pass CHECK ( pass ~* '^[A-Za-z0-9_\-]{8,20}$' ),
    CONSTRAINT proper_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
);

INSERT INTO user_profile(name, pass, email, user_role)
VALUES ('pawel1', '11111111', 'binf1611@gmail.com', 'admin'),
       ('alexey', '22222222', 'zorchenkov@gmail.com', 'admin') ON CONFLICT DO NOTHING ;


CREATE TABLE IF NOT EXISTS product_category
(
    product_category_id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS product
(
    product_id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    level1 TEXT NOT NULL,
    level2 TEXT NOT NULL,
    level3 TEXT NOT NULL,
    level4 TEXT NOT NULL,
    level5 TEXT NOT NULL,
    level6 TEXT NOT NULL,
    level7 TEXT NOT NULL
);