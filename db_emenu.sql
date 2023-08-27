-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: db_emenu
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-08-26 16:46:48.758','2023-08-26 16:46:48.758',NULL,'Foods'),(2,'2023-08-26 16:46:48.758','2023-08-26 16:46:48.758',NULL,'Beverages'),(3,'2023-08-26 23:25:28.126','2023-08-26 23:25:28.126',NULL,'Deserts');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `description` longtext,
  `price` bigint unsigned DEFAULT NULL,
  `image_url` longtext,
  `category_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_items_deleted_at` (`deleted_at`),
  KEY `fk_categories_items` (`category_id`),
  CONSTRAINT `fk_categories_items` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (1,'2023-08-27 14:20:32.276','2023-08-27 14:20:32.276',NULL,'Hong Kong Boba Tea','Hong Kong Boba Tea Kit for 6',60000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/108725/hong-kong-boba-tea-kit-for-6.63841de36d8e5edfafa13023fc303285.jpg',2),(2,'2023-08-27 14:24:30.301','2023-08-27 14:24:30.301',NULL,'Guy\'s Caliente Margaritas','Guy\'s Caliente Margaritas for 12',120000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/126836/guys-caliente-margaritas-for-12.ca8c6bc06b8f1039549385ffcebc749d.jpg',2),(3,'2023-08-27 14:25:49.949','2023-08-27 14:25:49.949',NULL,'Woodford Reserve Mint Julep Syrup','Woodford Reserve Mint Julep Syrup',15000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/134036/woodford-reserve-mint-julep-syrup.ef523ac7cbae5f4aba6b058207f490d2.jpg',2),(4,'2023-08-27 14:27:40.517','2023-08-27 14:27:40.517',NULL,'Gramercy Tavern Burger','The Gramercy Tavern Burger - 4 Pack',98000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/137148/Gramercy-Tavern-Burger-and-Kielbasa-Kit-6.4.21-72ppi-1x1-15.jpg',1),(5,'2023-08-27 14:29:13.205','2023-08-27 14:29:13.205',NULL,'Shake Shack ShackBurger','Shake Shack ShackBurger – 8 Pack',240000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/134862/shake-shack-shackburger-8-pack.973a5e26836ea86d7e86a327becea2b0.png',1),(6,'2023-08-27 14:30:20.780','2023-08-27 14:30:20.780',NULL,'Gott\'s Complete Cheeseburger','Gott\'s Complete Cheeseburger Kit for 4',84000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/132933/gotts-complete-cheeseburger-kit-for-4.7bdc74104b193427b3fe6eae39e05b5e.jpg',1),(7,'2023-08-27 14:31:55.309','2023-08-27 14:31:55.309',NULL,'German Chocolate Killer Brownie','German Chocolate Killer Brownie',20000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/132029/german-chocolate-killer-brownie-tin-pack.5ebc34160f28767a9d94c4da2e04c4b9.jpg',3),(8,'2023-08-27 14:33:22.463','2023-08-27 14:33:22.463',NULL,'Banana Pudding Bowl','World Famous Banana Pudding Bowl',26000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/84893/world-famous-banana-pudding-bowl-64-oz.85af650f8f51512f8f3181a11d6587d6.jpg',3),(9,'2023-08-27 14:34:59.174','2023-08-27 14:34:59.174',NULL,'Chocolate Chip Cookies','Jacques’ World Famous Chocolate Chip Cookies - 6 Pack',46000,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/81172/jacques-world-famous-chocolate-chip-cookies-6-pack.2217a14c443602493bba88aa9335319a.jpg',3),(10,'2023-08-27 14:35:37.542','2023-08-27 14:35:37.542','2023-08-27 15:08:53.955','Brownies','dummy description',99,'https://goldbelly.imgix.net/uploads/showcase_media_asset/image/81172/jacques-world-famous-chocolate-chip-cookies-6-pack.2217a14c443602493bba88aa9335319a.jpg',3);
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `quantity` bigint unsigned DEFAULT NULL,
  `item_id` bigint unsigned DEFAULT NULL,
  `order_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_details_deleted_at` (`deleted_at`),
  KEY `fk_orders_order_details` (`order_id`),
  KEY `fk_items_order_details` (`item_id`),
  CONSTRAINT `fk_items_order_details` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `fk_orders_order_details` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_details`
--

LOCK TABLES `order_details` WRITE;
/*!40000 ALTER TABLE `order_details` DISABLE KEYS */;
INSERT INTO `order_details` VALUES (1,'2023-08-27 23:44:50.153','2023-08-27 23:44:50.153',NULL,3,1,1),(2,'2023-08-27 23:45:16.424','2023-08-27 23:45:16.424',NULL,2,3,1),(3,'2023-08-27 23:45:36.358','2023-08-27 23:45:36.358',NULL,7,5,1);
/*!40000 ALTER TABLE `order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `order_status` varchar(191) DEFAULT 'waiting for checkout',
  `user_id` char(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_users_orders` (`user_id`),
  CONSTRAINT `fk_users_orders` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-08-27 21:18:19.201','2023-08-27 21:18:19.201',NULL,'waiting for payment','20454fc5-7923-4696-b84c-d29fa57c5f7c'),(2,'2023-08-27 21:22:02.545','2023-08-27 21:22:02.545','2023-08-27 21:23:22.539','cancelled','20454fc5-7923-4696-b84c-d29fa57c5f7c'),(3,'2023-08-27 22:20:38.357','2023-08-27 22:20:38.357',NULL,'waiting for checkout','20454fc5-7923-4696-b84c-d29fa57c5f7c');
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `rating` bigint unsigned DEFAULT NULL,
  `comment` longtext,
  `user_id` char(36) DEFAULT NULL,
  `item_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_reviews_deleted_at` (`deleted_at`),
  KEY `fk_users_reviews` (`user_id`),
  KEY `fk_items_reviews` (`item_id`),
  CONSTRAINT `fk_items_reviews` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `fk_users_reviews` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
INSERT INTO `reviews` VALUES (1,'2023-08-27 15:47:19.902','2023-08-27 15:47:19.902',NULL,4,'Hong kong boba enak','20454fc5-7923-4696-b84c-d29fa57c5f7c',1),(2,'2023-08-27 16:53:14.285','2023-08-27 16:53:14.285',NULL,5,'Margarita enak banget','20454fc5-7923-4696-b84c-d29fa57c5f7c',2),(3,'2023-08-27 16:54:07.881','2023-08-27 16:54:07.881',NULL,5,'Margarita ini enak','1d345c76-c839-4dfa-a687-fc8e2ef4dddb',2),(4,'2023-08-27 16:54:40.388','2023-08-27 16:54:40.388','2023-08-27 16:56:30.347',5,'','1d345c76-c839-4dfa-a687-fc8e2ef4dddb',2),(5,'2023-08-27 17:09:24.200','2023-08-27 17:09:24.200','2023-08-27 17:29:03.486',4,'Margarittaa','1d345c76-c839-4dfa-a687-fc8e2ef4dddb',2);
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `name` longtext NOT NULL,
  `role` varchar(191) DEFAULT 'user',
  `reset_token` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `email_2` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('1d345c76-c839-4dfa-a687-fc8e2ef4dddb','demo1@demo.com','$2a$12$G.Y1OVDyqwijqD1T7wr4ROnI/6Oc6Bo/SJAXabOxv3ChSzoE8O9XS','Demo1','user','','2023-08-26 21:54:34.831','2023-08-26 21:54:34.831',NULL),('20454fc5-7923-4696-b84c-d29fa57c5f7c','demo@demo.com','$2a$12$n19DbDGit1.UcWhTpqXrWeoN8ZNsQjifKBWveOILdU54EE5TSzu9C','Demo Account','user','','2023-08-26 21:53:51.553','2023-08-26 21:53:51.553',NULL),('81a67924-9ba6-4065-81cb-7debe47c3f2c','admin@demo.com','$2a$12$dawFqBxGmEvhpsoSnqVztu7p5xlmprWDIiExPqBN56XOgSkdc2wsO','Admin Account','admin','','2023-08-26 23:21:32.165','2023-08-26 23:21:32.165',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'db_emenu'
--

--
-- Dumping routines for database 'db_emenu'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-28  2:13:29
