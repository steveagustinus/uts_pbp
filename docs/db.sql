CREATE DATABASE mp_games;
USE mp_games;

CREATE TABLE accounts (
    id int(11) AUTO_INCREMENT,
    username varchar(16),
    PRIMARY KEY (id)
);

CREATE TABLE games (
    id int(11) AUTO_INCREMENT,
    name varchar(64),
    max_player int(11),
    PRIMARY KEY (id)
);

CREATE TABLE rooms (
    id int(11) AUTO_INCREMENT,
    room_name varchar(16),
    id_game int(11),
    PRIMARY KEY (id),
    FOREIGN KEY (id_game) REFERENCES games (id)
);

CREATE TABLE participants (
    id int(11) AUTO_INCREMENT,
    id_room int(11),
    id_account int(11),
    PRIMARY KEY (id),
    FOREIGN KEY (id_room) REFERENCES rooms (id),
    FOREIGN KEY (id_account) REFERENCES accounts (id)
);

INSERT INTO accounts VALUES
(1, "andi"),
(2, "budi"),
(3, "caca"),
(4, "dedi"),
(5, "evan")

INSERT INTO games VALUES
(1, "Terraria", 3),
(2, "Grand Theft Auto V Online", 16),
(3, "Mobile Legends", 10),
(4, "Forza Horizon 5", 20);

INSERT INTO rooms VALUES
(NULL, "Room 1", 2),
(NULL, "Room 2", 3),
(NULL, "Room 3", 3),
(NULL, "Room 4", 2),
(NULL, "Room 5", 1);

INSERT INTO participants VALUES
(NULL, 1, 1),
(NULL, 2, 2),
(NULL, 3, 3),
(NULL, 4, 4),
(NULL, 5, 5);
