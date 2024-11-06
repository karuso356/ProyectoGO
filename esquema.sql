create database ProjectEL;
use ProjectEL;

create table USER(
    idU int not null auto_increment,
    nombreU varchar(255) not null,
    password varchar(255) not null,
    email varchar(255) not null,
    primary key(idU)
);

create table LIBROS(
    idL int not null auto_increment,
    nombreL varchar(255) not null,
    descripcion varchar(255) not null,
    autor varchar(255) not null,
    contenido TEXT not null,
    primary key (idL)
);
