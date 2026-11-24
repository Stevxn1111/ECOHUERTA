/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-12.2.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: ecohuerta
-- ------------------------------------------------------
-- Server version	12.2.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `departamentos`
--

DROP TABLE IF EXISTS `departamentos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `departamentos` (
  `id_departamento` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `slug` varchar(50) NOT NULL,
  PRIMARY KEY (`id_departamento`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departamentos`
--

SET @OLD_AUTOCOMMIT=@@AUTOCOMMIT, @@AUTOCOMMIT=0;
LOCK TABLES `departamentos` WRITE;
/*!40000 ALTER TABLE `departamentos` DISABLE KEYS */;
INSERT INTO `departamentos` VALUES
(1,'Amazonas','amazonas'),
(2,'Antioquia','antioquia'),
(3,'Arauca','arauca'),
(4,'Atlántico','atlantico'),
(5,'Bogotá D.C.','bogota_dc'),
(6,'Bolívar','bolivar'),
(7,'Boyacá','boyaca'),
(8,'Caldas','caldas'),
(9,'Caquetá','caqueta'),
(10,'Casanare','casanare'),
(11,'Cauca','cauca'),
(12,'Cesar','cesar'),
(13,'Chocó','choco'),
(14,'Córdoba','cordoba'),
(15,'Cundinamarca','cundinamarca'),
(16,'Guainía','guainia'),
(17,'Guaviare','guaviare'),
(18,'Huila','huila'),
(19,'La Guajira','la_guajira'),
(20,'Magdalena','magdalena'),
(21,'Meta','meta'),
(22,'Nariño','narino'),
(23,'Norte de Santander','norte_de_santander'),
(24,'Putumayo','putumayo'),
(25,'Quindío','quindio'),
(26,'Risaralda','risaralda'),
(27,'San Andrés y Providencia','san_andres'),
(28,'Santander','santander'),
(29,'Sucre','sucre'),
(30,'Tolima','tolima'),
(31,'Valle del Cauca','valle_del_cauca'),
(32,'Vaupés','vaupes'),
(33,'Vichada','vichada');
/*!40000 ALTER TABLE `departamentos` ENABLE KEYS */;
UNLOCK TABLES;
COMMIT;
SET AUTOCOMMIT=@OLD_AUTOCOMMIT;

--
-- Table structure for table `plantas`
--

DROP TABLE IF EXISTS `plantas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `plantas` (
  `id_planta` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(120) NOT NULL,
  `descripcion` text NOT NULL,
  `caracteristicas` text NOT NULL,
  `cuidados` text NOT NULL,
  `beneficios` text NOT NULL,
  `id_tipo` int(11) NOT NULL,
  `id_departamento` int(11) NOT NULL,
  `imagen` varchar(255) DEFAULT '',
  PRIMARY KEY (`id_planta`),
  KEY `fk_planta_tipo` (`id_tipo`),
  KEY `fk_planta_departamento` (`id_departamento`),
  CONSTRAINT `fk_planta_departamento` FOREIGN KEY (`id_departamento`) REFERENCES `departamentos` (`id_departamento`),
  CONSTRAINT `fk_planta_tipo` FOREIGN KEY (`id_tipo`) REFERENCES `tipos_planta` (`id_tipo`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `plantas`
--

SET @OLD_AUTOCOMMIT=@@AUTOCOMMIT, @@AUTOCOMMIT=0;
LOCK TABLES `plantas` WRITE;
/*!40000 ALTER TABLE `plantas` DISABLE KEYS */;
INSERT INTO `plantas` VALUES
(1,'Frailejón Ernesto Perez (Espeletia)','Planta emblemática de los páramos andinos con hojas densamente cubiertas de vellosidad plateada.','Crecimiento lento (unos pocos centímetros al año), roseta de hojas lanudas y tallo leñoso.','Requiere clima frío de páramo, suelos húmedos con buen drenaje y exposición solar de altura; no trasplantar de su hábitat.','Capta y regula el agua de la niebla y protege los suelos de alta montaña',1,7,''),
(3,'Caléndula','Hierba aromática de flores amarillo-anaranjadas muy usada en huertas educativas y botiquín natural.','Hojas alargadas suaves y capítulos florales brillantes que abren con la luz solar.','Sol directo o semisombra, riego moderado sin encharcar y suelo ligero.','Uso tradicional calmante en piel y atrae polinizadores benéficos.',2,15,''),
(4,'Orquídea Cattleya','Una de las especies ornamentales más representativas de Colombia por el tamaño y color de sus flores.','Pseudobulbos engrosados, hojas carnosas y flores vistosas de aroma suave.','Luz filtrada brillante, buena ventilación, sustrato aireado para epífitas y riego controlado.','Alto valor ornamental y biodiversidad en bosques húmedos.',3,2,''),
(5,'Victoria Regia','Planta acuática amazónica famosa por sus enormes hojas circulares con borde levantado.','Hojas flotantes que pueden superar un metro de diámetro y flores que cambian de color.','Aguas cálidas y tranquilas, abundante luz solar y sustrato rico en materia orgánica.','Refugio para fauna acuática y equilibrio de ecosistemas de humedal.',5,1,''),
(6,'Chontaduro','Palma tropical de racimos con frutos rojizos o anaranjados de pulpa harinosa y nutritiva.','Tallo anillado con espinas y frutos ricos en grasas saludables y betacarotenos.','Clima cálido y húmedo, plena exposición solar y suelos profundos con buena humedad.','Alto aporte energético y nutricional en la gastronomía regional.',4,31,'');
/*!40000 ALTER TABLE `plantas` ENABLE KEYS */;
UNLOCK TABLES;
COMMIT;
SET AUTOCOMMIT=@OLD_AUTOCOMMIT;

--
-- Table structure for table `tipos_planta`
--

DROP TABLE IF EXISTS `tipos_planta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `tipos_planta` (
  `id_tipo` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  PRIMARY KEY (`id_tipo`),
  UNIQUE KEY `nombre` (`nombre`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tipos_planta`
--

SET @OLD_AUTOCOMMIT=@@AUTOCOMMIT, @@AUTOCOMMIT=0;
LOCK TABLES `tipos_planta` WRITE;
/*!40000 ALTER TABLE `tipos_planta` DISABLE KEYS */;
INSERT INTO `tipos_planta` VALUES
(5,'Bosque Tropical'),
(4,'Frutal y Alimenticia'),
(2,'Medicinal'),
(3,'Ornamental'),
(1,'Páramo y Alta Montaña');
/*!40000 ALTER TABLE `tipos_planta` ENABLE KEYS */;
UNLOCK TABLES;
COMMIT;
SET AUTOCOMMIT=@OLD_AUTOCOMMIT;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `usuario` varchar(50) NOT NULL,
  `clave` varchar(255) NOT NULL,
  `rol` enum('estudiante','docente','admin') NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `usuario` (`usuario`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

SET @OLD_AUTOCOMMIT=@@AUTOCOMMIT, @@AUTOCOMMIT=0;
LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(4,'docente','docente','$2a$10$KUuSIOvChrDacezKwZkp1.g2/oeWJ9AuFejzkHXmfmTa50Tj0iwcG','docente'),
(8,'adminnnn','adminnnn','$2a$10$Zc9ilkgEoq.vChfyFJl29Oe.U75w1FwdvEaY52uutU4FpfPMd0u7q','docente'),
(12,'felipe','felipe','$2a$10$soL56NLdfsjSu7eWWRAob.A0Xb7JH46mv7rwCjqnbC1NIsv2u4EEy','estudiante'),
(13,'Brayan','Brayan','$2a$10$JzozCCpfosvyLbtCauCtYe/8QTyv1rQyqNwef5UnN.lMUuzJ2NnAy','admin');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
COMMIT;
SET AUTOCOMMIT=@OLD_AUTOCOMMIT;

--
-- Dumping routines for database 'ecohuerta'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2026-11-24  8:39:39
