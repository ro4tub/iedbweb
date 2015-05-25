DROP TABLE IF EXISTS account;
CREATE TABLE account (
	id int(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	permission INT(11),
	add_time TIMESTAMP NOT NULL,
) engine=innodb;


DROP TABLE IF EXISTS game;
CREATE TABLE game (
	id int(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	url_name VARCHAR(255) NOT NULL,
	pic VARCHAR(255) NOT NULL,
	simple_desc VARCHAR(512) NOT NULL,
	detail_desc TEXT NOT NULL,
	tags VARCHAR(255) NOT NULL,
	platform VARCHAR(255) NOT NULL,
	genre VARCHAR(255) NOT NULL,
	release_date VARCHAR(255) NOT NULL,
	FULLTEXT (simple_desc)
) engine=innodb;


DROP TABLE IF EXISTS game_edit;
CREATE TABLE gameedit (
	id int(11) unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	url_name VARCHAR(255) NOT NULL,
	pic VARCHAR(255) NOT NULL,
	simple_desc VARCHAR(512) NOT NULL,
	detail_desc TEXT NOT NULL,
	tags VARCHAR(255) NOT NULL,
	platform VARCHAR(255) NOT NULL,
	genre VARCHAR(255) NOT NULL,
	release_date VARCHAR(255) NOT NULL,
	FULLTEXT (simple_desc)
) engine=innodb;