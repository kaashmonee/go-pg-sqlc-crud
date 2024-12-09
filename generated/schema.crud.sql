
  -- public.users CRUD Operations
  
  -- name: GetUsersByID :one
  SELECT id, created_at, updated_at, email, encrypted_password
  FROM public.users
  WHERE id = $1;
  
  -- name: ListUserss :many
  SELECT id, created_at, updated_at, email, encrypted_password
  FROM public.users
  ORDER BY created_at DESC
  LIMIT $1 OFFSET $2;
  
  -- name: CreateUsers :one
  INSERT INTO public.users (
	email,
	encrypted_password,
  )
  VALUES (
	$1,
	$2,
  )
  RETURNING id, created_at, updated_at, email, encrypted_password;
  
  -- name: UpdateUsers :one
  UPDATE public.users
  SET
	updated_at = $3,
	email = $4,
	encrypted_password = $5,
	updated_at = now()
  WHERE id = $1
  RETURNING id, created_at, updated_at, email, encrypted_password;
  
  -- name: DeleteUsers :exec
  DELETE FROM public.users
  WHERE id = $1;
  


  -- public.todos CRUD Operations
  
  -- name: GetTodosByID :one
  SELECT id, created_at, updated_at, title, description, is_done, user_id
  FROM public.todos
  WHERE id = $1;
  
  -- name: ListTodoss :many
  SELECT id, created_at, updated_at, title, description, is_done, user_id
  FROM public.todos
  ORDER BY created_at DESC
  LIMIT $1 OFFSET $2;
  
  -- name: CreateTodos :one
  INSERT INTO public.todos (
	title,
	description,
	is_done,
	user_id,
  )
  VALUES (
	$1,
	$2,
	$3,
	$4,
  )
  RETURNING id, created_at, updated_at, title, description, is_done, user_id;
  
  -- name: UpdateTodos :one
  UPDATE public.todos
  SET
	updated_at = $3,
	title = $4,
	description = $5,
	is_done = $6,
	user_id = $7,
	updated_at = now()
  WHERE id = $1
  RETURNING id, created_at, updated_at, title, description, is_done, user_id;
  
  -- name: DeleteTodos :exec
  DELETE FROM public.todos
  WHERE id = $1;
  


  -- public.sessions CRUD Operations
  
  -- name: GetSessionsByID :one
  SELECT id, created_at, updated_at, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at
  FROM public.sessions
  WHERE id = $1;
  
  -- name: ListSessionss :many
  SELECT id, created_at, updated_at, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at
  FROM public.sessions
  ORDER BY created_at DESC
  LIMIT $1 OFFSET $2;
  
  -- name: CreateSessions :one
  INSERT INTO public.sessions (
	user_id,
	refresh_token,
	user_agent,
	client_ip,
	is_blocked,
	expires_at,
  )
  VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
  )
  RETURNING id, created_at, updated_at, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at;
  
  -- name: UpdateSessions :one
  UPDATE public.sessions
  SET
	updated_at = $3,
	user_id = $4,
	refresh_token = $5,
	user_agent = $6,
	client_ip = $7,
	is_blocked = $8,
	expires_at = $9,
	updated_at = now()
  WHERE id = $1
  RETURNING id, created_at, updated_at, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at;
  
  -- name: DeleteSessions :exec
  DELETE FROM public.sessions
  WHERE id = $1;
  