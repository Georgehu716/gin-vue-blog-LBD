CREATE TABLE `product` (
    `code` varchar(20) NOT NULL DEFAULT 'x',
    `name` varchar(150) DEFAULT NULL,
    `qty` int(5) DEFAULT NULL,
    `last_updated` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;