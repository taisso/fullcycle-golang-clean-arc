CREATE TABLE orders (
  id varchar(255) NOT NULL,
  price float NOT NULL,
  tax float NOT NULL,
  final_price float NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO
  orders (id, price, tax, final_price)
VALUES
  ('1', 100.00, 10.00, 110.00),
  ('2', 75.50, 7.55, 83.05),
  ('3', 50.25, 5.03, 55.28);