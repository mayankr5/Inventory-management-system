--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Homebrew)
-- Dumped by pg_dump version 14.10 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: ims
--

CREATE TABLE public.categories (
    category_id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.categories OWNER TO ims;

--
-- Name: categories_category_id_seq; Type: SEQUENCE; Schema: public; Owner: ims
--

ALTER TABLE public.categories ALTER COLUMN category_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.categories_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: product; Type: TABLE; Schema: public; Owner: ims
--

CREATE TABLE public.product (
    product_id integer NOT NULL,
    name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price integer NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.product OWNER TO ims;

--
-- Name: product_product_id_seq; Type: SEQUENCE; Schema: public; Owner: ims
--

ALTER TABLE public.product ALTER COLUMN product_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.product_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: ims
--

COPY public.categories (category_id, name) FROM stdin;
5	Stationary
4	Snacks
6	Hardware
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: ims
--

COPY public.product (product_id, name, quantity, price, category_id) FROM stdin;
25	Chips	20	30	4
27	Pens	30	10	5
28	Marker	20	30	5
\.


--
-- Name: categories_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ims
--

SELECT pg_catalog.setval('public.categories_category_id_seq', 7, true);


--
-- Name: product_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ims
--

SELECT pg_catalog.setval('public.product_product_id_seq', 30, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: ims
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: ims
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (product_id);


--
-- Name: categories unique_category; Type: CONSTRAINT; Schema: public; Owner: ims
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT unique_category UNIQUE (name);


--
-- Name: product fk_category; Type: FK CONSTRAINT; Schema: public; Owner: ims
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(category_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

