KEYSPACE="delivery_management"
REPLICATION_FACTOR=3
TABLE="orders"
CQL="CREATE KEYSPACE IF NOT EXISTS $KEYSPACE WITH replication = {'class' : 'SimpleStrategy','replication_factor' : $REPLICATION_FACTOR};
     CREATE TABLE IF NOT EXISTS $KEYSPACE.$TABLE (
        order_id text,
     		awb_number text,
     		allow_open_parcel boolean,
     		created_date timestamp,
     		labels list<text>,
     		latitude float,
     		longitude float,
     		service_payment float,
     		receiver_address text,
     		receiver_address_locality text,
     		receiver_contact text,
     		receiver_name text,
     		receiver_phone text,
     		shipper_address text,
     		shipper_address_locality text,
     		shipper_contact text,
     		shipper_name text,
     		shipper_phone text,
     		status_group_id int,
     		today_important boolean,
     		PRIMARY KEY ((order_id, awb_number), created_date));"

until echo $CQL | cqlsh
 do
   echo "cqlsh: Cassandra is unavailable to initialize - will retry later"
   sleep 2
 done &

exec /docker-entrypoint.sh "$@"
