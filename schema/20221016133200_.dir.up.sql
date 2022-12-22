create table users(
	id serial not null unique ,
	email varchar(255) not null ,
	username varchar(255) not null ,
	passwordHash varchar(255) not null 
);
create table books_list (
	id serial not null unique ,
	title varchar(255) not null,
	description varchar(255)
);
create table users_list (
	id serial not null unique ,
	user_id int references users(id) on delete cascade not null,
	list_id int references books_list(id) on delete cascade not null
);
create table books_items(
	id serial not null unique ,
	title varchar(255) not null,
	description varchar(255),
	done bool not null default false
);
create table list_items (
	id serial not null unique ,
	item_id int references books_items(id) on delete cascade not null,
	list_id int references books_list(id) on delete cascade not null
);
