CREATE TABLE messages
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    recipient_id int not null default 0,
    message text null,
    author_id int not null default 0,
    status int not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null
)
    charset=utf8;

CREATE INDEX messages_recipient_id_idx ON messages (recipient_id, created_at);
CREATE INDEX messages_author_id_idx ON messages (author_id, created_at)

