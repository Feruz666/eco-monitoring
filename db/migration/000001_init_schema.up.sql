CREATE TABLE "air" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar NOT NULL,
  "date_of_sampling" varchar NOT NULL,
  "number_of_sample" varchar NOT NULL,
  "concentration" float4 NOT NULL,
  "unit_of_measure" varchar NOT NULL,
  "location" varchar NOT NULL,
  "longitude" varchar NOT NULL,
  "latitude" varchar NOT NULL,
  "license_area" varchar NOT NULL,
  "num_of_license" varchar NOT NULL,
  "company" varchar NOT NULL,
  "method_of_determ" varchar NOT NULL,
  "laboratory" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "snow" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar NOT NULL,
  "date_of_sampling" varchar NOT NULL,
  "number_of_sample" varchar NOT NULL,
  "concentration" float4 NOT NULL,
  "unit_of_measure" varchar NOT NULL,
  "location" varchar NOT NULL,
  "longitude" varchar NOT NULL,
  "latitude" varchar NOT NULL,
  "source_of_emission" varchar NOT NULL,
  "license_area" varchar NOT NULL,
  "num_of_license" varchar NOT NULL,
  "company" varchar NOT NULL,
  "method_of_determ" varchar NOT NULL,
  "laboratory" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "surface_water" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar NOT NULL,
  "date_of_sampling" varchar NOT NULL,
  "number_of_sample" varchar NOT NULL,
  "concentration" float4 NOT NULL,
  "unit_of_measure" varchar NOT NULL,
  "location" varchar NOT NULL,
  "longitude" varchar NOT NULL,
  "latitude" varchar NOT NULL,
  "reservoir" varchar NOT NULL,
  "license_area" varchar NOT NULL,
  "num_of_license" varchar NOT NULL,
  "company" varchar NOT NULL,
  "method_of_determ" varchar NOT NULL,
  "laboratory" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "bottom_sediments" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar NOT NULL,
  "date_of_sampling" varchar NOT NULL,
  "number_of_sample" varchar NOT NULL,
  "type_of_sediments" varchar NOT NULL,
  "concentration" float4 NOT NULL,
  "unit_of_measure" varchar NOT NULL,
  "location" varchar NOT NULL,
  "longitude" varchar NOT NULL,
  "latitude" varchar NOT NULL,
  "reservoir" varchar NOT NULL,
  "license_area" varchar NOT NULL,
  "num_of_license" varchar NOT NULL,
  "company" varchar NOT NULL,
  "method_of_determ" varchar NOT NULL,
  "laboratory" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "soil" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar NOT NULL,
  "date_of_sampling" varchar NOT NULL,
  "number_of_sample" varchar NOT NULL,
  "soil_subtype" varchar NOT NULL,
  "concentration" float4 NOT NULL,
  "unit_of_measure" varchar NOT NULL,
  "location" varchar NOT NULL,
  "longitude" varchar NOT NULL,
  "latitude" varchar NOT NULL,
  "license_area" varchar NOT NULL,
  "num_of_license" varchar NOT NULL,
  "company" varchar NOT NULL,
  "method_of_determ" varchar NOT NULL,
  "laboratory" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);
