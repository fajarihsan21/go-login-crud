-- Creating table schema
CREATE TABLE public.tb_user (
	user_id uuid NOT NULL,
	username varchar NULL,
	"password" varchar NULL,
	email varchar NULL,
	phone varchar NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT tb_user_pk PRIMARY KEY (user_id)
);

-- Insert dummy data
INSERT INTO public.tb_user (user_id,username,"password",email,phone,created_at,updated_at) VALUES
	 ('cb5c31a8-b731-4316-ab92-08def766e26f','user1','$2a$12$.gEan8Ztvcj7g8bybxyLPOAHuxfIt44vmtAxMWFzII5W7wPMCA7ua','','','2024-08-09 19:26:49.066174+07','2024-08-09 19:26:49.066174+07');
