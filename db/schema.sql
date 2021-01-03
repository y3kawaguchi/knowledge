CREATE TABLE IF NOT EXISTS articles (
  id bigint NOT NULL,
  author_id bigint NOT NULL,
  title character varying NOT NULL,
  content text NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  PRIMARY KEY (id)
);

CREATE SEQUENCE articles_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER TABLE ONLY articles ALTER COLUMN id SET DEFAULT nextval('articles_id_seq'::regclass);



CREATE TABLE IF NOT EXISTS users (
  id bigint NOT NULL,
  username character varying NOT NULL,
  email character varying NOT NULL,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  deleted_at timestamp with time zone,
  PRIMARY KEY (id)
);

CREATE SEQUENCE users_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

  ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);