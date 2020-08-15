create table article
(
  id SERIAL primary key,
  username VARCHAR(50) not null,
  text VARCHAR(255) not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);