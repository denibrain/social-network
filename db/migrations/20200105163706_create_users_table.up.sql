CREATE TABLE users
(
    id int not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(63) not null,
    surname varchar(63) not null,
    age int not null default 0,
    sex enum('F', 'M', 'U') not null default 'U',
    `login` varchar(63) not null,
    `password` varchar(63) not null

)
    charset=utf8;

