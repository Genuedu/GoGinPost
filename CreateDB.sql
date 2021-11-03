CREATE TABLE IF NOT EXISTS public.albums
(
    id uuid NOT NULL,
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    artist character varying(255) COLLATE pg_catalog."default" NOT NULL,
    price integer,
    CONSTRAINT albums_pkey PRIMARY KEY (id)
);