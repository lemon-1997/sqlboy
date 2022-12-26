// Code generated by sqlboy. DO NOT EDIT.
package sqlx

// compile-time assertion
func _() {
	_ = map[bool]struct {
	}{false: {}, order == `
-- order_info definition

CREATE TABLE 'order_info' (
  'id' int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  'order_id' varchar(20) NOT NULL DEFAULT '' COMMENT '订单号',
  'status' tinyint(3) NOT NULL DEFAULT '0' COMMENT '订单状态',
  'created_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  'updated_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY ('id'),
  UNIQUE KEY 'uk_order' ('order_id')
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';
`: {}}
	_ = map[bool]struct {
	}{false: {}, product == `
-- product_info definition

CREATE TABLE 'product_info' (
  'id' int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  'product_id' varchar(20) NOT NULL DEFAULT '' COMMENT '商品编号',
  'sku_id' varchar(20) NOT NULL DEFAULT '' COMMENT 'sku',
  'status' tinyint(3) NOT NULL DEFAULT '0' COMMENT '商品状态',
  'created_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  'updated_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY ('id'),
  UNIQUE KEY 'uk_product' ('product_id', 'sku_id')
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';
`: {}}
}
