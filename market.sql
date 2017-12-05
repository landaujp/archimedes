-- Create syntax for TABLE 'bitbank'
CREATE TABLE `bitbank` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'bitflyer'
CREATE TABLE `bitflyer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'btcbox'
CREATE TABLE `btcbox` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'coincheck'
CREATE TABLE `coincheck` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'diff_log'
CREATE TABLE `diff_log` (
  `ex_ask` varchar(255) NOT NULL DEFAULT '',
  `ex_bid` varchar(255) NOT NULL DEFAULT '',
  `ask` float(10,1) NOT NULL,
  `bid` float(10,1) NOT NULL,
  `rate` float(3,2) NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'fisco'
CREATE TABLE `fisco` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'kraken'
CREATE TABLE `kraken` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'quoine'
CREATE TABLE `quoine` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'users'
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `border1` float(3,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create syntax for TABLE 'zaif'
CREATE TABLE `zaif` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
