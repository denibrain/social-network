CREATE TABLE friendship_requests
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    author_id int not null default 0,
    invited_id int not null default 0,
    status int not null default 0,
    created_at timestamp not null,
    updated_at timestamp not null
)
    charset=utf8;

CREATE INDEX friendship_requests_my_invites_idx ON friendship_requests (author_id, created_at DESC);
CREATE INDEX friendship_requests_invites_idx ON friendship_requests (invited_id, created_at DESC);

