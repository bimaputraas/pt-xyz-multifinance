-- -------------------------------------------------------------
-- TablePlus 6.1.2(568)
--
-- https://tableplus.com/
--
-- Database: xyz_multifinance
-- Generation Time: 2024-08-18 20:28:21.1310
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `transactions` (
                                `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                `datetime` datetime DEFAULT CURRENT_TIMESTAMP,
                                `nomor_kontrak` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `otr` decimal(15,2) NOT NULL,
                                `admin_fee` decimal(15,2) NOT NULL,
                                `tenor` int NOT NULL,
                                `jumlah_cicilan` decimal(15,2) NOT NULL,
                                `jumlah_bunga` decimal(15,2) NOT NULL,
                                `nama_asset` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `user_id` bigint unsigned NOT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `nomor_kontrak` (`nomor_kontrak`),
                                KEY `user_id` (`user_id`),
                                FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_details` (
                                `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                `datetime` datetime DEFAULT CURRENT_TIMESTAMP,
                                `nik` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `full_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `legal_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `tempat_lahir` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `tanggal_lahir` date NOT NULL,
                                `gaji` decimal(15,2) NOT NULL,
                                `foto_ktp` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `foto_selfie` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                `user_id` bigint unsigned NOT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `nik` (`nik`),
                                KEY `user_id` (`user_id`),
                                FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_limits` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                               `user_id` bigint unsigned NOT NULL,
                               `tenor_1` decimal(15,2) NOT NULL,
                               `tenor_2` decimal(15,2) NOT NULL,
                               `tenor_3` decimal(15,2) NOT NULL,
                               `tenor_4` decimal(15,2) NOT NULL,
                               `datetime` datetime DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               KEY `user_id` (`user_id`),
                               FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `datetime` datetime DEFAULT CURRENT_TIMESTAMP,
                         `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;