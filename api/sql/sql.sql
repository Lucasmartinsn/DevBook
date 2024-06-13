
create table usuario (
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null,
    criacaoEm timestamp default current_timestamp()
)