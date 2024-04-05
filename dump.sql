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
INSERT INTO `employee_checkpoint` VALUES (5,8,'2024-03-29 21:36:57'),(5,9,'2024-03-29 21:37:02'),(6,8,'2024-03-29 21:37:12'),(6,9,'2024-03-29 21:37:17'),(6,11,'2024-03-29 21:37:27'),(7,11,'2024-03-30 16:24:31'),(8,8,'2024-04-05 17:27:07'),(9,8,'2024-04-05 18:17:23');
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (4,'$2a$10$IYbbxZYcpWn.CUGhmn3qFOz0oEoR8FVqmyVmq4WaXKX4jxEw.RqX6','fuckyou','1990-02-01 05:30:00.000','2024-03-29 21:31:55.108','packer','Engineering','123 Main St, Anytown','123123123','g@example.com','ADMIN'),(5,'$2a$10$wMjsacpkkAMs8PphSGNUYu6LpIDBf7hLldggtydZSnuERZafUq5eO','w1','1990-02-01 05:30:00.000','2024-03-29 21:35:13.002','packer','Engineering','123 Main St, Anytown','123','g1@example.com','WORKER'),(6,'$2a$10$LQ8.EjGwGqIKvY06M/k8U.SzpjHgePzor6pxP1Lmfxd6A.cZHHVrK','w2','1990-02-01 05:30:00.000','2024-03-29 21:35:34.235','packer','Engineering','123 Main St, Anytown','1234','g2@example.com','WORKER'),(7,'$2a$10$j/CVCHp5zv8/Mkr0lyVdjuUrBG7H4BVtt6NPOaDBpzUqQK1Gyuk7u','w3','1990-02-01 05:30:00.000','2024-03-30 16:20:13.713','packer','Engineering','123 Main St, Anytown','12345','g3@example.com','WORKER'),(8,'$2a$10$6f9L64up3SNbEx4T56Qr.eR2zDhw8aBiA2KFQfh3nZGMHr8rG6fLS','pagal','1990-02-01 05:30:00.000','2024-04-05 00:10:18.420','pkger','Engineering','123 Main St, Anytown','000000','ggg1@example.com','WORKER'),(9,'$2a$10$M6d2tTHrWHHvEVkyUCuT1Oh6Zx2/ByUmEA4nhdvUdPgZ.8NulleSO','subbu','1990-02-01 05:30:00.000','2024-04-05 18:14:34.600','pkger','Engineering','123 Main St, Anytown','456456456','subbu@example.com','SUPERVISOR');
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
  `date_of_arrival` datetime(3) DEFAULT NULL,
  `vehicle_number` varchar(50) NOT NULL,
  `batch_code` varchar(200) NOT NULL,
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
  `weighment_slip_number` varchar(20) DEFAULT NULL,
  `added_by_employee` int NOT NULL,
  KEY `batch_code` (`batch_code`),
  KEY `fk_employee_id` (`added_by_employee`),
  CONSTRAINT `fk_employee_id` FOREIGN KEY (`added_by_employee`) REFERENCES `employees` (`employee_id`),
  CONSTRAINT `incoming_raw_material_ibfk_1` FOREIGN KEY (`batch_code`) REFERENCES `master_tracking` (`batch_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incoming_raw_material`
--

LOCK TABLES `incoming_raw_material` WRITE;
/*!40000 ALTER TABLE `incoming_raw_material` DISABLE KEYS */;
INSERT INTO `incoming_raw_material` VALUES ('Tomato','2024-01-31 17:30:00.000','Your Vehicle Number','TO/05-04-24','Your Variety','Received From','Supplier',123.450,123.450,12.340,'Your Color','Your Texture','Your Size','Your Maturity','Your Aroma','Your Appearance',123.450,123.450,'Your Remarks','',8),('Beans','2024-01-31 17:30:00.000','Your Vehicle Number','BE/05-04-24','Your Variety','Received From','Supplier',123.450,123.450,12.340,'Your Color','Your Texture','Your Size','Your Maturity','Your Aroma','Your Appearance',123.450,123.450,'Your Remarks','',8),('yellow capsicum','2024-01-31 17:30:00.000','Your Vehicle Number','YC/05-04-24','Your Variety','Received From','Supplier',123.450,123.450,12.340,'Your Color','Your Texture','Your Size','Your Maturity','Your Aroma','Your Appearance',123.450,123.450,'Your Remarks','',8);
/*!40000 ALTER TABLE `incoming_raw_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_tracking`
--

DROP TABLE IF EXISTS `master_tracking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_tracking` (
  `active_status` tinyint(1) DEFAULT '1',
  `batch_code` varchar(20) NOT NULL,
  `date_added` datetime(3) NOT NULL,
  `checkpoint_1_passed` tinyint(1) DEFAULT '0',
  `checkpoint_1_checked_by` int DEFAULT NULL,
  `checkpoint_1_verified_by` int DEFAULT NULL,
  `checkpoint_1_clear_date` datetime(3) DEFAULT NULL,
  `checkpoint_2_passed` tinyint(1) DEFAULT '0',
  `checkpoint_2_checked_by` int DEFAULT NULL,
  `checkpoint_2_verified_by` int DEFAULT NULL,
  `checkpoint_2_clear_date` datetime(3) DEFAULT NULL,
  `checkpoint_3_passed` tinyint(1) DEFAULT '0',
  `checkpoint_3_checked_by` int DEFAULT NULL,
  `checkpoint_3_verified_by` int DEFAULT NULL,
  `checkpoint_3_clear_date` datetime(3) DEFAULT NULL,
  `checkpoint_4_passed` tinyint(1) DEFAULT '0',
  `checkpoint_4_checked_by` int DEFAULT NULL,
  `checkpoint_4_verified_by` int DEFAULT NULL,
  `use_by_date` datetime(3) DEFAULT NULL,
  `checkpoint_1_checked` tinyint(1) DEFAULT '0',
  `checkpoint_2_checked` tinyint(1) DEFAULT '0',
  `checkpoint_3_checked` tinyint(1) DEFAULT '0',
  `checkpoint_4_checked` tinyint(1) DEFAULT '0',
  `checkpoint_1_verified` tinyint(1) DEFAULT '0',
  `checkpoint_2_verified` tinyint(1) DEFAULT '0',
  `checkpoint_3_verified` tinyint(1) DEFAULT '0',
  `checkpoint_4_verified` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`batch_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_tracking`
--

LOCK TABLES `master_tracking` WRITE;
/*!40000 ALTER TABLE `master_tracking` DISABLE KEYS */;
INSERT INTO `master_tracking` VALUES (1,'BE/05-04-24','2024-01-31 17:30:00.000',0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,0,0,0,0,0,0,0),(1,'TO/05-04-24','2024-01-31 17:30:00.000',0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,0,0,0,0,0,0,0),(1,'YC/05-04-24','2024-01-31 17:30:00.000',0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,NULL,NULL,NULL,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `master_tracking` ENABLE KEYS */;
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
  `title` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`request_id`),
  KEY `request_from` (`request_from`),
  KEY `accepted_by` (`accepted_by`),
  CONSTRAINT `requests_raised_ibfk_1` FOREIGN KEY (`request_from`) REFERENCES `employees` (`employee_id`),
  CONSTRAINT `requests_raised_ibfk_2` FOREIGN KEY (`accepted_by`) REFERENCES `employees` (`employee_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `requests_raised`
--

LOCK TABLES `requests_raised` WRITE;
/*!40000 ALTER TABLE `requests_raised` DISABLE KEYS */;
INSERT INTO `requests_raised` VALUES (5,7,'tareekh pe tareekh',1,4,'jaa be','2024-04-04 23:58:23.883','2024-04-05 00:59:38.495',1,NULL),(6,8,'8 ki tareekh pe tareekh',0,4,'nooooooooooooooooooooooooo be','2024-04-05 00:14:49.828','2024-04-05 01:01:24.772',1,NULL),(7,8,'fuck you',0,NULL,NULL,'2024-04-06 01:37:18.916',NULL,0,NULL),(8,8,'ghanta bhai',0,NULL,NULL,'2024-04-06 02:03:12.815',NULL,0,NULL),(9,8,'ghanta bhai',0,NULL,NULL,'2024-04-06 02:19:04.699',NULL,0,'title check'),(10,8,'ghanta strrdfighighfghfxdgfhjklhgcxccvhvjbhklhgfg',0,NULL,NULL,'2024-04-06 02:22:04.208',NULL,0,'title check');
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

-- Dump completed on 2024-04-06  4:23:38
