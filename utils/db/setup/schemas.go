package dbSetup

const initialSchema string = `
    create table if not exists stacks
    (
        id           varchar not null
            constraint stacks_pk
                primary key,
        name         varchar not null,
        type         varchar,
        network_name varchar,
        url          varchar
    );

    alter table stacks
        owner to whm;

    create unique index if not exists stacks_id_uindex
        on stacks (id);

    create unique index if not exists stacks_name_uindex
        on stacks (name);

    create table if not exists containers
    (
        stack_id varchar
            constraint containers_stacks_id_fk
                references stacks
                on delete cascade,
        id       varchar not null
            constraint container_pk
                primary key
    );

    alter table containers
        owner to whm;

    create unique index if not exists containers_id_uindex
        on containers (id);

    create table if not exists roles
    (
        role varchar not null
            constraint roles_pk
                primary key,
        name varchar not null
    );

    alter table roles
        owner to whm;

    create table if not exists users
    (
        name  varchar not null,
        id    varchar not null
            constraint users_pk
                primary key,
        email varchar not null,
        role  varchar default 'admin'::character varying
            constraint users_roles_role_fk
                references roles
    );

    alter table users
        owner to whm;

    create table if not exists domains
    (
        name varchar not null
            constraint domains_pk
                primary key
    );

    alter table domains
        owner to whm;

    insert into roles (role, name) values ('admin', 'Administrator') ON CONFLICT DO NOTHING
`
