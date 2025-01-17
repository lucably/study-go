# study-go

## Para a etapa SQL  

**Comando para a criação do container MYSQL**  

_Criacao do Container:_$ docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=senha -e MYSQL_DATABASE=meubanco -e MYSQL_USER=usuario -e MYSQL_PASSWORD=senha_usuario -p 3306:3306 -d mysql:latest

_Criacao do Container com Volume:_ $ docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=senha -v /meu/caminho/local:/var/lib/mysql -p 3306:3306 -d mysql:latest

**No qual:**  
_--name mysql-container:_ Nome do contêiner.  
_-e MYSQL_ROOT_PASSWORD=senha:_ Define a senha para o usuário root.  
_-e MYSQL_DATABASE=meubanco:_ (Opcional) Cria um banco de dados inicial chamado meubanco.  
_-e MYSQL_USER=usuario:_ (Opcional) Cria um usuário adicional chamado usuario.  
_-e MYSQL_PASSWORD=senha_usuario:_ (Opcional) Define a senha para o usuário adicional.  
_-p 3306:3306:_ Mapeia a porta 3306 do contêiner para a porta 3306 do host.  
_-d:_ Executa o contêiner em segundo plano.  
_mysql:latest:_ Usa a versão mais recente da imagem MySQL.  

_-v:_ Substitua /meu/caminho/local por um diretório no seu sistema onde os dados serão armazenados.  

**Gerenciamento do Container**  

_Iniciar o Container:_ docker start my-container  

_Parar o Conteiner:_ docker stop my-container  

_Remover o Conteiner:_ docker rm -f my-container  

**Para conetar**  

_Pelo terminal Admin:_ $ docker exec -it my-container mysql -u root -p  
_Pelo terminal User:_ $ docker exec -it my-container mysql -u usuario -p   


**Comandos dentro do Mysql**  

_Para criar um banco: CREATE DATABASE nome-banco;  

_Para ver os databases_: show databases;  

_Para entrar em um database_: USE name-database;  

_Para criar um usuario: Precisa estar dentro de um banco-de-dados, $ CREATE USER 'usuario'@'localhost' IDENTIFIED BY 'senha'; obs: esta "**localhost**" pois o servidor é local e caso dê o ERROR: Access denied for user 'golang'@'IP' (using password: YES). pega o valor do **IP** e coloque no local do **localhost**.  

_Permissões para o usuario: GRANT ALL PRIVILEGES ON nome-banco.tabela-banco TO 'usuario'@'localhost'; obs: Para dar permissão para todas as tabelas basta digitar nome-banco.*  


### Comando dentro do database  

_Criacao da tabela_: CREATE TABLE usuarios( id int auto_increment primary key, nome varchar(50) not null, email varchar(50) not null)
ENGINE=INNODB;  

_Para ver as tabelas:_ show tables;  

_Para encerrar o mysql:_ exit;  