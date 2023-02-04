--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: online_shop; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA online_shop;


ALTER SCHEMA online_shop OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: deliveries; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.deliveries (
    id integer NOT NULL,
    order_id integer NOT NULL,
    address_id integer NOT NULL,
    target_date timestamp without time zone,
    type_id integer
);


ALTER TABLE online_shop.deliveries OWNER TO postgres;

--
-- Name: deliveries_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.deliveries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.deliveries_id_seq OWNER TO postgres;

--
-- Name: deliveries_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.deliveries_id_seq OWNED BY online_shop.deliveries.id;


--
-- Name: items; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.items (
    id integer NOT NULL,
    item_name character varying,
    price numeric(2,0) DEFAULT 1.00,
    description character varying,
    image_ids integer[],
    CONSTRAINT items_price CHECK ((price > 0.00))
);


ALTER TABLE online_shop.items OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.items_id_seq OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.items_id_seq OWNED BY online_shop.items.id;


--
-- Name: order_items; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.order_items (
    id integer NOT NULL,
    order_id integer NOT NULL,
    item_id integer NOT NULL,
    quantity integer,
    default_price numeric(2,0) DEFAULT 1.00,
    discount integer DEFAULT 0,
    final_price numeric(2,0) GENERATED ALWAYS AS (((default_price * ((100 - discount))::numeric) / (100)::numeric)) STORED,
    CONSTRAINT order_items_default_price CHECK ((default_price > 0.00)),
    CONSTRAINT order_items_final_price CHECK ((final_price > 0.00))
);


ALTER TABLE online_shop.order_items OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.order_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.order_items_id_seq OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.order_items_id_seq OWNED BY online_shop.order_items.id;


--
-- Name: order_statuses; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.order_statuses (
    id integer NOT NULL,
    status_name character varying
);


ALTER TABLE online_shop.order_statuses OWNER TO postgres;

--
-- Name: order_statuses_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.order_statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.order_statuses_id_seq OWNER TO postgres;

--
-- Name: order_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.order_statuses_id_seq OWNED BY online_shop.order_statuses.id;


--
-- Name: orders; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.orders (
    id integer NOT NULL,
    display_number character varying NOT NULL,
    user_id integer NOT NULL,
    status_id integer DEFAULT 1,
    cancel_reason character varying,
    payment_id integer NOT NULL,
    total_price numeric(2,0) DEFAULT 1.00,
    creation_date timestamp without time zone DEFAULT now(),
    modification_date timestamp without time zone DEFAULT now()
);


ALTER TABLE online_shop.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.orders_id_seq OWNED BY online_shop.orders.id;


--
-- Name: users; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.users (
    id integer NOT NULL,
    user_name character varying,
    user_surname character varying,
    user_patronymic character varying,
    phone character varying NOT NULL,
    birthdate timestamp without time zone,
    password_hash character varying,
    mail character varying,
    role_id integer,
    token character varying
);


ALTER TABLE online_shop.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.users_id_seq OWNED BY online_shop.users.id;


--
-- Name: deliveries id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.deliveries ALTER COLUMN id SET DEFAULT nextval('online_shop.deliveries_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.items ALTER COLUMN id SET DEFAULT nextval('online_shop.items_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_items ALTER COLUMN id SET DEFAULT nextval('online_shop.order_items_id_seq'::regclass);


--
-- Name: order_statuses id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_statuses ALTER COLUMN id SET DEFAULT nextval('online_shop.order_statuses_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.orders ALTER COLUMN id SET DEFAULT nextval('online_shop.orders_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.users ALTER COLUMN id SET DEFAULT nextval('online_shop.users_id_seq'::regclass);


--
-- Data for Name: deliveries; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.deliveries (id, order_id, address_id, target_date, type_id) FROM stdin;
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.items (id, item_name, price, description, image_ids) FROM stdin;
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.order_items (id, order_id, item_id, quantity, default_price, discount) FROM stdin;
\.


--
-- Data for Name: order_statuses; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.order_statuses (id, status_name) FROM stdin;
1	Новый
2	Выдан
3	Отменен
4	Формируется
5	Готов к выдаче
6	Передан курьеру
7	Оплачен, готов к выдаче
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.orders (id, display_number, user_id, status_id, cancel_reason, payment_id, total_price, creation_date, modification_date) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.users (id, user_name, user_surname, user_patronymic, phone, birthdate, password_hash, mail, role_id, token) FROM stdin;
1	admin	admin	admin	89000000000	\N	nhabgnkasgnbkiasg	admin@mail.ru	1	dfasdnmgfasdngfadsjkfg
\.


--
-- Name: deliveries_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.deliveries_id_seq', 1, false);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.items_id_seq', 1, false);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.order_items_id_seq', 1, false);


--
-- Name: order_statuses_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.order_statuses_id_seq', 7, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.orders_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.users_id_seq', 1, true);


--
-- Name: deliveries deliveries_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.deliveries
    ADD CONSTRAINT deliveries_pkey PRIMARY KEY (id);


--
-- Name: items items_name; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.items
    ADD CONSTRAINT items_name UNIQUE (item_name);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: order_statuses order_statuses_name; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_statuses
    ADD CONSTRAINT order_statuses_name UNIQUE (status_name);


--
-- Name: order_statuses order_statuses_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_statuses
    ADD CONSTRAINT order_statuses_pkey PRIMARY KEY (id);


--
-- Name: orders orders_display_num; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.orders
    ADD CONSTRAINT orders_display_num UNIQUE (display_number);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: users users_mail; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.users
    ADD CONSTRAINT users_mail UNIQUE (mail);


--
-- Name: users users_phone; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.users
    ADD CONSTRAINT users_phone UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: deliveries deliveries_order_fkey; Type: FK CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.deliveries
    ADD CONSTRAINT deliveries_order_fkey FOREIGN KEY (order_id) REFERENCES online_shop.orders(id);


--
-- Name: order_items order_items_item_fkey; Type: FK CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_items
    ADD CONSTRAINT order_items_item_fkey FOREIGN KEY (item_id) REFERENCES online_shop.items(id);


--
-- Name: order_items order_items_order_fkey; Type: FK CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.order_items
    ADD CONSTRAINT order_items_order_fkey FOREIGN KEY (order_id) REFERENCES online_shop.orders(id);


--
-- Name: orders orders_status_fk; Type: FK CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.orders
    ADD CONSTRAINT orders_status_fk FOREIGN KEY (status_id) REFERENCES online_shop.order_statuses(id);


--
-- Name: orders orders_user_fk; Type: FK CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.orders
    ADD CONSTRAINT orders_user_fk FOREIGN KEY (user_id) REFERENCES online_shop.users(id);


--
-- PostgreSQL database dump complete
--

