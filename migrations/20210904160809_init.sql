-- +goose Up
CREATE TABLE "user"
(
    id   SERIAL PRIMARY KEY,
    name varchar(128)
);

CREATE TABLE "hobby"
(
    id      uuid PRIMARY KEY,
    name    varchar(128),
    kind    int8,
    user_id int,
    CONSTRAINT "user_id"
        FOREIGN KEY (user_id)
            REFERENCES "user" (id)
);

-- +goose Down
DROP TABLE "hobby";
DROP TABLE "user";
