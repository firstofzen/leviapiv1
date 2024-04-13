create keyspace user with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
create keyspace product with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
create keyspace trade with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
create type product.rating_user(
                                   email_user text,
                                   feedback   text,
                                   score      float
                               );
create type trade.orders_status(
                                   last_update timestamp,
                                   status      text
                               );
create type trade.product_identify(
                                      category_2 text,
                                      id         uuid
                                  );

create table user.account
(
    email      text,
    name       text,
    avatar_url text,
    primary key ( email )
);
create table user.detail
(
    email      text,
    phone      text,
    address    text,
    is_violate boolean,
    primary key ( email )
);
create table trade.orders
(
    email            text,
    id               uuid,
    product_identify frozen<trade.product_identify>,
    create_at        timestamp,
    total_amount     float,
    status           trade.orders_status,
    primary key ( email, create_at )
) with clustering order by (create_at desc);

create table trade.booth
(
    email                   text,
    brand                   text,
    is_authentic            boolean,
    description             text,
    order_id_wait_confirm   set<text>,
    order_id_delivering     set<text>,
    product_identifies      set<frozen<trade.product_identify>>,
    product_sold_identifies set<frozen<trade.product_identify>>,
    primary key ( email )
);

create table trade.cart
(
    email     text,
    order_ids set<text>,
    primary key ( email)
);

create table product.product
(
    category_1   text,
    category_2   text,
    id           uuid,
    email        text,
    avatar_url   text,
    pictures_url set<text>,
    title        text,
    description  text,
    price        float,
    total_rating float,
    sale         float,
    sold         int,
    inventory    int,
    publish_at   timestamp,
    primary key ( category_2, id )
) with caching = {'enabled': 'true'};
create table product.rating
(
    product_category_2 text,
    product_id         uuid,
    rating_users       set<frozen<product.rating_user>>,
    count_rating       int,
    primary key ((product_category_2, product_id))
);
create table product.ranking
(
    product_category_2 text,
    product_id         uuid,
    sold               int,
    total_rating       float,
    primary key ((product_category_2, product_id), sold, total_rating )
); 

