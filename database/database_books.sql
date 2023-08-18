CREATE DATABASE book;
use book;

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

INSERT INTO authors (first_name,second_name) VALUES ("Федор","Достоевский"), ("Лев","Толстой"), ("Антон","Чехов");

INSERT INTO books (title,description,author_id) 
VALUES ("Преступление и наказание","Про преступление и наказание",1),
("Война и мир","Про войну и мир",2), ("Каштанка","Про собачку",3);