-- products
CREATE TABLE IF NOT EXISTS products (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  type varchar(100),
  flavor varchar(100),
  size varchar(50),
  price integer DEFAULT 0,
  quantity integer DEFAULT 0,
  created_at timestamptz DEFAULT now()
);

-- customers
CREATE TABLE IF NOT EXISTS customers (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  points integer DEFAULT 0,
  created_at timestamptz DEFAULT now()
);

-- transactions
CREATE TABLE IF NOT EXISTS transactions (
  id varchar(36) PRIMARY KEY,
  customer_id integer REFERENCES customers(id) ON DELETE SET NULL,
  product_id integer REFERENCES products(id) ON DELETE SET NULL,
  size varchar(50),
  flavor varchar(100),
  quantity integer DEFAULT 1,
  created_at timestamptz DEFAULT now()
);
