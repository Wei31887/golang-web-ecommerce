
DROP TABLE images ;
DROP TABLE cart_items ;
DROP TABLE products ;
DROP TABLE carts ;
DROP TABLE sessions ;
DROP TABLE users ;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  role       integer not null,
  created_at timestamp not null   
);

CREATE TABLE sessions (
  id serial primary key,
  uuid varchar(64) not null unique,
  user_id integer references users(id),
  email varchar(255) not null unique,
  created_at timestamp not null   
);

CREATE TABLE products (
  id serial primary key,
  title varchar(100) not null unique,
  prices NUMERIC(10, 2),
  stocks integer not null,
  sales integer not null,
  created_at timestamp not null   
);

-- cart
CREATE TABLE carts (
  id serial primary key,
  uuid varchar(64) not null unique,
  user_id integer references users(id),
  total_count integer not null,
  total_amount NUMERIC(10, 2) not null,
  created_at timestamp not null    
);

-- cart item
CREATE TABLE cart_items (
  id serial primary key,
  count integer not null,
	amount NUMERIC(10, 2) not null,
  cart_id integer references carts(id),
  product_id integer references products(id)
);

-- image
CREATE TABLE images (
  id  serial primary key,
  image_path 	varchar(50),
	image_name  varchar(50),
  product_id  integer references products(id)
);