CREATE TABLE items(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (id),
    `user_id` BIGINT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id),
    `name` VARCHAR(128) NOT NULL,
    `photo` VARCHAR(128) NULL,
    `price` DECIMAL(10,2) NULL,
    `notes` TEXT NULL
)
ENGINE=InnoDB;

