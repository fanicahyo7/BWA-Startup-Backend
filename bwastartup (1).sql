-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 24, 2021 at 08:34 AM
-- Server version: 10.4.13-MariaDB
-- PHP Version: 7.4.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bwastartup`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'Campaign 1', 'Ini Campaign 1', 'Ini Campaign 1 yaaa', 'anu1,anu2', 0, 10000000, 0, 'campaign-satu', '2021-10-22 11:12:27', '2021-10-22 11:12:21'),
(2, 1, 'Campaign 2', 'Ini Campaign 2', 'Ini Campaign 2', 'anu3,anu4,anu5', 0, 50000000, 0, 'campaign-dua', '2021-10-22 11:15:57', '2021-10-22 11:16:02'),
(3, 2, 'Campaign 3', 'Ini Campaign 3', 'Ini Campaign 3', 'anu6', 0, 75000000, 0, 'campaign-tiga', '2021-10-22 11:16:57', '2021-10-22 11:17:01'),
(4, 1, 'Campaign service', 'campaign', 'Campaign dari service', 'haha 1,haha 2', 0, 120000000, 0, 'campaign-service-1', '2021-11-04 11:22:08', '2021-11-04 11:22:08'),
(5, 1, 'edit dari handler', 'campaign ini edit dari handler', 'campaign coba handler', 'dfgdfg,dsgfgf,sdgfdg', 0, 80000000, 0, 'coba-campaign-dari-handler-1', '2021-11-04 14:04:04', '2021-11-05 13:30:41');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'images/satu.jpg', 0, '2021-10-28 09:02:37', '2021-11-08 10:29:55'),
(2, 1, 'images/dua.jpg', 0, '2021-10-28 09:03:02', '2021-11-08 10:29:55'),
(3, 2, 'images/tiga.jpg', 0, '2021-10-28 09:03:31', '2021-11-08 10:48:49'),
(6, 1, 'images/1-Target Kerja.png', 0, '2021-11-08 10:19:55', '2021-11-08 10:29:55'),
(7, 1, 'images/1-Target Kerja.png', 1, '2021-11-08 10:29:55', '2021-11-08 10:29:55'),
(8, 2, 'images/1-Target Kerja.png', 1, '2021-11-08 10:48:49', '2021-11-08 10:48:49'),
(9, 1, 'images/1-Target Kerja.png', 0, '2021-11-08 11:09:37', '2021-11-08 11:09:37');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 80000, '1', 'cbcvbcx', '2021-11-09 13:27:51', '2021-11-09 13:27:55'),
(2, 1, 2, 40000, '1', 'fbfghfgh', '2021-11-09 14:55:09', '2021-11-09 14:55:12');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Fani Dwi Cahyo', 'programmer', 'fanidc7@gmail.com', '$2a$04$N6s0M9SMcrJr/ZzY.1Den.3CWNhVYkhPSUTfV.V57khJ9SKGSFcBG', 'images/1-satu.png', 'USER', '', '2021-10-19 13:42:32', '2021-10-22 09:04:12'),
(2, 'Melly Aswanti', 'ui', 'melly@gmail.com', '$2a$04$YF4Rw3X22HeZosnuCbdWMeZJY54pmdVU5uCh2cBIuVkVLQB3pRjoi', 'images/2-formatexcel withdrawal.png', 'USER', '', '2021-10-22 08:58:34', '2021-10-22 09:01:51');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
