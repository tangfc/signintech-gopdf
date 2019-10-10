/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-10-09 16:45:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_member
-- ----------------------------
DROP TABLE IF EXISTS `tb_member`;
CREATE TABLE `tb_member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(18) NOT NULL COMMENT '用户名',
  `nickname` varchar(50) NOT NULL COMMENT '昵称',
  `cellphone` char(11) NOT NULL COMMENT '手机号码',
  `realname` varchar(50) NOT NULL COMMENT '备注名',
  `sex` int(11) unsigned NOT NULL COMMENT '性别',
  `signature` varchar(255) DEFAULT NULL COMMENT '签名',
  `province` varchar(50) NOT NULL COMMENT '省份',
  `city` varchar(50) NOT NULL COMMENT '市',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='会员表';

-- ----------------------------
-- Records of tb_member
-- ----------------------------
INSERT INTO `tb_member` VALUES ('1', '千手斗罗', '千手修罗', '13100000001', '唐三', '1', '过去的，别再遗憾；未来的，无须忧虑；现在的，加倍珍惜。', '安徽', '芜湖', '2019-07-12 10:11:52');
INSERT INTO `tb_member` VALUES ('2', '柔骨斗罗', '柔骨魅兔', '13100000002', '小舞', '2', '不乱于心，不困于情。不畏将来，不念过往。如此，安好。', '云南', '昆明', '2019-07-12 12:14:33');
INSERT INTO `tb_member` VALUES ('3', '白虎斗罗', '邪眸白虎', '13100000003', '戴沐白', '1', '.心若计较，时时都有怨言；心若宽容，处处都是晴天。', '浙江', '杭州', '2019-07-12 12:39:16');
INSERT INTO `tb_member` VALUES ('4', '食神斗罗', '香肠专卖', '13100000004', '奥斯卡', '1', '有一天你会明白，善良比聪明更难。聪明是一种天赋，而善良是一种选择。', '江苏', '苏州', '2019-07-12 13:42:05');
INSERT INTO `tb_member` VALUES ('5', '凤凰斗罗', '邪火凤凰', '13100000005', '马红俊', '1', '不要在小人小事上浪费时间，将军有剑，不斩苍蝇。', '山西', '太原', '2019-07-12 18:32:20');
INSERT INTO `tb_member` VALUES ('6', '九彩斗罗', '九宝琉璃', '13100000006', '宁荣荣', '2', '所有人生的过渡，当时是百般艰难，当有一天蓦然回首，原来已经是飞渡了千山。', '内蒙古', '呼和浩特', '2019-07-15 22:31:10');
INSERT INTO `tb_member` VALUES ('7', '幽冥斗罗', '幽冥灵猫', '13100000007', '朱竹清', '1', '我不需要多么完美的爱情，我只需要有一个人，永远不会放弃我。', '湖南', '长沙', '2019-10-09 10:18:13');
