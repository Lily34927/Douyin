create database if not exists douyin;
use douyin;

/*drop table if exists `user`;
create table `user`(
        `id` bigint(20) not null AUTO_INCREMENT,
        `userid` bigint(20) not null,
        `username` varchar(32) collate utf8mb4_general_ci not null unique , -- collate排序
        `password` varchar(32) collate utf8mb4_general_ci not null,
        `follow_count` bigint DEFAULT NULL,
        `follower_count` bigint DEFAULT NULL,
        `is_follow` tinyint(1) DEFAULT NULL,
        `avatar` varchar(255) DEFAULT NULL,
        `background_image` varchar(255) DEFAULT NULL,
        `signature` varchar(255) DEFAULT NULL,
        `total_favorited` varchar(255) DEFAULT NULL,
        `work_count` int DEFAULT NULL,
        `favorite_count` int DEFAULT NULL,
        `create_time` timestamp not null default current_timestamp,
        `update_time` timestamp not null default current_timestamp on update current_timestamp  COMMENT '用户信息更新时间',
        primary key (`id`),
        unique key `idx_username` (`username`) using btree ,
        unique key `idx_user_id` (`userid`) using btree
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
 */

