CREATE TABLE investor (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE car (
    id UUID PRIMARY KEY, 
    state_number VARCHAR NOT NULL,
    model  VARCHAR NOT NULL,
    status VARCHAR DEFAULT 'in-stock',
    price NUMERIC NOT NULL,
    daily_limit INT NOT NULL,
    over_limit INT NOT NULL,
    investor_percentage NUMERIC NOT NULL,
    investor_id UUID NOT NULL REFERENCES investor(id),
    km INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE client (
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE "order" (
    id UUID PRIMARY KEY,
    car_id UUID NOT NULL REFERENCES car(id),
    client_id UUID NOT NULL REFERENCES client(id),
    total_price NUMERIC NOT NULL,
    paid_price NUMERIC DEFAULT 0,
    day_count INT NOT NULL,
    give_km NUMERIC,
    receive_km NUMERIC,
    status VARCHAR NOT NULL DEFAULT 'new',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION order_investor_trigger() RETURNS trigger
LANGUAGE PLPGSQL
AS
$$
    BEGIN

        IF new.paid_price != 0
            THEN
                UPDATE foyda
	            SET
			sum = paid_price * investor_percentage / 100
	        from car join "order" on "order".car_id = car.id
	        where car.investor_id = foyda.investor_id;
        END IF;

        return new;
    END;
$$;

CREATE TRIGGER order_trigger_invest
AFTER INSERT OR UPDATE ON "order"
FOR EACH ROW EXECUTE PROCEDURE order_investor_trigger();

CREATE TABLE foyda (
    investor_id UUID NOT NULL REFERENCES investor(id),
    investor_first_name VARCHAR NOT NULL,
    sum NUMERIC DEFAULT 0
);


CREATE OR REPLACE FUNCTION order_status_trigger() RETURNS trigger
LANGUAGE PLPGSQL
AS
$$
    BEGIN

        IF new.status = 'new'
            THEN
                UPDATE car SET status = 'booked' WHERE id = new.car_id;
        ELSIF 
            old.status = 'new' AND 
            new.status = 'client_took'
                THEN
                    UPDATE car SET status = 'in_use' WHERE id = new.car_id;
        ELSIF 
            old.status = 'client_took' AND 
            new.status = 'client_returned'
                THEN
                    UPDATE car SET status = 'in_stock' WHERE id = new.car_id;
        END IF;

        return new;
    END;
$$;

CREATE TRIGGER order_trigger
AFTER INSERT OR UPDATE ON "order"
FOR EACH ROW EXECUTE PROCEDURE order_status_trigger();






