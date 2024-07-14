--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7
-- Dumped by pg_dump version 16.0

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
-- Name: championship_predictions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.championship_predictions (
    id integer NOT NULL,
    team_name character varying(255) NOT NULL,
    week integer NOT NULL,
    probability double precision NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.championship_predictions OWNER TO postgres;

--
-- Name: championship_predictions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.championship_predictions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.championship_predictions_id_seq OWNER TO postgres;

--
-- Name: championship_predictions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.championship_predictions_id_seq OWNED BY public.championship_predictions.id;


--
-- Name: matches; Type: TABLE; Schema: public; Owner: football_user
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    home_team_id integer,
    away_team_id integer,
    home_goals integer,
    away_goals integer,
    week integer
);


ALTER TABLE public.matches OWNER TO football_user;

--
-- Name: matches_id_seq; Type: SEQUENCE; Schema: public; Owner: football_user
--

CREATE SEQUENCE public.matches_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.matches_id_seq OWNER TO football_user;

--
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: football_user
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
-- Name: settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.settings (
    id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    current_week integer,
    league_ended boolean
);


ALTER TABLE public.settings OWNER TO postgres;

--
-- Name: settings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.settings_id_seq OWNER TO postgres;

--
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


--
-- Name: teams; Type: TABLE; Schema: public; Owner: football_user
--

CREATE TABLE public.teams (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    power integer,
    points integer,
    wins integer,
    draws integer,
    losses integer,
    goals_for integer,
    goals_against integer,
    championship_probability numeric,
    played integer,
    goal_difference integer
);


ALTER TABLE public.teams OWNER TO football_user;

--
-- Name: teams_id_seq; Type: SEQUENCE; Schema: public; Owner: football_user
--

CREATE SEQUENCE public.teams_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.teams_id_seq OWNER TO football_user;

--
-- Name: teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: football_user
--

ALTER SEQUENCE public.teams_id_seq OWNED BY public.teams.id;


--
-- Name: championship_predictions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.championship_predictions ALTER COLUMN id SET DEFAULT nextval('public.championship_predictions_id_seq'::regclass);


--
-- Name: matches id; Type: DEFAULT; Schema: public; Owner: football_user
--

ALTER TABLE ONLY public.matches ALTER COLUMN id SET DEFAULT nextval('public.matches_id_seq'::regclass);


--
-- Name: settings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


--
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: football_user
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- Data for Name: championship_predictions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: football_user
--



--
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.settings VALUES
	(5, '2024-07-14 11:47:28.605243', '2024-07-14 11:47:28.605243', NULL, 0, false);


--
-- Data for Name: teams; Type: TABLE DATA; Schema: public; Owner: football_user
--

INSERT INTO public.teams VALUES
	(169, '2024-07-14 11:47:28.612967+03', '2024-07-14 11:47:28.617938+03', NULL, 'Chelsea', 92, 0, 0, 0, 0, 0, 0, 0, 0, 0),
	(170, '2024-07-14 11:47:28.615177+03', '2024-07-14 11:47:28.620819+03', NULL, 'Arsenal', 18, 0, 0, 0, 0, 0, 0, 0, 0, 0),
	(171, '2024-07-14 11:47:28.615774+03', '2024-07-14 11:47:28.623166+03', NULL, 'Manchester City', 32, 0, 0, 0, 0, 0, 0, 0, 0, 0),
	(172, '2024-07-14 11:47:28.616393+03', '2024-07-14 11:47:28.624613+03', NULL, 'Liverpool', 6, 0, 0, 0, 0, 0, 0, 0, 0, 0);


--
-- Name: championship_predictions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.championship_predictions_id_seq', 292, true);


--
-- Name: matches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: football_user
--

SELECT pg_catalog.setval('public.matches_id_seq', 1526, true);


--
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.settings_id_seq', 5, true);


--
-- Name: teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: football_user
--

SELECT pg_catalog.setval('public.teams_id_seq', 172, true);


--
-- Name: championship_predictions championship_predictions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.championship_predictions
    ADD CONSTRAINT championship_predictions_pkey PRIMARY KEY (id);


--
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: football_user
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: football_user
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (id);


--
-- Name: idx_matches_deleted_at; Type: INDEX; Schema: public; Owner: football_user
--

CREATE INDEX idx_matches_deleted_at ON public.matches USING btree (deleted_at);


--
-- Name: idx_settings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_settings_deleted_at ON public.settings USING btree (deleted_at);


--
-- Name: idx_teams_deleted_at; Type: INDEX; Schema: public; Owner: football_user
--

CREATE INDEX idx_teams_deleted_at ON public.teams USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

