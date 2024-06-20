CREATE TABLE countries(
    country_id serial PRIMARY KEY,
    country_name text NOT NULL UNIQUE
);

CREATE TABLE warehouses(
    warehouse_id serial PRIMARY KEY,
    country_id int NOT NULL REFERENCES countries(country_id) ON DELETE CASCADE,
    city_name text NOT NULL,
    UNIQUE(country_id, city_name)
);

CREATE TABLE car_models(
    car_model_id serial PRIMARY KEY,
    car_model_name text NOT NULL UNIQUE
);

CREATE TABLE auto_parts(
    auto_part_id serial PRIMARY KEY,
    car_model_id int NOT NULL REFERENCES car_models(car_model_id) ON DELETE CASCADE,
    auto_part_name text NOT NULL UNIQUE
);

CREATE TABLE auto_part_components(
    auto_part_component_id serial PRIMARY KEY,
    auto_part_id int REFERENCES auto_parts(auto_part_id) ON DELETE CASCADE,
    parent_auto_part_component_id int REFERENCES auto_part_components(auto_part_component_id) ON DELETE CASCADE,
    auto_part_component_name text NOT NULL,
    UNIQUE(parent_id, auto_part_component_name)
);

CREATE TABLE warehouse_positions(
    warehouse_id int NOT NULL REFERENCES warehouses(warehouse_id) ON DELETE CASCADE,
    auto_part_component_id int NOT NULL REFERENCES auto_part_components(auto_part_component_id) ON DELETE CASCADE,
    quantity int NOT NULL,
    CONSTRAINT warehouse_position_id PRIMARY KEY (warehouse_id, auto_part_component_id)
);
