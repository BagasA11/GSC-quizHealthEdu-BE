-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 15, 2024 at 01:22 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gsc-quiz-healthedu`
--

-- --------------------------------------------------------

--
-- Table structure for table `blacklist_tokens`
--

CREATE TABLE `blacklist_tokens` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `token` varchar(191) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `blacklist_tokens`
--

INSERT INTO `blacklist_tokens` (`id`, `created_at`, `updated_at`, `deleted_at`, `token`) VALUES
(1, '2024-02-12 09:50:42.455', '2024-02-12 09:50:42.456', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJ1c2VyVXJ1dEtlMyIsInRva2VuVHlwZSI6InVzZXIiLCJleHAiOjE3MDc3MDk3NTB9.pR6pNiWQdF__MFxMvpigjjshmf-vMSprlM0w8ANOPSI'),
(2, '2024-02-12 20:55:53.800', '2024-02-12 20:55:53.800', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NSwidXNlcm5hbWUiOiJhZG1pbnN1cHJlbWUiLCJ0b2tlblR5cGUiOiJhZG1pbiIsImV4cCI6MTcwNzc0OTQ2OH0.IyboiCe1iNRUXYT8bn-5ycFtTy7_OJXvXbrhaghAEis');

-- --------------------------------------------------------

--
-- Table structure for table `options`
--

CREATE TABLE `options` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `alphabet` varchar(1) NOT NULL,
  `text` varchar(200) NOT NULL,
  `question_id` bigint(20) UNSIGNED DEFAULT NULL,
  `color` varchar(30) NOT NULL DEFAULT 'white'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `options`
--

INSERT INTO `options` (`id`, `created_at`, `updated_at`, `deleted_at`, `alphabet`, `text`, `question_id`, `color`) VALUES
(1, '2024-02-12 20:39:34.460', '2024-02-12 20:52:00.474', NULL, 'A', 'Wortel', 2, 'white'),
(2, '2024-02-12 20:39:59.207', '2024-02-12 20:39:59.207', NULL, 'B', 'Tahu', 2, 'red'),
(3, '2024-02-12 20:40:29.168', '2024-02-12 20:40:29.168', NULL, 'C', 'Bayam', 2, 'blue'),
(4, '2024-02-12 20:41:03.616', '2024-02-12 20:41:03.616', NULL, 'D', 'Kangkung', 2, 'white'),
(5, '2024-02-12 20:53:59.337', '2024-02-12 20:53:59.337', NULL, 'D', 'Kangkung', 2, 'white');

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `question` longtext NOT NULL,
  `answer` varchar(1) NOT NULL,
  `quiz_id` bigint(20) UNSIGNED DEFAULT NULL,
  `img` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`id`, `created_at`, `updated_at`, `deleted_at`, `question`, `answer`, `quiz_id`, `img`) VALUES
(2, '2024-02-12 15:33:31.146', '2024-02-12 16:53:39.608', NULL, 'Makanan manakah yang merupakan sumber protein nabati', 'B', 2, 'asset/img/question/2e611d24-bcee-4490-926b-7234266c8dbd.jpg'),
(3, '2024-02-12 16:56:09.508', '2024-02-12 16:56:09.508', '2024-02-12 16:58:10.996', 'Soal gagal', 'A', 2, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `quizzes`
--

CREATE TABLE `quizzes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) NOT NULL,
  `topic` varchar(50) NOT NULL,
  `free` tinyint(1) NOT NULL DEFAULT 1,
  `price` bigint(20) UNSIGNED DEFAULT NULL,
  `disc` int(11) NOT NULL DEFAULT 0,
  `img` longtext DEFAULT NULL,
  `desc` longtext NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `quizzes`
--

INSERT INTO `quizzes` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `topic`, `free`, `price`, `disc`, `img`, `desc`, `verified`) VALUES
(2, '2024-02-07 12:47:45.000', '2024-02-07 12:47:45.000', NULL, 'Coba', 'coba', 1, NULL, 0, NULL, 'percobaan', 0);

-- --------------------------------------------------------

--
-- Table structure for table `scores`
--

CREATE TABLE `scores` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `point` double NOT NULL DEFAULT 0,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `quiz_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `top_ups`
--

CREATE TABLE `top_ups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `quiz_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `qty` bigint(20) UNSIGNED NOT NULL,
  `pay` tinyint(1) NOT NULL DEFAULT 0,
  `cancel` tinyint(1) NOT NULL DEFAULT 0,
  `pay_at` datetime(3) DEFAULT NULL,
  `cancel_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `price` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` longtext NOT NULL,
  `bio` varchar(300) DEFAULT NULL,
  `avatar` varchar(110) DEFAULT NULL,
  `admin` tinyint(1) NOT NULL DEFAULT 0,
  `wallet` int(11) NOT NULL DEFAULT 0,
  `block` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `email`, `password`, `bio`, `avatar`, `admin`, `wallet`, `block`) VALUES
(5, '2024-02-07 13:09:55.000', '2024-02-07 13:09:55.000', NULL, 'adminsupreme', 'admin01@gmail.com', '$2y$10$P3ZC/prw1dWv6u6/U4Q5peOq7CVrkNGw5rxLIpDMz2nN1/MtfQgnq', NULL, NULL, 1, 0, 0),
(6, '2024-02-07 13:19:47.293', '2024-02-07 13:19:47.293', NULL, 'bb3f31d0-9ce3-4e53-86ac-2aaced734fd3', 'admin2@gmail.com', '$2a$10$VggpHER4/UfIhNqR.lkzQOpej0VCojMliCAq.V7zmP66/3EDT87VK', NULL, NULL, 1, 0, 0),
(7, '2024-02-10 23:07:19.261', '2024-02-12 10:39:30.404', NULL, 'userUrutKe3', 'user3@gmail.com', '$2a$10$x10q6MW6U2h7KibYMeu61exyYwaeibDCvIkZ8KE87GfLCJHLYTNlK', NULL, '/asset/img/user/691a1c05-177b-44fd-8ada-ecac1f90e422.png', 0, 0, 0),
(8, '2024-02-12 16:43:02.846', '2024-02-12 16:43:02.846', NULL, 'userFiktif1', 'fictiveuser@gmail.com', '$2a$10$yLWCwJxlcI0UiV/TWX2RP.e9nX6iUbTsWvQO2TjheNMF5EuBRSKVy', NULL, NULL, 0, 0, 0),
(9, '2024-02-15 19:17:04.000', '2024-02-15 19:19:48.000', NULL, 'admin_ke002', 'admin02@gmail.com', '$2y$10$Zsykp1rCn94sknTfGcOSGOzI.AeF6xB9xYWfoyBPjC44QLm9zDmTS', NULL, NULL, 0, 0, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `blacklist_tokens`
--
ALTER TABLE `blacklist_tokens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `token` (`token`),
  ADD KEY `idx_blacklist_tokens_deleted_at` (`deleted_at`);

--
-- Indexes for table `options`
--
ALTER TABLE `options`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_options_deleted_at` (`deleted_at`),
  ADD KEY `fk_questions_option` (`question_id`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_questions_deleted_at` (`deleted_at`),
  ADD KEY `fk_quizzes_question` (`quiz_id`);

--
-- Indexes for table `quizzes`
--
ALTER TABLE `quizzes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `title` (`title`),
  ADD KEY `idx_quizzes_deleted_at` (`deleted_at`);

--
-- Indexes for table `scores`
--
ALTER TABLE `scores`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_scores_deleted_at` (`deleted_at`),
  ADD KEY `fk_quizzes_score` (`quiz_id`),
  ADD KEY `fk_users_score` (`user_id`);

--
-- Indexes for table `top_ups`
--
ALTER TABLE `top_ups`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_top_ups_deleted_at` (`deleted_at`),
  ADD KEY `fk_quizzes_top_up` (`quiz_id`),
  ADD KEY `fk_users_top_up` (`user_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transactions_deleted_at` (`deleted_at`),
  ADD KEY `fk_users_transaction` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username_2` (`username`),
  ADD UNIQUE KEY `email_2` (`email`),
  ADD UNIQUE KEY `username_3` (`username`),
  ADD UNIQUE KEY `email_3` (`email`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `blacklist_tokens`
--
ALTER TABLE `blacklist_tokens`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `options`
--
ALTER TABLE `options`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `quizzes`
--
ALTER TABLE `quizzes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `scores`
--
ALTER TABLE `scores`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `top_ups`
--
ALTER TABLE `top_ups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `options`
--
ALTER TABLE `options`
  ADD CONSTRAINT `fk_questions_option` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`);

--
-- Constraints for table `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `fk_quizzes_question` FOREIGN KEY (`quiz_id`) REFERENCES `quizzes` (`id`);

--
-- Constraints for table `scores`
--
ALTER TABLE `scores`
  ADD CONSTRAINT `fk_quizzes_score` FOREIGN KEY (`quiz_id`) REFERENCES `quizzes` (`id`),
  ADD CONSTRAINT `fk_users_score` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `top_ups`
--
ALTER TABLE `top_ups`
  ADD CONSTRAINT `fk_quizzes_top_up` FOREIGN KEY (`quiz_id`) REFERENCES `quizzes` (`id`),
  ADD CONSTRAINT `fk_users_top_up` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_users_transaction` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
