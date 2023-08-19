CREATE DATABASE books;
use books;

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
    author_id   int  not null,
    FOREIGN KEY (author_id) REFERENCES authors (id)

);

INSERT INTO authors (first_name,second_name) VALUES ("Leo","Tolstoy"), ("Anton","Chekhov"), ("Fedor","Dostoevsky");

INSERT INTO books (title,description,author_id) 
VALUES ("War and Peace","About war and peace",1),
("Ward No. 6","About Ward No. 6",2), ("The Brothers Karamazov","About brothers",3);