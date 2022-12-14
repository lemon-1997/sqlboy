// Code generated by sqlboy. DO NOT EDIT.
package sqlx

import "context"

type OrderInfoDao interface {
	CreateOrderInfo(ctx context.Context, orderInfo *OrderInfo) error
	BatchCreateOrderInfo(ctx context.Context, list []*OrderInfo, batchSize int) error
	FindOrderInfo(ctx context.Context, id uint32) (*OrderInfo, error)
	UpdateOrderInfo(ctx context.Context, orderInfo *OrderInfo) error
	DeleteOrderInfo(ctx context.Context, id uint32) error
	FindByOrderId(ctx context.Context, orderId string) (*OrderInfo, error)
	UpdateByOrderId(ctx context.Context, orderInfo *OrderInfo) error
	DeleteByOrderId(ctx context.Context, orderId string) error
}

type OrderInfoImpl struct {
	dao *Dao
}

func NewOrderInfoDao(dao *Dao) OrderInfoDao {
	return &OrderInfoImpl{
		dao: dao,
	}
}

func (d *OrderInfoImpl) CreateOrderInfo(ctx context.Context, orderInfo *OrderInfo) error {
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "INSERT INTO `order_info` (`id`,`order_id`,`status`,`created_at`,`updated_at`) VALUES (:id,:order_id,:status,:created_at,:updated_at)", orderInfo)
	return err
}

func (d *OrderInfoImpl) BatchCreateOrderInfo(ctx context.Context, list []*OrderInfo, batchSize int) error {
	return d.dao.InTx(ctx, func(ctx context.Context) error {
		for i := 0; i < len(list); i += batchSize {
			ends := i + batchSize
			if ends > len(list) {
				ends = len(list)
			}
			_, err := d.dao.DB(ctx).NamedExecContext(ctx, "INSERT INTO `order_info` (`id`,`order_id`,`status`,`created_at`,`updated_at`) VALUES (:id,:order_id,:status,:created_at,:updated_at)", list[i:ends])
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *OrderInfoImpl) FindOrderInfo(ctx context.Context, id uint32) (*OrderInfo, error) {
	var orderInfo OrderInfo
	if err := d.dao.DB(ctx).QueryRowxContext(ctx, "SELECT `id`,`order_id`,`status`,`created_at`,`updated_at` FROM `order_info` WHERE `id` = ?", id).StructScan(&orderInfo); err != nil {
		return nil, err
	}
	return &orderInfo, nil
}

func (d *OrderInfoImpl) UpdateOrderInfo(ctx context.Context, orderInfo *OrderInfo) error {
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "UPDATE `order_info` SET `id` = :id AND `order_id` = :order_id AND `status` = :status AND `created_at` = :created_at AND `updated_at` = :updated_at WHERE `id` = :id", orderInfo)
	return err
}

func (d *OrderInfoImpl) DeleteOrderInfo(ctx context.Context, id uint32) error {
	_, err := d.dao.DB(ctx).ExecContext(ctx, "DELETE FROM `order_info` WHERE `id` = ?", id)
	return err
}

func (d *OrderInfoImpl) FindByOrderId(ctx context.Context, orderId string) (*OrderInfo, error) {
	var orderInfo OrderInfo
	if err := d.dao.DB(ctx).QueryRowxContext(ctx, "SELECT `id`,`order_id`,`status`,`created_at`,`updated_at` FROM `order_info` WHERE `order_id` = ?", orderId).StructScan(&orderInfo); err != nil {
		return nil, err
	}
	return &orderInfo, nil
}

func (d *OrderInfoImpl) UpdateByOrderId(ctx context.Context, orderInfo *OrderInfo) error {
	_, err := d.dao.DB(ctx).NamedExecContext(ctx, "UPDATE `order_info` SET `id` = :id AND `order_id` = :order_id AND `status` = :status AND `created_at` = :created_at AND `updated_at` = :updated_at WHERE `order_id` = :order_id", orderInfo)
	return err
}

func (d *OrderInfoImpl) DeleteByOrderId(ctx context.Context, orderId string) error {
	_, err := d.dao.DB(ctx).ExecContext(ctx, "DELETE FROM `order_info` WHERE `order_id` = ?", orderId)
	return err
}
