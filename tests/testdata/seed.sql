CREATE SEQUENCE IF NOT EXISTS public.carts_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;



CREATE SEQUENCE IF NOT EXISTS public.orders_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;



CREATE SEQUENCE IF NOT EXISTS public.products_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;


CREATE TABLE IF NOT EXISTS public.products
(
    id bigint NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
                             name text COLLATE pg_catalog."default",
                             price numeric,
                             CONSTRAINT products_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.products
    OWNER to postgres;
-- Index: idx_products_deleted_at

-- DROP INDEX IF EXISTS public.idx_products_deleted_at;

CREATE INDEX IF NOT EXISTS idx_products_deleted_at
    ON public.products USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS public.carts
(
    id bigint NOT NULL DEFAULT nextval('carts_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
                             user_id bigint,
                             CONSTRAINT carts_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.carts
    OWNER to postgres;
-- Index: idx_carts_deleted_at

-- DROP INDEX IF EXISTS public.idx_carts_deleted_at;

CREATE INDEX IF NOT EXISTS idx_carts_deleted_at
    ON public.carts USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS public.orders
(
    id bigint NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
                             user_id bigint,
                             status text COLLATE pg_catalog."default",
                             CONSTRAINT orders_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.orders
    OWNER to postgres;
-- Index: idx_orders_deleted_at

-- DROP INDEX IF EXISTS public.idx_orders_deleted_at;

CREATE INDEX IF NOT EXISTS idx_orders_deleted_at
    ON public.orders USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS public.cart_items
(
    cart_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity bigint,
    CONSTRAINT cart_items_pkey PRIMARY KEY (cart_id, product_id),
    CONSTRAINT fk_carts_items FOREIGN KEY (cart_id)
    REFERENCES public.carts (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
    CONSTRAINT fk_products_cart_items FOREIGN KEY (product_id)
    REFERENCES public.products (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.cart_items
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public.order_items
(
    order_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity bigint,
    CONSTRAINT order_items_pkey PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_orders_items FOREIGN KEY (order_id)
    REFERENCES public.orders (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
    CONSTRAINT fk_products_order_items FOREIGN KEY (product_id)
    REFERENCES public.products (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.order_items
    OWNER to postgres;

ALTER SEQUENCE public.carts_id_seq
    OWNED BY public.carts.id;

ALTER SEQUENCE public.carts_id_seq
    OWNER TO postgres;

ALTER SEQUENCE public.orders_id_seq
    OWNED BY public.orders.id;

ALTER SEQUENCE public.orders_id_seq
    OWNER TO postgres;

ALTER SEQUENCE public.products_id_seq
    OWNED BY public.products.id;

ALTER SEQUENCE public.products_id_seq
    OWNER TO postgres;






INSERT INTO carts (id, user_id) VALUES
                                        (1, 1),
                                        (2, 2);

-- Создание товаров
INSERT INTO products (id, name, price) VALUES
                                                        (1, 'Laptop', 999.99),
                                                        (2, 'Mouse', 29.99);

-- Создание заказов
INSERT INTO orders (id, user_id, status, created_at) VALUES
                                                         (1, 1, 'pending', NOW()),
                                                         (2, 2, 'completed', NOW());

-- Связь заказов и продуктов (order_items или аналогичная таблица)
INSERT INTO order_items (order_id, product_id, quantity) VALUES
                                                                    (1, 1, 1),
                                                                    (1, 2, 2),
                                                                    (2, 2, 1);