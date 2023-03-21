///1. 
//a. Insert 5 operators pada table operators.
USE alta_online_shop

INSERT INTO operators (name) VALUES
('Telkomsel'),
('Indosat'),
('XL'),
('Tri'),
('Smartfren')

//b.Insert 3 product type.
USE alta_online_shop
INSERT INTO product_type (name) VALUES
('Pulsa'),
('Paket Data'),
('Voucher Game')

//c. Insert 2 product dengan product type id = 1, dan operators id = 3. 
USE alta_online_shop

INSERT INTO product (id ,name ,product_type_id ,operator_id ,price, description_id ) VALUES
('1', 'Pulsa XL 10.000', '1', '3', '10000', '1' ),
('2', 'Pulsa XL 20.000', '2', '3', '20000', '1')

//d. Insert 3 product dengan product type id = 2, dan operators id = 1.
USE alta_online_shop

INSERT INTO product (id ,name ,product_type_id ,operator_id ,price, description_id ) VALUES
('4', 'Paket Data Telkomsel 1GB', '2', '1', '50000', '1' ),
('5', 'Paket Data Telkomsel 2GB', '2', '1', '60000', '1'),
('6', 'Paket Data Telkomsel 3GB', '2', '1', '80000', '1')

//e. Insert 3 product dengan product type id = 3, dan operators id = 4.
USE alta_online_shop

INSERT INTO product (id ,name ,product_type_id ,operator_id ,price, description_id ) VALUES
('7', 'Voucher Game Free Fire 50 Diamond', '3', '4', '50000', '1' ),
('8', 'Voucher Game Mobile Legends 100 Diamond', '3', '4', '100000', '1'),
('9', 'Voucher Game PUBG 300 UC', '3', '4', '40000', '1')
//f. Insert product description pada setiap product.
USE alta_online_shop

INSERT INTO product_description (product_id, description) VALUES
(1, 'Pulsa XL 10.000'),
(2, 'Pulsa XL 20.000'),
(3, 'Paket Data Telkomsel 1GB'),
(4, 'Paket Data Telkomsel 2GB'),
(5, 'Paket Data Telkomsel 5GB'),
(6, 'Voucher Game Free Fire 50 Diamond'),
(7, 'Voucher Game Mobile Legends 100 Diamond'),
(8, 'Voucher Game PUBG 300 UC');
//g.Insert 3 payment methods.
USE alta_online_shop
INSERT INTO payment_method (name) VALUES
('1','OVO'),
('2','GoPay'),
('3','Bank')
//h.  Insert 5 user pada tabel user.
USE alta_online_shop

INSERT INTO user (id, name, email, password) VALUES
('1', 'Aciw', 'aciw@gmail.com', '081234567890'),
('2', 'rensy', 'rensy@yahoo.com', '085678901234'),
('3', 'lala', 'lala@gmail.com', '081234567891'),
('4', 'devi', 'devi@lala.com', '081234567892'),
('5', 'lia', 'lia@haha.com', '085678901235')

//i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
USE alta_online_shop;

INSERT INTO transactions (id ,user_id , payment_method_id , total_price , transaction_date) VALUES
('1', '1', '1', '120000', '	1753-01-01'),
('2', '1', '2', '150000', '	1753-01-01'),
('3', '1', '3', '90000', '	1753-01-01'),
('4', '2', '1', '220000', '	1753-01-01'),
('5', '2', '2', '175000', '	1753-01-01'),
('6', '2', '3', '110000', '	1753-01-01'),
('7', '3', '1', '320000', '	1753-01-01'),
('8', '3', '2', '255000', '	1753-01-01'),
('9', '3', '3', '180000', '	1753-01-01'),
('10', '4', '1','420000', '	1753-01-01'),
('11', '4', '2','315000', '	1753-01-01'),
('12', '4', '3','200000', '	1753-01-01'),
('13', '5', '1','520000', '	1753-01-01'),
('14', '5', '2','405000', '	1753-01-01'),
('15', '5', '3','250000', '	1753-01-01');
//j. Insert 3 product di masing-masing transaksi.
USE alta_online_shop

INSERT INTO transaction_detail (id , transaction_id , product_id , quantity) VALUES
('1', '1', '1', '2'),
('2', '1', '2', '1'),
('3', '1', '3', '3'),
('4', '2', '2', '2'),
('5', '2', '4', '1'),
('6', '2', '5', '2'),
('7', '3', '3', '1'),
('8', '3', '6', '4'),
('9', '3', '7', '2')

///2. SELECT
///a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
USE alta_online_shop

SELECT name FROM user WHERE gender='L'
//b. Tampilkan product dengan id = 3.
USE alta_online_shop

SELECT * FROM product WHERE id=3;
/// c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
USE alta_online_shop

SELECT * FROM user WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';
///d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
USE alta_online_shop

SELECT name FROM user WHERE gender='P'
/// e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
USE alta_online_shop

SELECT * FROM user ORDER BY name ASC;
///f. Tampilkan 5 data pada data product
USE alta_online_shop

SELECT * FROM product LIMIT 5;

///3. UPDATE
///a.Ubah data product id 1 dengan nama ‘product dummy’.
USE alta_online_shop


UPDATE product SET name = 'product dummy' WHERE id = 1
//b.Update qty = 3 pada transaction detail dengan product id = 1.
USE alta_online_shop

UPDATE transaction_detail SET quantity = 3 WHERE product_id = 1;

///4.Delete
//a. Delete data pada tabel product dengan id = 1.
USE alta_online_shop

DELETE FROM product WHERE id = 1
///b. Delete pada pada tabel product dengan product type id 1.
USE alta_online_shop

DELETE FROM product WHERE product_type_id = 1

//Join, Union, Sub query, Function
//1. Gabungkan data transaksi dari user id 1 dan user id 2.
USE alta_online_shop

SELECT * FROM transactions  WHERE user_id = 1
UNION
SELECT * FROM transactions  WHERE user_id = 2
//2. Tampilkan jumlah harga transaksi user id 1.
USE alta_online_shop

SELECT SUM(total_price) FROM transactions  WHERE user_id = 1
//3. Tampilkan total transaksi dengan product type 2.
USE alta_online_shop

SELECT COUNT(*) FROM transaction_detail
JOIN product ON transaction_detail.product_id = product.id
WHERE product.product_type_id = 2
//4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
USE alta_online_shop

SELECT product.*, product_type.name AS product_type_name FROM product
JOIN product_type ON product.product_type_id = product_type.id
//5. Tampilkan semua field table transaction, field name table product dan field name table user.
//6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
USE alta_online_shop

DELIMITER $$
CREATE FUNCTION delete_transaction_detail(id INT)
RETURNS BOOLEAN
BEGIN
    DECLARE deleted_rows INT DEFAULT 0;
    DELETE FROM transaction_detail WHERE transaction_id = id;
    SET deleted_rows = ROW_COUNT();
    RETURN deleted_rows > 0;
END$$
DELIMITER ;


//7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
USE alta_online_shop

DELIMITER $$
CREATE FUNCTION update_total_qty(id INT)
RETURNS BOOLEAN
BEGIN 
    DECLARE updated_rows INT DEFAULT 0;
    UPDATE transaction t
    SET total_qty = (SELECT SUM(qty) FROM transaction_detail WHERE transaction_id = t.id)
    WHERE t.id = id;
    SET updated_rows = ROW_COUNT();
    RETURN updated_rows > 0;
END$$
DELIMITER ;

//8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

USE alta_online_shop

SELECT *
FROM product
WHERE id NOT IN (
    SELECT product_id
    FROM transaction_detail
)