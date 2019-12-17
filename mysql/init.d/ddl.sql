DROP DATABASE IF EXISTS `auth_sample`;
CREATE DATABASE `auth_sample`;

DROP TABLE IF EXISTS `auth_sample`.`users`;
CREATE TABLE IF NOT EXISTS `auth_sample`.`users`(
    `id` VARCHAR(32) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `password` VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);

