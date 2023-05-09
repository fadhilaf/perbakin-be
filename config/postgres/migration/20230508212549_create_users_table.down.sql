-- migrate create -ext sql create_users_table

DROP TABLE IF EXISTS scorers;
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS supers;

DROP TABLE IF EXISTS users;
