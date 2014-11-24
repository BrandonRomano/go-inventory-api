CREATE TABLE sessions(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (id),
    `user_id` BIGINT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id),
    `token` VARCHAR(64) NOT NULL,
    `expiration` int(11) NOT NULL
)
ENGINE=InnoDB;

