CREATE TABLE authors
(
    id   int PRIMARY KEY AUTO_INCREMENT,
    first_name  varchar(255) not null,
    second_name varchar(255) not null

);

CREATE TABLE books
(
    id          int PRIMARY KEY AUTO_INCREMENT,
    title       varchar(255) not null,
    description varchar(512),
    author_id   int  not null unique,
    FOREIGN KEY (author_id) REFERENCES authors (id)

);
