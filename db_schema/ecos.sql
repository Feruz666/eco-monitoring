CREATE TABLE "air" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar,
  "date_of_sampling" timestamptz,
  "number_of_sample" varchar,
  "concentration" float4,
  "unit_of_measure" varchar,
  "location" varchar,
  "longitude" varchar,
  "latitude" varchar,
  "license_area" varchar,
  "num_of_license" varchar,
  "company" varchar,
  "method_of_determ" varchar,
  "laboratory" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "snow" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar,
  "date_of_sampling" timestamptz,
  "number_of_sample" varchar,
  "concentration" float4,
  "unit_of_measure" varchar,
  "location" varchar,
  "longitude" varchar,
  "latitude" varchar,
  "source_of_emission" varchar,
  "license_area" varchar,
  "num_of_license" varchar,
  "company" varchar,
  "method_of_determ" varchar,
  "laboratory" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "surface_water" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar,
  "date_of_sampling" timestamptz,
  "number_of_sample" varchar,
  "concentration" float4,
  "unit_of_measure" varchar,
  "location" varchar,
  "longitude" varchar,
  "latitude" varchar,
  "reservoir" varchar,
  "license_area" varchar,
  "num_of_license" varchar,
  "company" varchar,
  "method_of_determ" varchar,
  "laboratory" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "bottom_sediments" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar,
  "date_of_sampling" timestamptz,
  "number_of_sample" varchar,
  "type_of_sediments" varchar,
  "concentration" float4,
  "unit_of_measure" varchar,
  "location" varchar,
  "longitude" varchar,
  "latitude" varchar,
  "reservoir" varchar,
  "license_area" varchar,
  "num_of_license" varchar,
  "company" varchar,
  "method_of_determ" varchar,
  "laboratory" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "soil" (
  "id" bigserial PRIMARY KEY,
  "substanse_name" varchar,
  "date_of_sampling" timestamptz,
  "number_of_sample" varchar,
  "soil_subtype" varchar,
  "concentration" float4,
  "unit_of_measure" varchar,
  "location" varchar,
  "longitude" varchar,
  "latitude" varchar,
  "license_area" varchar,
  "num_of_license" varchar,
  "company" varchar,
  "method_of_determ" varchar,
  "laboratory" varchar,
  "created_at" timestamptz DEFAULT (now())
);
