CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    coins         INTEGER      NOT NULL DEFAULT 0
);

CREATE TABLE coin_transactions
(
    id           SERIAL PRIMARY KEY,
    from_user_id INTEGER NOT NULL,
    to_user_id   INTEGER NOT NULL,
    amount       INTEGER NOT NULL,
    CONSTRAINT fk_from_user FOREIGN KEY (from_user_id) REFERENCES users (id),
    CONSTRAINT fk_to_user FOREIGN KEY (to_user_id) REFERENCES users (id)
);

CREATE TABLE inventory
(
    user_id   INTEGER      NOT NULL,
    item_type VARCHAR(255) NOT NULL,
    quantity  INTEGER      NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, item_type),
    CONSTRAINT fk_user_inventory FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE items
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL UNIQUE,
    price INTEGER      NOT NULL
);

INSERT INTO items (name, price)
VALUES ('t-shirt', 80),
       ('cup', 20),
       ('book', 50),
       ('pen', 10),
       ('powerbank', 200),
       ('hoody', 300),
       ('umbrella', 200),
       ('socks', 10),
       ('wallet', 50),
       ('pink-hoody', 500);
