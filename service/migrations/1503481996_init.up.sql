CREATE TABLE public.auctions (
    "id" UUID,
    "week" text NOT NULL,
    "country" varchar(2) NOT NULL,
    "dc" varchar(2) NOT NULL,
    "ingredient" text NOT NULL,
    "duration" integer,
    "start_date" TIMESTAMP NOT NULL DEFAULT 'now()',
    "qty" float8 NOT NULL,
    "threshold" integer[],
    "max_price" float8 DEFAULT '0',
    PRIMARY KEY ("id")
);

CREATE TABLE public.bids (
    "id" UUID,
    "auction_id" UUID,
    "user_id" UUID NOT NULL,
    "threshold" integer NOT NULL,
    "value" float8 NOT NULL,
    "created" TIMESTAMP NOT NULL DEFAULT 'now()',
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_auctions" FOREIGN KEY ("auction_id") REFERENCES "public"."auctions"("id") ON DELETE CASCADE
);

CREATE INDEX threshold_index ON bids (threshold);
