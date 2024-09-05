CREATE TABLE "contacts"(
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "first_name" VARCHAR not null,
    "last_name" VARCHAR not null,
    "phone_number" VARCHAR not null,
    "street" VARCHAR not null,
    "created_at" TIMESTAMP not null default CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP not null,
    constraint "contacts_pkey" primary key("id") 
);

create unique index "contacts_phone_number_key" on "contacts"("phone_number"); 