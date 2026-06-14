CREATE TABLE `t_auction` (
`id` BIGINT UNSIGNED NOT NULL COMMENT '主键',
`seller` VARCHAR(64) NOT NULL COMMENT '出售者',
`nft_contract` VARCHAR(64) NOT NULL COMMENT 'NFT合约地址',
`nft_token_id` BIGINT UNSIGNED NOT NULL COMMENT 'NFT主键',
`start_price` DECIMAL(10, 2) NOT NULL COMMENT '初始价格',
`start_time` DATETIME NOT NULL COMMENT '拍卖开始时间',
`end_time` DATETIME NOT NULL COMMENT '拍卖结束时间',
`status` TINYINT NOT NULL DEFAULT '1' COMMENT '拍卖状态 - 1:拍卖中;2:拍卖结束;3:流拍',
`winner` VARCHAR(64) COMMENT '中拍者',
`win_price` DECIMAL(10, 2) COMMENT '中拍价格',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin COMMENT '拍卖表';

CREATE TABLE `t_auction_place_bid_log` (
`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
`auction_id` BIGINT UNSIGNED NOT NULL COMMENT '拍卖主键',
`bidder` VARCHAR(64) NOT NULL COMMENT '竞拍者',
`bid_price` DECIMAL(10, 2) NOT NULL COMMENT '竞拍价格',
`bid_time` DATETIME NOT NULL COMMENT '竞拍时间',
PRIMARY KEY (`id`),
KEY `idx_auction_id` (`auction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin COMMENT '拍卖竞拍记录表';