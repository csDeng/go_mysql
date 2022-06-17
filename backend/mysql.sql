DROP DATABASE IF EXISTS `blog`;
-- 创建数据库
CREATE DATABASE blog default character set utf8mb4
default collate utf8mb4_unicode_ci;

USE blog;


-- 创建表之前先删除旧表
DROP TABLE IF EXISTS `blog_user`;
DROP TABLE IF EXISTS `blog_tag`;
DROP TABLE IF EXISTS `blog_article`;
DROP TABLE IF EXISTS `blog_comment`;
DROP TABLE IF EXISTS `blog_thumb`;
DROP TABLE IF EXISTS `blog_browser`;


DROP TABLE IF EXISTS `blog_tag_user`;
DROP TABLE IF EXISTS `blog_tag_article`;
DROP TABLE IF EXISTS `blog_article_user`;
DROP TABLE IF EXISTS `blog_article_thumb`;
DROP TABLE IF EXISTS `blog_article_comment`;
DROP TABLE IF EXISTS `blog_article_browser`;
DROP TABLE IF EXISTS `blog_user_comment`;




-- 创建用户表

CREATE TABLE `blog_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(256) DEFAULT '' COMMENT '密码',
  `level` char(2) DEFAULT '0' COMMENT '0是普通用户,1是管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';


-- 创建标签表
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签表';


-- 创建文章表
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

-- 创建评论表
CREATE TABLE `blog_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表';


-- 点赞表
CREATE TABLE `blog_thumb` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `thumb` int(10) unsigned  DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='点赞表';

-- 浏览表
CREATE TABLE `blog_browser` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `browser` int(10) unsigned  DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='浏览表';

-- 创建标签用户表
CREATE TABLE `blog_tag_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tid` int(10) unsigned NOT NULL,
    `uid` int(10) unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`tid`) REFERENCES `blog_tag`(`id`),
    FOREIGN KEY(`uid`) REFERENCES `blog_user`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户标签关系表';

-- 创建文章标签关系表
CREATE TABLE `blog_tag_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tid` int(10) unsigned NOT NULL,
    `aid` int(10) unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`tid`) REFERENCES `blog_tag`(`id`),
    FOREIGN KEY(`aid`) REFERENCES `blog_article`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签关系表';

-- 创建文章用户关系表
CREATE TABLE `blog_article_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `aid` int(10) unsigned NOT NULL,
    `uid` int(10) unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`uid`) REFERENCES `blog_user`(`id`),
    FOREIGN KEY(`aid`) REFERENCES `blog_article`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章用户关系表';

-- 文章评论关系表
CREATE TABLE `blog_article_comment` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `aid` int(10) unsigned NOT NULL,
    `cid` int(10) unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`cid`) REFERENCES `blog_comment`(`id`),
    FOREIGN KEY(`aid`) REFERENCES `blog_article`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章评论关系表';


-- 文章点赞关系表
CREATE TABLE `blog_article_thumb` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `aid` int(10) unsigned NOT NULL,
    `thumb_id` int(10) unsigned NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`aid`) REFERENCES `blog_article`(`id`),
    FOREIGN KEY(`thumb_id`) REFERENCES `blog_thumb`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章点赞关系表';

-- 文章浏览量关系表
CREATE TABLE `blog_article_browser` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `aid` int(10) unsigned NOT NULL,
    `browser_id` int(10) unsigned  NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`aid`) REFERENCES `blog_article`(`id`),
    FOREIGN KEY(`browser_id`) REFERENCES `blog_browser`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章浏览关系表';

-- 用户评论关系表
CREATE TABLE `blog_user_comment` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uid` int(10) unsigned NOT NULL,
    `cid` int(10) unsigned  NOT NULL,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`uid`) REFERENCES `blog_user`(`id`),
    FOREIGN KEY(`cid`) REFERENCES `blog_comment`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户评论关系表';



INSERT INTO `blog`.`blog_user` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
