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

create table seguidores (
    usuarioId int not null,
    seguidoresId int not null,
    foreign key (usuarioId) references usuario(id) on delete cascade,
    foreign key (seguidoresId) references usuario(id) on delete cascade,
    primary key(usuarioId, seguidoresId)
)

create table publicacao (
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autorId int not null,
    autorNick varchar(50) not null,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp(),
    foreign key (autorId) references usuario(id) on delete cascade
);