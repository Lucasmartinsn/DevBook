create database if not exists devbook;
use devbook;

drop table if exists usuario;
create table usuario (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criacaoEm timestamp default current_timestamp()
);