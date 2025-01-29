                                                                                                                                                                                                                           CREATE TABLE users
(
    id            serial       PRIMARY KEY,
    username      varchar(255) NOT NULL UNIQUE,
    name          varchar(255) NOT NULL,
    surname       varchar(255) NOT NULL,
    email         varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE events
(
    id          serial       PRIMARY KEY,
    author_id   integer      NOT NULL,
    title       varchar(255) NOT NULL,
    content     varchar(255) NOT NULL,
    photo_url   varchar(255),
    event_date  TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at  date         NOT NULL,
    updated_at  date         NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TYPE status AS ENUM ('Yes', 'No', 'Pending');

CREATE TABLE participants
(
    user_id           integer      NOT NULL,
    event_id          integer      NOT NULL,
    current_status    status       NOT NULL,
    status_updated_at date         NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, event_id)
);

CREATE TABLE invites
(
    id serial primary key ,
    event_id integer not null ,
    token varchar unique not null,
    created_at timestamp not null default current_timestamp,
    expires_at timestamp not null,
    constraint fk_event foreign key (event_id) references events (id) on delete cascade
);
