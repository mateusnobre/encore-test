CREATE TABLE "lecturer" (
    "id" TEXT NOT NULL,
    "name" character varying NOT NULL,
    "created_at" timestamp default NOW(),
    "updated_at" timestamp default NOW(),
    CONSTRAINT "lecturer_pk" PRIMARY KEY ("id"))