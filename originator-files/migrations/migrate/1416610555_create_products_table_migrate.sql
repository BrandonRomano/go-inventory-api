CREATE TABLE products(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (id),
    `name` VARCHAR(255) NULL,
    `description` VARCHAR(255) NULL,
    `price` FLOAT NULL
)
ENGINE=InnoDB;
