CREATE TABLE
  orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
  );

INSERT INTO
  orders (id, price, tax, final_price)
VALUES
  ("1", 10.99, 0.5, 54.99)
  ("2", 9.99, 0.5, 49.99);