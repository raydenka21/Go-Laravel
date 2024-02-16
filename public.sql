/*
 Navicat Premium Data Transfer

 Source Server         : LOCAL POSTGRES
 Source Server Type    : PostgreSQL
 Source Server Version : 150006 (150006)
 Source Host           : localhost:5432
 Source Catalog        : rest
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150006 (150006)
 File Encoding         : 65001

 Date: 16/02/2024 15:53:46
*/


-- ----------------------------
-- Sequence structure for category_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."category_id_seq";
CREATE SEQUENCE "public"."category_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."category_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS "public"."category";
CREATE TABLE "public"."category" (
  "id" int4 NOT NULL DEFAULT nextval('category_id_seq'::regclass),
  "name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
  "deleted_at" timestamp(6),
  "updated_at" timestamp(6),
  "created_at" timestamp(6)
)
;
ALTER TABLE "public"."category" OWNER TO "postgres";

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO "public"."category" ("id", "name", "deleted_at", "updated_at", "created_at") VALUES (2, 'test23131231', NULL, '2024-02-14 09:51:41.652851', '2024-02-13 14:00:23');
INSERT INTO "public"."category" ("id", "name", "deleted_at", "updated_at", "created_at") VALUES (3, 'test1', '2024-02-14 10:30:06.267575', '2024-02-14 07:57:08.379576', '2024-02-14 07:57:08.379576');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."category_id_seq"
OWNED BY "public"."category"."id";
SELECT setval('"public"."category_id_seq"', 3, true);

-- ----------------------------
-- Primary Key structure for table category
-- ----------------------------
ALTER TABLE "public"."category" ADD CONSTRAINT "category_pkey" PRIMARY KEY ("id");
