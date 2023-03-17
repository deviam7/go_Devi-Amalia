//1. Create database alta_online_shop.
CREATE DATABASE alta_online_shop;
//Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
///a. Create Table User
CREATE TABLE alta_online_shop.user (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)
///b. Create table product, product type, operators, product description, payment_method.
CREATE TABLE alta_online_shop.product (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  price DECIMAL(10, 2),
  stock INT,
  operator_id INT,
  product_type_id INT,
  product_description_id INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (operator_id) REFERENCES operators(id),
  FOREIGN KEY (product_type_id) REFERENCES product_type(id),
  FOREIGN KEY (product_description_id) REFERENCES product_description(id)
)

CREATE TABLE alta_online_shop.product_type (
  id INT PRIMARY KEY,
  name VARCHAR(255)
)

CREATE TABLE alta_online_shop.operators (
  id INT PRIMARY KEY,
  name VARCHAR(255)
)

CREATE TABLE alta_online_shop.product_description (
  id INT PRIMARY KEY,
  description TEXT
);

CREATE TABLE alta_online_shop.payment_method (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
)

///c. Create table transaction, transaction detail.

CREATE TABLE alta_online_shop.transaction (
  id INT PRIMARY KEY,
  user_id INT,
  payment_method_id INT,
  transaction_date TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);

CREATE TABLE alta_online_shop.transaction_detail (
  id INT PRIMARY KEY,
  transaction_id INT,
  product_id INT,
  quantity INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transaction(id),
  FOREIGN KEY (product_id) REFERENCES product(id)
);

//3. Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE alta_online_shop.kurir (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

//4. Tambahkan ongkos_dasar column di tabel kurir.
SELECT id, name, created_at, updated_at
FROM alta_online_shop.kurir
ALTER TABLE kurir ADD ongkos_dasar DECIMAL(10,2)

//5. Rename tabel kurir menjadi shipping.
ALTER TABLE alta_online_shop.kurir RENAME TO shipping;

///6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE alta_online_shop.shipping;

//7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
/// a. 1-to-1: payment method description.
CREATE TABLE alta_online_shop.payment_method_description (
  id INT PRIMARY KEY AUTO_INCREMENT,
  description TEXT,
  payment_method_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
)
///b. 1-to-many: user dengan alamat.
CREATE TABLE alta_online_shop.alamat (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT,
  nama_jalan VARCHAR(255),
  kota VARCHAR(255),
  provinsi VARCHAR(255),
  kode_pos VARCHAR(10),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user(id)
)
///c. 3. many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE alta_online_shop.user_payment_method_detail (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT,
  payment_method_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
)