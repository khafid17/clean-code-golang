-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.11-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.3.0.6589
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for esb
CREATE DATABASE IF NOT EXISTS `esb` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `esb`;

-- Dumping structure for table esb.customers
CREATE TABLE IF NOT EXISTS `customers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table esb.customers: ~5 rows (approximately)
INSERT INTO `customers` (`id`, `company`, `created_at`, `updated_at`) VALUES
	(1, 'Company Name 1', '2023-12-26 00:34:12', '2023-12-26 00:34:12'),
	(2, 'Company Name 2', '2023-12-26 00:34:29', '2023-12-26 00:34:29'),
	(3, 'Company Name 3', '2023-12-26 00:34:34', '2023-12-26 00:34:34'),
	(4, 'Company Name 4', '2023-12-26 00:34:39', '2023-12-26 00:34:39'),
	(5, 'Company Name 5', '2023-12-26 00:34:44', '2023-12-26 00:34:44');

-- Dumping structure for table esb.invoices
CREATE TABLE IF NOT EXISTS `invoices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subject` varchar(255) DEFAULT NULL,
  `issue_date` datetime DEFAULT current_timestamp(),
  `due_date` datetime DEFAULT current_timestamp(),
  `address` text DEFAULT NULL,
  `customer` varchar(255) DEFAULT NULL,
  `total_items` int(11) DEFAULT NULL,
  `sub_total` decimal(10,2) DEFAULT NULL,
  `grand_total` decimal(10,2) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table esb.invoices: ~17 rows (approximately)
INSERT INTO `invoices` (`id`, `subject`, `issue_date`, `due_date`, `address`, `customer`, `total_items`, `sub_total`, `grand_total`, `status`, `created_at`, `updated_at`) VALUES

-- Dumping structure for table esb.invoice_items
CREATE TABLE IF NOT EXISTS `invoice_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `invoice_id` int(11) DEFAULT NULL,
  `qty` decimal(10,2) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table esb.invoice_items: ~18 rows (approximately)
INSERT INTO `invoice_items` (`id`, `invoice_id`, `qty`, `price`, `amount`, `created_at`, `updated_at`) VALUES

-- Dumping structure for table esb.items
CREATE TABLE IF NOT EXISTS `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table esb.items: ~4 rows (approximately)
INSERT INTO `items` (`id`, `item_name`, `type`, `created_at`, `updated_at`) VALUES
	(1, 'Design', 'Service', '2023-12-26 00:31:39', '2023-12-26 00:31:51'),
	(2, 'Development', 'Service', '2023-12-26 00:31:39', '2023-12-26 00:31:51'),
	(3, 'Meetings', 'Service', '2023-12-26 00:31:39', '2023-12-26 00:31:51'),
	(4, 'Printer', 'Hardware', '2023-12-26 00:31:39', '2023-12-26 00:31:51'),
	(5, 'Monitor', 'Hardware', '2023-12-26 00:31:39', '2023-12-26 00:31:51');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
