-- DROP TABLE IF EXISTS todos;
-- DROP TABLE IF EXISTS expenses;
-- DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL
);

-- CREATE TABLE IF NOT EXISTS todos (
--   id SERIAL PRIMARY KEY,
--   userId INTEGER NOT NULL,
--   done BOOLEAN NOT NULL,
--   todo VARCHAR(250) NOT NULL,
--   attachment VARCHAR(250),
--   CONSTRAINT fk_user FOREIGN KEY(userId) REFERENCES users(id)
-- );

CREATE TABLE IF NOT EXISTS expenses (
  id SERIAL PRIMARY KEY,
  userId INTEGER NOT NULL,
  currency CHAR(3),
  amount REAL,
  note VARCHAR(250),
  date TIMESTAMP NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(userId) REFERENCES users(id)
);

insert into users(username, password) values('user1', '123456'), ('user2', '123456');

-- insert into todos(userId, done, todo) values(1, false, 'user1 todo1')

insert into expenses(userId, currency, amount, date) 
values
	(1, 'SGD', 100.10, current_timestamp),
	(1, 'SGD', 100.20, current_timestamp),
	(1, 'SGD', 100.30, current_timestamp),
  (2, 'SGD', 100.40, current_timestamp),
  (2, 'SGD', 100.50, current_timestamp);

-- SELECT 
-- 	userId, 
--     currency, 
--     ROUND(SUM(amount) * 100)/100 as total_amount
-- FROM expenses 
-- GROUP BY 
-- 	userId, currency;