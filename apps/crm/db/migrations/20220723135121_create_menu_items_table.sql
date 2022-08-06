CREATE TABLE menu_items
(
    `id`         INT          NOT NULL AUTO_INCREMENT,
    `menu_id`    INT          NOT NULL,
    `label`      VARCHAR(255) NOT NULL,
    `path`       VARCHAR(255),
    `ordering`   INT,
    `parent_id`  INT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`menu_id`) REFERENCES menus (id) ON DELETE CASCADE
);