CREATE TABLE IF NOT EXISTS users(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "password" VARCHAR (255) NOT NULL,
    "email" VARCHAR (255) UNIQUE NOT NULL,
    "user_role_id" INTEGER REFERENCES user_roles (id),
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
ALTER TABLE users ADD CONSTRAINT fk_users_user_roles FOREIGN KEY (user_role_id) REFERENCES user_roles (id);

CREATE TABLE IF NOT EXISTS hotels(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "user_id" INTEGER REFERENCES users (id),
    "name" VARCHAR (255) NOT NULL,
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
ALTER TABLE hotels ADD CONSTRAINT fk_hotels_users FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS hotel_types(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "hotel_id" INTEGER REFERENCES hotels (id),
    "name" VARCHAR(255) NOT NULL,
    "price" INTEGER NOT NULL,
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
ALTER TABLE hotel_types ADD CONSTRAINT fk_hotel_types_hotels FOREIGN KEY (hotel_id) REFERENCES hotels (id);

CREATE TABLE IF NOT EXISTS hotel_rooms(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "hotel_id" INTEGER REFERENCES hotels (id),
    "hotel_types_id" INTEGER REFERENCES hotel_types (id),
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
ALTER TABLE hotel_rooms ADD CONSTRAINT fk_hotel_rooms_hotels FOREIGN KEY (hotel_id) REFERENCES hotels (id);
ALTER TABLE hotel_rooms ADD CONSTRAINT fk_hotel_rooms_hotel_types FOREIGN KEY (hotel_types_id) REFERENCES hotel_types (id);

CREATE TABLE IF NOT EXISTS transactions(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "user_id" INTEGER REFERENCES users (id),
    "is_paid" BOOLEAN DEFAULT false,
    "hotel_room_id" INTEGER REFERENCES hotel_rooms (id),
    "paid_price" INTEGER NOT NULL,
    "paid_at" TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NULL,
    "checkin_date" DATE NOT NULL,
    "checkout_date" DATE NOT NULL,
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
ALTER TABLE transactions ADD CONSTRAINT fk_transactions_hotel_rooms FOREIGN KEY (hotel_room_id) REFERENCES hotel_rooms (id);
ALTER TABLE transactions ADD CONSTRAINT fk_transactions_users FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS user_roles(
    "id" BIGSERIAL PRIMARY KEY NOT NULl,
    "name" VARCHAR(255) NOT NULL,
    "created_at" timestamp(0) without time zone null, 
    "updated_at" timestamp(0) without time zone null
);
