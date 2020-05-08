CREATE TABLE posts
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    user_id int not null default 0,
    message_id int not null default 0,
    message text null,
    author_id int not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null
)
charset=utf8;

CREATE INDEX feeds_user_id_idx ON posts (user_id, created_at)

