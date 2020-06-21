CREATE TABLE IF NOT EXISTS articles (
  id bigint NOT NULL,
  title character varying NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (id)
);

CREATE SEQUENCE articles_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;