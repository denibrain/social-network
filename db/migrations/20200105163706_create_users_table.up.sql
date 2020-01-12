CREATE TABLE users
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    `login` varchar(63) not null,
    `password` varchar(63) not null,
    name varchar(63) not null,
    surname varchar(63) not null,
    city varchar(63) not null,
    age int not null default 0,
    sex enum('F', 'M', 'U') not null default 'U',
    interests text not null
)
    charset=utf8;

CREATE UNIQUE INDEX user_email_idx ON users (login)

