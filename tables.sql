CREATE TABLE IF NOT EXISTS users(
  id bigserial primary key,
  name text not null,
  email text not null,
  created_at timestamp(0) with time zone not null default now()
);
