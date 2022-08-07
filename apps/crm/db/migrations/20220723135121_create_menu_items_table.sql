CREATE TABLE menu_items
(
    `id`         INT          NOT NULL AUTO_INCREMENT,
    `menus_id`   INT          NOT NULL,
    `label`      VARCHAR(255) NOT NULL,
    `path`       VARCHAR(255),
    `ordering`   INT,
    `parent_id`  INT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`menus_id`) REFERENCES menus (id) ON DELETE CASCADE
);