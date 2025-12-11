INSERT INTO users (
    user_id,
    name,
    email,
    password_hash,
    role,
    created_at,
    updated_at
  )
VALUES (
    user_id:integer,
    'name:character varying',
    'email:character varying',
    'password_hash:character varying',
    'role:character varying',
    'created_at:timestamp without time zone',
    'updated_at:timestamp without time zone'
  );