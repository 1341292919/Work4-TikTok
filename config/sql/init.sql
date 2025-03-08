CREATE TABLE casaos.`user`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL ,
    `password` varchar(255) NOT NULL ,
    `avatar_url` varchar(255) NOT NULL DEFAULT "no",
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos.`video`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL ,
    `video_url` varchar(255) NOT NULL ,
    `cover_url` varchar(255) NOT NULL ,
    `title` varchar(255) NOT NULL ,
    `description` varchar(255)NOT NULL ,
    `visit_count` bigint NOT NULL DEFAULT 0,
    `like_count` bigint NOT NULL DEFAULT 0,
    `comment_count` bigint NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos.`comment`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL ,
    `content` varchar(255) NOT NULL ,
    `video_id` bigint NOT NULL ,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos.`user_follows`
(
    `follower_id` bigint NOT NULL ,-- 关注者的用户ID
    `followee_id` bigint NOT NULL ,-- 被关注者的用户ID
    `followed_at` timestamp NOT NULL DEFAULT current_timestamp, -- 关注时间
    PRIMARY KEY (follower_id, followee_id), -- 联合主键，防止重复关注
    FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE, -- 外键约束
    FOREIGN KEY (followee_id) REFERENCES user(id) ON DELETE CASCADE -- 外键约束
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos.`video_likes`
(
    `user_id` bigint NOT NULL , -- 点赞用户的ID
    `video_id` bigint NOT NULL , -- 被点赞的视频ID
    `liked_at` timestamp NOT NULL DEFAULT current_timestamp, -- 点赞时间
    PRIMARY KEY (user_id, video_id), -- 联合主键，防止重复点赞
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE, -- 外键约束
    FOREIGN KEY (video_id) REFERENCES video(id) ON DELETE CASCADE -- 外键约束
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;