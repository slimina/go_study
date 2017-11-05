mysql,sqlite3

CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`departname` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)



postgresql

CREATE TABLE userinfo
(
	uid serial NOT NULL,
	username character varying(100) NOT NULL,
	departname character varying(500) NOT NULL,
	Created date,
	CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);

CREATE TABLE userdeatail
(
	uid integer,
	intro character varying(100),
	profile character varying(100)
)
WITH(OIDS=FALSE);