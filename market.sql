# ************************************************************
# Sequel Pro SQL dump
# バージョン 4499
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# ホスト: 127.0.0.1 (MySQL 5.7.20)
# データベース: market
# 作成時刻: 2017-11-14 08:33:15 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# テーブルのダンプ bitbank
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bitbank`;

CREATE TABLE `bitbank` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ bitflyer
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bitflyer`;

CREATE TABLE `bitflyer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ btcbox
# ------------------------------------------------------------

DROP TABLE IF EXISTS `btcbox`;

CREATE TABLE `btcbox` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ coincheck
# ------------------------------------------------------------

DROP TABLE IF EXISTS `coincheck`;

CREATE TABLE `coincheck` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(7,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ fisco
# ------------------------------------------------------------

DROP TABLE IF EXISTS `fisco`;

CREATE TABLE `fisco` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ kraken
# ------------------------------------------------------------

DROP TABLE IF EXISTS `kraken`;

CREATE TABLE `kraken` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ market
# ------------------------------------------------------------

DROP TABLE IF EXISTS `market`;

CREATE TABLE `market` (
  `exchange` varchar(50) NOT NULL DEFAULT '',
  `last` float(10,1) DEFAULT NULL,
  `depth` text,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`exchange`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ quoine
# ------------------------------------------------------------

DROP TABLE IF EXISTS `quoine`;

CREATE TABLE `quoine` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# テーブルのダンプ zaif
# ------------------------------------------------------------

DROP TABLE IF EXISTS `zaif`;

CREATE TABLE `zaif` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `last` float(10,1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
