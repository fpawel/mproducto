DO $$
    BEGIN
        CREATE TYPE USER_ROLE AS ENUM ('admin', 'manager', 'regular_user');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END $$;

CREATE TABLE IF NOT EXISTS credential
(
    credential_id SERIAL PRIMARY KEY NOT NULL,
    name          TEXT               NOT NULL,
    pass          TEXT               NOT NULL,
    email         TEXT               NOT NULL,
    user_role     USER_ROLE          NOT NULL,

    CONSTRAINT proper_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_name UNIQUE (name),
    CONSTRAINT proper_name CHECK ( char_length(name) BETWEEN 5 AND 20 ),
    CONSTRAINT proper_pass CHECK ( char_length(pass) BETWEEN 5 AND 20 )
);

INSERT INTO credential(name, pass, email, user_role)
VALUES ('pawel', '11111', 'binf1611@gmail.com', 'admin'),
       ('alexey', '22222', 'zorchenkov@gmail.com', 'admin');