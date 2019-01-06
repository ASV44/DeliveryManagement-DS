package db

const (
	InsertOrder = `INSERT INTO orders(order_id, awb_number, allow_open_parcel, created_date, 
											labels, latitude, longitude, service_payment, 
											receiver_address, receiver_address_locality, 
											receiver_contact, receiver_name, receiver_phone, 
											shipper_address, shipper_address_locality, 
											shipper_contact, shipper_name, shipper_phone, 
											status_group_id, today_important)
						VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	GetAllOrders   = "SELECT * FROM orders"
	GetOrderById   = "SELECT * FROM orders WHERE order_id = ? LIMIT 1 ALLOW FILTERING"
	GetOrdersByAwb = "SELECT * FROM orders WHERE awb_number = ? ALLOW FILTERING"
)
