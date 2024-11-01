-- MySQL dump 10.13  Distrib 5.7.39, for Win64 (x86_64)
--
-- Host: 172.16.1.254    Database: readline_rns
-- ------------------------------------------------------
-- Server version	5.7.30-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `rd_file`
--

DROP TABLE IF EXISTS `rd_file`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rd_file` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文件id',
  `source_code` varchar(30) DEFAULT '' COMMENT '来源code',
  `source_type` varchar(30) DEFAULT '' COMMENT '来源类型(inspect:请验单附件,task_result:检验结果附件,task_map:检验图谱附件)',
  `file_code` varchar(30) DEFAULT '' COMMENT '文件code',
  `title` varchar(90) DEFAULT '' COMMENT '文件名称',
  `extension` varchar(30) DEFAULT '' COMMENT '扩展名',
  `size` int(11) DEFAULT '0' COMMENT '文件大小',
  `mime_type` varchar(200) DEFAULT '' COMMENT '文件mime类型',
  `origin_name` varchar(200) DEFAULT '' COMMENT '原始名称',
  `rel_path` varchar(200) DEFAULT '' COMMENT '相对路径',
  `file_url` varchar(300) DEFAULT '' COMMENT '文件地址',
  `fms_id` varchar(30) DEFAULT '' COMMENT 'fms文件id',
  `downloads` int(11) DEFAULT '0' COMMENT '下载次数',
  `extra` varchar(255) DEFAULT '' COMMENT '额外信息',
  `created_by` bigint(20) DEFAULT '0' COMMENT '创建人ID',
  `created_name` varchar(50) DEFAULT '' COMMENT '创建人名称',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT '0' COMMENT '更新人ID',
  `updated_name` varchar(50) DEFAULT '' COMMENT '更新人名称',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` bigint(20) DEFAULT '0' COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `undex_code` (`source_code`,`source_type`,`file_code`)
) ENGINE=InnoDB AUTO_INCREMENT=1595 DEFAULT CHARSET=utf8mb4 COMMENT='附件表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-01 14:53:20
