INSERT INTO users (
  username, password, name
) VALUES (
  'super', 'super', 'Super Admin Satu'
) RETURNING id INTO supers (user_id);
