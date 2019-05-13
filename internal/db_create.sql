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


CREATE TABLE IF NOT EXISTS product
(
    product_id SERIAL PRIMARY KEY NOT NULL,
    code_sap INTEGER NOT NULL,
    name1 TEXT NOT NULL,
    pack_amount INTEGER NOT NULL,
    status_sap TEXT,
    code1 BIGINT,
    code2 BIGINT,
    name2 TEXT NOT NULL,
    barcode_piece BIGINT,
    barcode_group BIGINT,
    subgroup_id INTEGER NOT NULL,
    subgroup TEXT NOT NULL,
    hierarchy_level1 TEXT NOT NULL,
    hierarchy_level2 TEXT NOT NULL,
    hierarchy_level3 TEXT NOT NULL,
    hierarchy_level4 TEXT NOT NULL,
    hierarchy_level5 TEXT NOT NULL,
    hierarchy_level6 TEXT NOT NULL,
    hierarchy_level7 TEXT NOT NULL
);