CREATE TABLE friends
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    user_a_id int not null default 0,
    user_b_id int not null default 0,
    status int not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null
)
    charset=utf8;

CREATE INDEX friends_user_id_idx ON friends (user_a_id, created_at DESC);

