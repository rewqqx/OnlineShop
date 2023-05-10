--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

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
    price integer DEFAULT 100,
    description character varying,
    image_ids integer[],
    tag_ids integer[],
    CONSTRAINT items_price CHECK ((price > 0))
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
    default_price integer DEFAULT 100,
    discount integer DEFAULT 0,
    final_price integer GENERATED ALWAYS AS (ceil((((default_price * (100 - discount)) / 100))::double precision)) STORED,
    CONSTRAINT order_items_default_price CHECK ((default_price > 0)),
    CONSTRAINT order_items_final_price CHECK ((final_price > 0))
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
    total_price integer DEFAULT 100,
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
-- Name: payments; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.payments (
    id integer NOT NULL,
    payment_value integer,
    type_id integer,
    status_id integer,
    creation_date timestamp without time zone
);


ALTER TABLE online_shop.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.payments_id_seq OWNED BY online_shop.payments.id;


--
-- Name: tags; Type: TABLE; Schema: online_shop; Owner: postgres
--

CREATE TABLE online_shop.tags (
    id integer NOT NULL,
    tag_name character varying,
    parent_id integer
);


ALTER TABLE online_shop.tags OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: online_shop; Owner: postgres
--

CREATE SEQUENCE online_shop.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE online_shop.tags_id_seq OWNER TO postgres;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: online_shop; Owner: postgres
--

ALTER SEQUENCE online_shop.tags_id_seq OWNED BY online_shop.tags.id;


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
    sex integer DEFAULT 0,
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
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id integer NOT NULL,
    payment_value integer,
    type_id integer,
    status_id integer,
    creation_date timestamp without time zone
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


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
-- Name: payments id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.payments ALTER COLUMN id SET DEFAULT nextval('online_shop.payments_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.tags ALTER COLUMN id SET DEFAULT nextval('online_shop.tags_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.users ALTER COLUMN id SET DEFAULT nextval('online_shop.users_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Data for Name: deliveries; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.deliveries (id, order_id, address_id, target_date, type_id) FROM stdin;
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.items (id, item_name, price, description, image_ids, tag_ids) FROM stdin;
10	Pencil	30	Blue Pencil	{16}	{9}
4	Sweeter	15	Usual One	{6,7,8}	{5}
5	Ford	30	Very Fast Car	{9,10}	{3}
1	Apple	1	Sweee Apple	{1,2}	{1,8}
8	Iphone	90	The Last Iphone	{13,14}	{2,6}
6	Book	50	Can Be Useful	{11}	{9}
7	Candy	60	Berry Candy	{12}	{1}
9	Pizza	40	Still Hot	{15}	{1}
2	Orange	1	Soar Orange	{3,4}	{1,8}
3	Computer	1	Powerful Computer	{5}	{2}
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
-- Data for Name: payments; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.payments (id, payment_value, type_id, status_id, creation_date) FROM stdin;
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.tags (id, tag_name, parent_id) FROM stdin;
6	Phones	2
7	House	4
8	Fruits	1
9	Education	-1
1	Food	-1
5	Clothes	-1
4	Buildings	-1
2	Electronics	-1
3	Cars	-1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: online_shop; Owner: postgres
--

COPY online_shop.users (id, user_name, user_surname, user_patronymic, phone, birthdate, sex, password_hash, mail, role_id, token) FROM stdin;
1	Bogdan	Madzhuga	Andreevich		\N	0	da3814786f99c0c3bb53b36bd85599398a37d8f8	madzhuga@mail.ru	1	680ee3efa31e13b750bcb34874b9e89390b8a5de5b633bc9e086a306cae54d33
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, payment_value, type_id, status_id, creation_date) FROM stdin;
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
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.payments_id_seq', 1, false);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.tags_id_seq', 9, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: online_shop; Owner: postgres
--

SELECT pg_catalog.setval('online_shop.users_id_seq', 1, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 1, false);


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
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: online_shop; Owner: postgres
--

ALTER TABLE ONLY online_shop.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


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
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


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

