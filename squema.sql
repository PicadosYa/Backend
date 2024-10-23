create database picadosya;
use picadosya;

create table USERS (
    id int not null auto_increment,
    email VARCHAR(255) not null,
    name VARCHAR(255) not null,
    lastname VARCHAR(255) not null,
    password VARCHAR(255) not null,
    telephone VARCHAR(255) not null,
    profile_photo VARCHAR(255) not null,
    registry_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);

create table ROLES(
    id int not null auto_increment,
    name VARCHAR(255) not null,
    primary key (id)
);

create table USER_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);

INSERT INTO USERS (email, name, lastname, password, telephone, profile_photo)
VALUES ('luis@example.com', 'Luis', 'Martínez', 'hashed_password_here', '5551234567', '/imagenes/perfil/luis.jpg');

insert into ROLES (name) values ("admin");
insert into ROLES (name) values ("cliente");
insert into ROLES (name) values ("canchero");


-- Json para loguearse en localhost:8080/users/login
-- {
--    "email": "pedro@example.com",
--    "password": "yet_another_hashed_password_here"
--}

-- Json para crear usuario en localhost:8080/users/register
--{
--"email": "pedro@example.com",
--"name": "Pedro",
--"lastname": "López",
--"password": "yet_another_hashed_password_here", 
--"telephone": "4567891230",
--"profile_photo": "/imagenes/perfil/pedro.jpg"
--}

