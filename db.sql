CREATE TABLE IF NOT EXISTS `bb_catalog` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `ident` varchar(255) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `resume` varchar(255) NOT NULL,
    `display_order` integer NOT NULL,
    `img_url` varchar(255) NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS `bb_blog` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `ident` varchar(255) NOT NULL UNIQUE,
   `title` varchar(255) NOT NULL,
   `keywords` varchar(255),
   `catalog_id` bigint NOT NULL,
   `blog_content_id` bigint NOT NULL UNIQUE,
   `blog_content_last_update` bigint NOT NULL,
   `type` tinyint NOT NULL,
   `status` tinyint NOT NULL,
   `views` bigint NOT NULL,
   `created` datetime NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
CREATE INDEX `bb_blog_catalog_id` ON `bb_blog` (`catalog_id`);


CREATE TABLE IF NOT EXISTS `bb_blog_content` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `content` longtext NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
