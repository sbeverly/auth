DROP TABLE IF EXISTS user_account;

CREATE TABLE user_account (
	name VARCHAR(100) NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password VARCHAR(150) NOT NULL,
	is_admin BOOLEAN NOT NULL DEFAULT false
);

