DROP TABLE IF EXISTS user_account;

CREATE TABLE user_account (
	email TEXT UNIQUE NOT NULL,
	password VARCHAR(50) NOT NULL,
	is_admin BOOLEAN NOT NULL DEFAULT true
);

