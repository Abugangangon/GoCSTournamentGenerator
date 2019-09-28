CREATE TABLE public."Tournament"
(
    "Id" integer NOT NULL DEFAULT nextval('"Tournament_Id_seq"'::regclass),
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Tournament_pkey" PRIMARY KEY ("Id", "Name"),
    CONSTRAINT "Unique_Tournament_Id" UNIQUE ("Id")

)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Tournament"
    OWNER to postgres;

CREATE TABLE public."Team"
(
    "Id" integer NOT NULL DEFAULT nextval('"Team_Id_seq"'::regclass),
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Team_pkey" PRIMARY KEY ("Id", "Name"),
    CONSTRAINT "Unique_Team_Id" UNIQUE ("Id")
,
    CONSTRAINT "Unique_Team_Name" UNIQUE ("Name")

)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Team"
    OWNER to postgres;

CREATE TABLE public."Group"
(
    "Id" integer NOT NULL DEFAULT nextval('"Group_Id_seq"'::regclass),
    "Number" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( CYCLE INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 16 CACHE 1 ),
    "Tournament_Id" integer NOT NULL DEFAULT nextval('"Group_Tournament_Id_seq"'::regclass),
    CONSTRAINT "Group_pkey" PRIMARY KEY ("Id"),
    CONSTRAINT "Group_Tournament_Id" FOREIGN KEY ("Tournament_Id")
        REFERENCES public."Tournament" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Group"
    OWNER to postgres;

CREATE TABLE public."Score"
(
    "Id" integer NOT NULL DEFAULT nextval('"Score_Id_seq"'::regclass),
    "Team_Id" integer NOT NULL DEFAULT nextval('"Score_Team_Id_seq"'::regclass),
    "Group_Id" integer NOT NULL DEFAULT nextval('"Score_Group_Id_seq"'::regclass),
    "Points" integer DEFAULT 0,
    CONSTRAINT "Score_pkey" PRIMARY KEY ("Id"),
    CONSTRAINT "Score_Group_Id" FOREIGN KEY ("Group_Id")
        REFERENCES public."Group" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Score_Team_Id" FOREIGN KEY ("Team_Id")
        REFERENCES public."Team" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Score"
    OWNER to postgres;

CREATE TABLE public."Game"
(
    "Id" integer NOT NULL DEFAULT nextval('"Game_Id_seq"'::regclass),
    "Tournament_Id" integer NOT NULL DEFAULT nextval('"Game_Tournament_Id_seq"'::regclass),
    "Is_Group" boolean NOT NULL,
    "Team1_Id" integer NOT NULL DEFAULT nextval('"Game_Team1_Id_seq"'::regclass),
    "Team1_Score" integer NOT NULL,
    "Team2_Id" integer NOT NULL DEFAULT nextval('"Game_Team2_Id_seq"'::regclass),
    "Team2_Score" integer NOT NULL,
    CONSTRAINT "Game_pkey" PRIMARY KEY ("Id"),
    CONSTRAINT "Team1_Id" FOREIGN KEY ("Team1_Id")
        REFERENCES public."Team" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Team2_Id" FOREIGN KEY ("Team2_Id")
        REFERENCES public."Team" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Tournament_Id" FOREIGN KEY ("Tournament_Id")
        REFERENCES public."Tournament" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Game"
    OWNER to postgres;

CREATE TABLE public."Playoff"
(
    "Id" integer NOT NULL DEFAULT nextval('"Playoff_Id_seq"'::regclass),
    "Tournament_Id" integer NOT NULL DEFAULT nextval('"Playoff_Tournament_Id_seq"'::regclass),
    "Team1_Id" integer NOT NULL DEFAULT nextval('"Playoff_Team1_Id_seq"'::regclass),
    "Team2_Id" integer NOT NULL DEFAULT nextval('"Playoff_Team2_Id_seq"'::regclass),
    "Type" text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Playoff_pkey" PRIMARY KEY ("Id"),
    CONSTRAINT "Team1_Id" FOREIGN KEY ("Team1_Id")
        REFERENCES public."Team" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Team2_Id" FOREIGN KEY ("Team2_Id")
        REFERENCES public."Team" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT "Tournament_Id" FOREIGN KEY ("Tournament_Id")
        REFERENCES public."Tournament" ("Id") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public."Playoff"
    OWNER to postgres;

