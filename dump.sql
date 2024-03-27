-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: localhost    Database: kla_crm
-- ------------------------------------------------------
-- Server version	8.0.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `checkpoints`
--

DROP TABLE IF EXISTS `checkpoints`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `checkpoints` (
  `checkpoint_id` int NOT NULL AUTO_INCREMENT,
  `checkpoint_name` varchar(191) NOT NULL,
  PRIMARY KEY (`checkpoint_id`),
  UNIQUE KEY `checkpoint_name` (`checkpoint_name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `checkpoints`
--

LOCK TABLES `checkpoints` WRITE;
/*!40000 ALTER TABLE `checkpoints` DISABLE KEYS */;
INSERT INTO `checkpoints` VALUES (5,'blanching_area_1'),(3,'cutting_area_1'),(4,'cutting_area_2'),(6,'storage_area_1'),(1,'washing_area_1'),(2,'washing_area_2');
/*!40000 ALTER TABLE `checkpoints` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee_checkpoint`
--

DROP TABLE IF EXISTS `employee_checkpoint`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employee_checkpoint` (
  `employee_id` int DEFAULT NULL,
  `checkpoint_id` int DEFAULT NULL,
  `assigned_at` datetime DEFAULT NULL,
  KEY `employee_id` (`employee_id`),
  KEY `checkpoint_id` (`checkpoint_id`),
  CONSTRAINT `employee_checkpoint_ibfk_1` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`employee_id`),
  CONSTRAINT `employee_checkpoint_ibfk_2` FOREIGN KEY (`checkpoint_id`) REFERENCES `checkpoints` (`checkpoint_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee_checkpoint`
--

LOCK TABLES `employee_checkpoint` WRITE;
/*!40000 ALTER TABLE `employee_checkpoint` DISABLE KEYS */;
INSERT INTO `employee_checkpoint` VALUES (1,1,'2024-03-27 12:34:31');
/*!40000 ALTER TABLE `employee_checkpoint` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employees` (
  `employee_id` int NOT NULL AUTO_INCREMENT,
  `password` longtext NOT NULL,
  `name` longtext NOT NULL,
  `date_of_birth` datetime(3) NOT NULL,
  `date_of_joining` datetime(3) NOT NULL,
  `designation` longtext NOT NULL,
  `department` longtext,
  `address` longtext NOT NULL,
  `phone` varchar(191) NOT NULL,
  `email` varchar(191) DEFAULT NULL,
  `role` longtext NOT NULL,
  PRIMARY KEY (`employee_id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'$2a$10$81mWNHo9uHV115ylIpaKTufCcdiebmI6zZzGgW6hDazGLboNSCinW','shreya','1990-02-01 05:30:00.000','2024-03-27 12:31:32.338','developer','Engineering','123 Main St, Anytown','1234321','shreya@example.com','ADMIN'),(2,'$2a$10$hgYUjCJhE2w.m4cO0QNV6OLMYlZ5/k8kleW6cUHhXlX3u0uxw9kge','vaishnavi','1990-02-01 05:30:00.000','2024-03-27 12:32:52.432','developer','Engineering','123 Main St, ghosttown','345434','vaishnavi@example.com','ADMIN');
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-27 12:48:48
