-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: kla_crm
-- ------------------------------------------------------
-- Server version	8.0.36

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
-- Table structure for table `batches`
--

DROP TABLE IF EXISTS `batches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `batches` (
  `batch_code` varchar(200) NOT NULL,
  `date` date NOT NULL,
  `dispatched` tinyint(1) NOT NULL,
  `entity` varchar(10) NOT NULL,
  KEY `batch_code` (`batch_code`),
  KEY `entity` (`entity`),
  CONSTRAINT `batches_ibfk_1` FOREIGN KEY (`batch_code`) REFERENCES `incoming_raw_material` (`lot_number`) ON DELETE CASCADE,
  CONSTRAINT `batches_ibfk_2` FOREIGN KEY (`entity`) REFERENCES `raw_material_code` (`entity_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `batches`
--

LOCK TABLES `batches` WRITE;
/*!40000 ALTER TABLE `batches` DISABLE KEYS */;
INSERT INTO `batches` VALUES ('BE/30-03-24','2030-03-24',0,'BE');
/*!40000 ALTER TABLE `batches` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `checkpoints`
--

LOCK TABLES `checkpoints` WRITE;
/*!40000 ALTER TABLE `checkpoints` DISABLE KEYS */;
INSERT INTO `checkpoints` VALUES (9,'blanching_area'),(10,'cold_storage_area'),(11,'dispatch_area'),(8,'incoming_raw_material');
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
  CONSTRAINT `employee_checkpoint_ibfk_1` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`employee_id`) ON DELETE CASCADE,
  CONSTRAINT `employee_checkpoint_ibfk_2` FOREIGN KEY (`checkpoint_id`) REFERENCES `checkpoints` (`checkpoint_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee_checkpoint`
--

LOCK TABLES `employee_checkpoint` WRITE;
/*!40000 ALTER TABLE `employee_checkpoint` DISABLE KEYS */;
INSERT INTO `employee_checkpoint` VALUES (5,8,'2024-03-29 21:36:57'),(5,9,'2024-03-29 21:37:02'),(6,8,'2024-03-29 21:37:12'),(6,9,'2024-03-29 21:37:17'),(6,11,'2024-03-29 21:37:27'),(7,11,'2024-03-30 16:24:31');
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (4,'$2a$10$IYbbxZYcpWn.CUGhmn3qFOz0oEoR8FVqmyVmq4WaXKX4jxEw.RqX6','fuckyou','1990-02-01 05:30:00.000','2024-03-29 21:31:55.108','packer','Engineering','123 Main St, Anytown','123123123','g@example.com','ADMIN'),(5,'$2a$10$wMjsacpkkAMs8PphSGNUYu6LpIDBf7hLldggtydZSnuERZafUq5eO','w1','1990-02-01 05:30:00.000','2024-03-29 21:35:13.002','packer','Engineering','123 Main St, Anytown','123','g1@example.com','WORKER'),(6,'$2a$10$LQ8.EjGwGqIKvY06M/k8U.SzpjHgePzor6pxP1Lmfxd6A.cZHHVrK','w2','1990-02-01 05:30:00.000','2024-03-29 21:35:34.235','packer','Engineering','123 Main St, Anytown','1234','g2@example.com','WORKER'),(7,'$2a$10$j/CVCHp5zv8/Mkr0lyVdjuUrBG7H4BVtt6NPOaDBpzUqQK1Gyuk7u','w3','1990-02-01 05:30:00.000','2024-03-30 16:20:13.713','packer','Engineering','123 Main St, Anytown','12345','g3@example.com','WORKER'),(8,'$2a$10$6f9L64up3SNbEx4T56Qr.eR2zDhw8aBiA2KFQfh3nZGMHr8rG6fLS','pagal','1990-02-01 05:30:00.000','2024-04-05 00:10:18.420','pkger','Engineering','123 Main St, Anytown','000000','ggg1@example.com','WORKER');
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `incoming_raw_material`
--

DROP TABLE IF EXISTS `incoming_raw_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `incoming_raw_material` (
  `name` varchar(200) NOT NULL,
  `date_of_arrival` date NOT NULL,
  `vehicle_number` varchar(50) NOT NULL,
  `lot_number` varchar(200) NOT NULL,
  `variety` varchar(200) NOT NULL,
  `received_from` varchar(200) NOT NULL,
  `supplier` varchar(200) NOT NULL,
  `weight_supplier` decimal(10,3) NOT NULL,
  `weight_WM` decimal(10,3) NOT NULL,
  `Rate` decimal(10,3) NOT NULL,
  `color` varchar(200) DEFAULT NULL,
  `texture` varchar(200) DEFAULT NULL,
  `size` varchar(200) DEFAULT NULL,
  `maturity` varchar(200) DEFAULT NULL,
  `aroma` varchar(200) DEFAULT NULL,
  `appearance` varchar(200) DEFAULT NULL,
  `weight_accepted` decimal(10,3) NOT NULL,
  `quantity_rejected` decimal(10,3) NOT NULL,
  `remarks` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`lot_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incoming_raw_material`
--

LOCK TABLES `incoming_raw_material` WRITE;
/*!40000 ALTER TABLE `incoming_raw_material` DISABLE KEYS */;
INSERT INTO `incoming_raw_material` VALUES ('beans','2024-03-28','ABC123','BE/30-03-24','Type A','Supplier XYZ','Supplier XYZ',100.250,98.750,50.250,'Green','Smooth','Large','Fully Ripe','Fruity','Good',98.500,1.750,'None');
/*!40000 ALTER TABLE `incoming_raw_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchase_register`
--

DROP TABLE IF EXISTS `purchase_register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `purchase_register` (
  `order_number` int NOT NULL AUTO_INCREMENT,
  `order_date` date NOT NULL,
  `broker_name` varchar(200) NOT NULL,
  `product_name` varchar(200) NOT NULL,
  `condition_of_product` varchar(200) NOT NULL,
  `amount` decimal(10,3) NOT NULL,
  `qty_bags` int NOT NULL,
  `qty_kgs` decimal(10,3) NOT NULL,
  `vehicle_number` varchar(200) NOT NULL,
  `recovery` varchar(200) NOT NULL,
  `lot_number` varchar(200) NOT NULL,
  `date_received` date NOT NULL,
  `reject_reason` varchar(200) DEFAULT NULL,
  `purchased_by` int DEFAULT NULL,
  `remark` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`order_number`),
  KEY `purchased_by` (`purchased_by`),
  CONSTRAINT `purchase_register_ibfk_1` FOREIGN KEY (`purchased_by`) REFERENCES `employees` (`employee_id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchase_register`
--

LOCK TABLES `purchase_register` WRITE;
/*!40000 ALTER TABLE `purchase_register` DISABLE KEYS */;
/*!40000 ALTER TABLE `purchase_register` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `raw_material_code`
--

DROP TABLE IF EXISTS `raw_material_code`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `raw_material_code` (
  `entity_code` varchar(10) NOT NULL,
  `entity` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`entity_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `raw_material_code`
--

LOCK TABLES `raw_material_code` WRITE;
/*!40000 ALTER TABLE `raw_material_code` DISABLE KEYS */;
INSERT INTO `raw_material_code` VALUES ('BE','Beans'),('BG','Bitter Gourd'),('BRO','Broccoli'),('CA','Cauliflower'),('DS','Drumsticks'),('EG','Eggplant'),('GC','Green Chilli'),('GP','Green Peas'),('GRC','Green Capsicum'),('LF','Ladyfinger'),('ON','Onion'),('RC','Red Capsicum'),('SC','Sweet Corn'),('SP','Spinach'),('TO','Tomato'),('YC','Yellow Capsicum');
/*!40000 ALTER TABLE `raw_material_code` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `requests_raised`
--

DROP TABLE IF EXISTS `requests_raised`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `requests_raised` (
  `request_id` int NOT NULL AUTO_INCREMENT,
  `request_from` int NOT NULL,
  `request_description` text NOT NULL,
  `accepted` tinyint(1) DEFAULT '0',
  `accepted_by` int DEFAULT NULL,
  `admin_comment` text,
  `request_date` datetime(3) NOT NULL,
  `resolve_date` datetime(3) DEFAULT NULL,
  `resolved` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`request_id`),
  KEY `request_from` (`request_from`),
  KEY `accepted_by` (`accepted_by`),
  CONSTRAINT `requests_raised_ibfk_1` FOREIGN KEY (`request_from`) REFERENCES `employees` (`employee_id`),
  CONSTRAINT `requests_raised_ibfk_2` FOREIGN KEY (`accepted_by`) REFERENCES `employees` (`employee_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `requests_raised`
--

LOCK TABLES `requests_raised` WRITE;
/*!40000 ALTER TABLE `requests_raised` DISABLE KEYS */;
INSERT INTO `requests_raised` VALUES (5,7,'tareekh pe tareekh',0,NULL,NULL,'2024-04-04 23:58:23.883',NULL,0),(6,8,'8 ki tareekh pe tareekh',0,NULL,NULL,'2024-04-05 00:14:49.828',NULL,0);
/*!40000 ALTER TABLE `requests_raised` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-05  0:55:32
