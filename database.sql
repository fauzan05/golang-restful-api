create database golang_restful_api;
create database golang_restful_api_test;

use golang_restful_api;
use golang_restful_api_test;

create table categories 
(
	id integer primary key auto_increment,
    name varchar(100) not null
) engine = InnoDB;

select * from categories;