# sqlboy
a generate SQL curd code tool
![image](img/sqlboy.png)

## Examples

https://github.com/lemon-1997/sqlboy/tree/main/examples

## Installation

```
go install github.com/lemon-1997/sqlboy/cmd/sqlboy@latest
```

## Get started

1. create a new file and add your table statement:

    ```go
    const (
        order = `
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
    `
        product = `
    -- product_info definition
    
    CREATE TABLE 'product_info' (
    'id' int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    'product_id' varchar(20) NOT NULL DEFAULT '' COMMENT '产品编号',
    'sku_id' varchar(20) NOT NULL DEFAULT '' COMMENT 'sku',
    'status' tinyint(3) NOT NULL DEFAULT '0' COMMENT '产品状态',
    'created_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    'updated_at' timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY ('id'),
    UNIQUE KEY 'uk_product' ('product_id', 'sku_id')
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';
    `
    )
    ```

Note: it must be const for compile-time assertion

2. execute command

    ```
    sqlboy <statementFile>.go -mode gorm
    ```

Note: mode can be gorm or sqlx
