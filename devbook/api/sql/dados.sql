insert into usuarios (nome, nick, email, senha) values ('João', 'joao', 'joao@email.com', '$2a$10$I6dVTkU92R6PKkc5e3dg9uzkxFKqhe0zB5ysm.Nx/djewj5ogNO6G');
insert into usuarios (nome, nick, email, senha) values ('Maria', 'maria', 'maria@email.com', '$2a$10$I6dVTkU92R6PKkc5e3dg9uzkxFKqhe0zB5ysm.Nx/djewj5ogNO6G');
insert into usuarios (nome, nick, email, senha) values ('José', 'jose', 'jose@email.com', '$2a$10$I6dVTkU92R6PKkc5e3dg9uzkxFKqhe0zB5ysm.Nx/djewj5ogNO6G');

insert into seguidores(usuario_id, seguidor_id)
  values 
    (1, 2),
    (1, 3),
    (2, 1),
    (3, 1);