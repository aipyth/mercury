CREATE TABLE IF NOT EXISTS users (
    id          BIGSERIAL   PRIMARY KEY,
    created_at  timestamp   NOT NULL DEFAULT NOW(),
    updated_at  timestamp   ,
    email       text        NOT NULL UNIQUE,
    password    text        NOT NULL,
    token       text        NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS timeschemas (
    id          BIGSERIAL   PRIMARY KEY,
    created_at  timestamp   NOT NULL DEFAULT NOW(),
    updated_at  timestamp   ,
    name        text        NOT NULL UNIQUE,
    items       jsonb
);

CREATE TABLE IF NOT EXISTS rooms (
    id              BIGSERIAL   PRIMARY KEY,
    created_at      timestamp   NOT NULL DEFAULT NOW(),
    updated_at      timestamp   ,
    name            text        NOT NULL,
    slug            text        NOT NULL UNIQUE,
    period          int         NOT NULL,
    start_date      timestamp   NOT NULL,
    end_date        timestamp   NOT NULL,
    public          boolean     NOT NULL DEFAULT false,
    timeschema_id   bigint      NOT NULL,
    creator         bigint      NOT NULL,

    FOREIGN KEY (timeschema_id) REFERENCES timeschemas(id),
    FOREIGN KEY (creator) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS subjects (
    id              BIGSERIAL   PRIMARY KEY,
    created_at      timestamp   NOT NULL DEFAULT NOW(),
    updated_at      timestamp   ,
    name            text        NOT NULL,
    days_and_orders jsonb       ,
    room_id         bigint      ,
    lector          text        ,
    extra           text        ,

    FOREIGN KEY (room_id) REFERENCES rooms(id)
);
