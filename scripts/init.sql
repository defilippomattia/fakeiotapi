CREATE TABLE FARM
(
  id INT NOT NULL,
  name VARCHAR NOT NULL,
  location VARCHAR NOT NULL,
  date_established date NOT NULL,
  latitude FLOAT NOT NULL,
  longitude FLOAT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE EMPLOYEE
(
  id INT NOT NULL,
  name VARCHAR NOT NULL,
  surname VARCHAR NOT NULL,
  role VARCHAR NOT NULL,
  phone VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE FIELD
(
  id INT NOT NULL,
  type VARCHAR NOT NULL,
  m2 FLOAT NOT NULL,
  farm_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (farm_id) REFERENCES FARM(id)
);

CREATE TABLE farm2employee
(
  farm_id INT NOT NULL,
  employee_id INT NOT NULL,
  PRIMARY KEY (farm_id, employee_id),
  FOREIGN KEY (farm_id) REFERENCES FARM(id),
  FOREIGN KEY (employee_id) REFERENCES EMPLOYEE(id)
);

CREATE TABLE SENSOR
(
  id INT NOT NULL,
  manufacturer VARCHAR NOT NULL,
  type VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  field_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (field_id) REFERENCES FIELD(id)
);

CREATE TABLE TEMPERATURE
(
  timestamp timestamp NOT NULL,
  value FLOAT NOT NULL,
  sensor_id INT NOT NULL,
  PRIMARY KEY (sensor_id, timestamp),
  FOREIGN KEY (sensor_id) REFERENCES SENSOR(id)
);

CREATE TABLE HUMIDITY
(
  timestamp timestamp NOT NULL,
  value FLOAT NOT NULL,
  sensor_id INT NOT NULL,
  PRIMARY KEY (sensor_id, timestamp),
  FOREIGN KEY (sensor_id) REFERENCES SENSOR(id)
);

CREATE TABLE PRESSURE
(
  timestamp timestamp NOT NULL,
  value FLOAT NOT NULL,
  sensor_id INT NOT NULL,
  PRIMARY KEY (sensor_id, timestamp),
  FOREIGN KEY (sensor_id) REFERENCES SENSOR(id)
);


-- Inserting data into the FARM table
INSERT INTO FARM (id, name, location, date_established, latitude, longitude)
VALUES (1, 'Jonhson Familiy Farm', '123 Main St', '2022-01-15', 40.123456, -75.654321),
       (2, 'Green Pastures', '456 Elm St', '2020-05-10', 38.987654, -72.123456),
       (3, 'Fields and Fields', '789 Oak St', '2018-09-02', 42.345678, -71.234567);

-- Inserting data into the EMPLOYEE table
INSERT INTO EMPLOYEE (id, name, surname, role, phone, email)
VALUES (1, 'Jane', 'Doe', 'Manager', '555-1234', 'john.doe@example.com'),
       (2, 'John', 'Smith', 'Farmhand', '555-5678', 'jane.smith@example.com'),
       (3, 'Robert', 'Johnson', 'Manager', '555-9876', 'robert.johnson@example.com'),
       (4, 'Emily', 'Davis', 'Farmhand', '555-1111', 'emily.davis@example.com'),
       (5, 'Michael', 'Brown', 'Manager', '555-2222', 'michael.brown@example.com'),
       (6, 'Sarah', 'Wilson', 'Farmhand', '555-3333', 'sarah.wilson@example.com'),
       (7, 'Christopher', 'Taylor', 'Farmhand', '555-4444', 'christopher.taylor@example.com'),
       (8, 'Jessica', 'Thomas', 'Farmhand', '555-5555', 'jessica.thomas@example.com');


-- Inserting data into the FIELD table
INSERT INTO FIELD (id, type, m2, farm_id)
VALUES (1, 'Corn', 1000.0, 1),
       (2, 'Potato', 500.0, 1),
       (3, 'Beans', 200.0, 1),
       (4, 'Peas', 1500.0, 1),
       (5, 'Rice', 800.0, 2),
       (6, 'Corn', 1000.0, 2),
       (7, 'Wheat', 1200.0, 2),
       (8, 'Tomato', 800.0, 2),
       (9, 'Potato', 500.0, 3),
       (10, 'Olives', 1000.0, 3),
       (11, 'Tomato', 400.0, 3),
       (12, 'Rice', 800.0, 3),
       (13, 'Sunflower', 900.0, 3),
       (14, 'Cucumber', 600.0, 3);

-- Inserting data into the farm2employee table
INSERT INTO farm2employee (farm_id, employee_id)
VALUES (1, 1),
       (1, 2),
       (2, 3),
       (2, 4),
       (3, 5),
       (3, 6),
       (3, 7),
       (3, 8);

-- Inserting sensors for Field 1
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (1, 'SensorCo', 'Temperature', 'Temp Sensor 1', 1),
       (2, 'SensorCo', 'Humidity', 'Humidity Sensor 1', 1),
       (3, 'SensorCo', 'Pressure', 'Pressure Sensor 1', 1);

-- Inserting sensors for Field 2
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (4, 'SensorCo', 'Temperature', 'Temp Sensor 2', 2),
       (5, 'SmartSense', 'Humidity', 'Humidity Sensor 2', 2),
       (6, 'SmartSense', 'Pressure', 'Pressure Sensor 2', 2);

-- Inserting sensors for Field 3
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (7, 'SmartSense', 'Temperature', 'Temp Sensor 3', 3),
       (8, 'SensorCo', 'Humidity', 'Humidity Sensor 3', 3),
       (9, 'SmartSense', 'Pressure', 'Pressure Sensor 3', 3);

-- Inserting sensors for Field 4
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (10, 'FSNSen', 'Temperature', 'Temp Sensor 4', 4),
       (11, 'FSNSen', 'Humidity', 'Humidity Sensor 4', 4),
       (12, 'SensorCo', 'Pressure', 'Pressure Sensor 4', 4);

-- Inserting sensors for Field 5
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (13, 'FSNSen', 'Temperature', 'Temp Sensor 5', 5),
       (14, 'SensorCo', 'Humidity', 'Humidity Sensor 5', 5),
       (15, 'SmartSense', 'Pressure', 'Pressure Sensor 5', 5);

-- Inserting sensors for Field 6
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (16, 'SensorCo', 'Temperature', 'Temp Sensor 6', 6),
       (17, 'CoSenCo', 'Humidity', 'Humidity Sensor 6', 6),
       (18, 'CoSenCo', 'Pressure', 'Pressure Sensor 6', 6);

-- Inserting sensors for Field 7
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (19, 'SensorCo', 'Temperature', 'Temp Sensor 7', 7),
       (20, 'SensorCo', 'Humidity', 'Humidity Sensor 7', 7),
       (21, 'CoSenCo', 'Pressure', 'Pressure Sensor 7', 7);

-- Inserting sensors for Field 8
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (22, 'CoSenCo', 'Temperature', 'Temp Sensor 8', 8),
       (23, 'CoSenCo', 'Humidity', 'Humidity Sensor 8', 8),
       (24, 'CoSenCo', 'Pressure', 'Pressure Sensor 8', 8);

-- Inserting sensors for Field 9
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (25, 'SensorCo', 'Temperature', 'Temp Sensor 9', 9),
       (26, 'SensorCo', 'Humidity', 'Humidity Sensor 9', 9),
       (27, 'SensorCo', 'Pressure', 'Pressure Sensor 9', 9);

-- Inserting sensors for Field 10
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (28, 'SmartSense', 'Temperature', 'Temp Sensor 10', 10),
       (29, 'SensorCo', 'Humidity', 'Humidity Sensor 10', 10),
       (30, 'SmartSense', 'Pressure', 'Pressure Sensor 10', 10);

-- Inserting sensors for Field 11
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (31, 'SensorCo', 'Temperature', 'Temp Sensor 11', 11),
       (32, 'SmartSense', 'Humidity', 'Humidity Sensor 11', 11),
       (33, 'SensorCo', 'Pressure', 'Pressure Sensor 11', 11);

-- Inserting sensors for Field 12
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (34, 'SmartSense', 'Temperature', 'Temp Sensor 12', 12),
       (35, 'SensorCo', 'Humidity', 'Humidity Sensor 12', 12),
       (36, 'SensorCo', 'Pressure', 'Pressure Sensor 12', 12);

-- Inserting sensors for Field 13
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (37, 'SensorCo', 'Temperature', 'Temp Sensor 13', 13),
       (38, 'SensorCo', 'Humidity', 'Humidity Sensor 13', 13),
       (39, 'SensorCo', 'Pressure', 'Pressure Sensor 13', 13);

-- Inserting sensors for Field 14
INSERT INTO SENSOR (id, manufacturer, type, name, field_id)
VALUES (40, 'SmartSense', 'Temperature', 'Temp Sensor 14', 14),
       (41, 'SmartSense', 'Humidity', 'Humidity Sensor 14', 14),
       (42, 'SmartSense', 'Pressure', 'Pressure Sensor 14', 14);



-- Inserting sensor readings for Temperature Sensor 1
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 1),
       ('2023-07-01 08:15:00', 26.0, 1),
       ('2023-07-01 08:30:00', 26.2, 1),
       ('2023-07-01 08:45:00', 26.5, 1),
       ('2023-07-01 09:00:00', 26.8, 1);

-- Inserting sensor readings for Humidity Sensor 1
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 2),
       ('2023-07-01 08:15:00', 61.0, 2),
       ('2023-07-01 08:30:00', 61.5, 2),
       ('2023-07-01 08:45:00', 61.8, 2),
       ('2023-07-01 09:00:00', 62.1, 2);

-- Inserting sensor readings for Pressure Sensor 1
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 3),
       ('2023-07-01 08:15:00', 1012.9, 3),
       ('2023-07-01 08:30:00', 1012.7, 3),
       ('2023-07-01 08:45:00', 1012.5, 3),
       ('2023-07-01 09:00:00', 1012.3, 3);

-- Inserting sensor readings for Temperature Sensor 2
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 4),
       ('2023-07-01 08:15:00', 26.0, 4),
       ('2023-07-01 08:30:00', 26.2, 4),
       ('2023-07-01 08:45:00', 26.5, 4),
       ('2023-07-01 09:00:00', 26.8, 4);

-- Inserting sensor readings for Humidity Sensor 2
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 5),
       ('2023-07-01 08:15:00', 61.0, 5),
       ('2023-07-01 08:30:00', 61.5, 5),
       ('2023-07-01 08:45:00', 61.8, 5),
       ('2023-07-01 09:00:00', 62.1, 5);

-- Inserting sensor readings for Pressure Sensor 2
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 6),
       ('2023-07-01 08:15:00', 1012.9, 6),
       ('2023-07-01 08:30:00', 1012.7, 6),
       ('2023-07-01 08:45:00', 1012.5, 6),
       ('2023-07-01 09:00:00', 1012.3, 6);

-- Inserting sensor readings for Temperature Sensor 3
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 7),
       ('2023-07-01 08:15:00', 26.0, 7),
       ('2023-07-01 08:30:00', 26.2, 7),
       ('2023-07-01 08:45:00', 26.5, 7),
       ('2023-07-01 09:00:00', 26.8, 7);

-- Inserting sensor readings for Humidity Sensor 3
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 8),
       ('2023-07-01 08:15:00', 61.0, 8),
       ('2023-07-01 08:30:00', 61.5, 8),
       ('2023-07-01 08:45:00', 61.8, 8),
       ('2023-07-01 09:00:00', 62.1, 8);

-- Inserting sensor readings for Pressure Sensor 3
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 9),
       ('2023-07-01 08:15:00', 1012.9, 9),
       ('2023-07-01 08:30:00', 1012.7, 9),
       ('2023-07-01 08:45:00', 1012.5, 9),
       ('2023-07-01 09:00:00', 1012.3, 9);
    
-- Inserting sensor readings for Temperature Sensor 4
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 10),
       ('2023-07-01 08:15:00', 26.0, 10),
       ('2023-07-01 08:30:00', 26.2, 10),
       ('2023-07-01 08:45:00', 26.5, 10),
       ('2023-07-01 09:00:00', 26.8, 10);

-- Inserting sensor readings for Humidity Sensor 4
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 11),
       ('2023-07-01 08:15:00', 61.0, 11),
       ('2023-07-01 08:30:00', 61.5, 11),
       ('2023-07-01 08:45:00', 61.8, 11),
       ('2023-07-01 09:00:00', 62.1, 11);

-- Inserting sensor readings for Pressure Sensor 4
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 12),
       ('2023-07-01 08:15:00', 1012.9, 12),
       ('2023-07-01 08:30:00', 1012.7, 12),
       ('2023-07-01 08:45:00', 1012.5, 12),
       ('2023-07-01 09:00:00', 1012.3, 12);

-- Inserting sensor readings for Temperature Sensor 5
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 13),
       ('2023-07-01 08:15:00', 26.0, 13),
       ('2023-07-01 08:30:00', 26.2, 13),
       ('2023-07-01 08:45:00', 26.5, 13),
       ('2023-07-01 09:00:00', 26.8, 13);

-- Inserting sensor readings for Humidity Sensor 5
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 14),
       ('2023-07-01 08:15:00', 61.0, 14),
       ('2023-07-01 08:30:00', 61.5, 14),
       ('2023-07-01 08:45:00', 61.8, 14),
       ('2023-07-01 09:00:00', 62.1, 14);

-- Inserting sensor readings for Pressure Sensor 5
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 15),
       ('2023-07-01 08:15:00', 1012.9, 15),
       ('2023-07-01 08:30:00', 1012.7, 15),
       ('2023-07-01 08:45:00', 1012.5, 15),
       ('2023-07-01 09:00:00', 1012.3, 15);

-- Inserting sensor readings for Temperature Sensor 6
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 16),
       ('2023-07-01 08:15:00', 26.0, 16),
       ('2023-07-01 08:30:00', 26.2, 16),
       ('2023-07-01 08:45:00', 26.5, 16),
       ('2023-07-01 09:00:00', 26.8, 16);
    
-- Inserting sensor readings for Humidity Sensor 6
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 17),
       ('2023-07-01 08:15:00', 61.0, 17),
       ('2023-07-01 08:30:00', 61.5, 17),
       ('2023-07-01 08:45:00', 61.8, 17),
       ('2023-07-01 09:00:00', 62.1, 17);

-- Inserting sensor readings for Pressure Sensor 6
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 18),
       ('2023-07-01 08:15:00', 1012.9, 18),
       ('2023-07-01 08:30:00', 1012.7, 18),
       ('2023-07-01 08:45:00', 1012.5, 18),
       ('2023-07-01 09:00:00', 1012.3, 18);

-- Inserting sensor readings for Temperature Sensor 7
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 19),
       ('2023-07-01 08:15:00', 26.0, 19),
       ('2023-07-01 08:30:00', 26.2, 19),
       ('2023-07-01 08:45:00', 26.5, 19),
       ('2023-07-01 09:00:00', 26.8, 19);

-- Inserting sensor readings for Humidity Sensor 7
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 20),
       ('2023-07-01 08:15:00', 61.0, 20),
       ('2023-07-01 08:30:00', 61.5, 20),
       ('2023-07-01 08:45:00', 61.8, 20),
       ('2023-07-01 09:00:00', 62.1, 20);

-- Inserting sensor readings for Pressure Sensor 7
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 21),
       ('2023-07-01 08:15:00', 1012.9, 21),
       ('2023-07-01 08:30:00', 1012.7, 21),
       ('2023-07-01 08:45:00', 1012.5, 21),
       ('2023-07-01 09:00:00', 1012.3, 21);

-- Inserting sensor readings for Temperature Sensor 8
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 22),
       ('2023-07-01 08:15:00', 26.0, 22),
       ('2023-07-01 08:30:00', 26.2, 22),
       ('2023-07-01 08:45:00', 26.5, 22),
       ('2023-07-01 09:00:00', 26.8, 22);

-- Inserting sensor readings for Humidity Sensor 8
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 23),
       ('2023-07-01 08:15:00', 61.0, 23),
       ('2023-07-01 08:30:00', 61.5, 23),
       ('2023-07-01 08:45:00', 61.8, 23),
       ('2023-07-01 09:00:00', 62.1, 23);

-- Inserting sensor readings for Pressure Sensor 8
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 24),
       ('2023-07-01 08:15:00', 1012.9, 24),
       ('2023-07-01 08:30:00', 1012.7, 24),
       ('2023-07-01 08:45:00', 1012.5, 24),
       ('2023-07-01 09:00:00', 1012.3, 24);

-- Inserting sensor readings for Temperature Sensor 9
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 25),
       ('2023-07-01 08:15:00', 26.0, 25),
       ('2023-07-01 08:30:00', 26.2, 25),
       ('2023-07-01 08:45:00', 26.5, 25),
       ('2023-07-01 09:00:00', 26.8, 25);

-- Inserting sensor readings for Humidity Sensor 9
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 26),
       ('2023-07-01 08:15:00', 61.0, 26),
       ('2023-07-01 08:30:00', 61.5, 26),
       ('2023-07-01 08:45:00', 61.8, 26),
       ('2023-07-01 09:00:00', 62.1, 26);

-- Inserting sensor readings for Pressure Sensor 9
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 27),
       ('2023-07-01 08:15:00', 1012.9, 27),
       ('2023-07-01 08:30:00', 1012.7, 27),
       ('2023-07-01 08:45:00', 1012.5, 27),
       ('2023-07-01 09:00:00', 1012.3, 27);

-- Inserting sensor readings for Temperature Sensor 10
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 28),
       ('2023-07-01 08:15:00', 26.0, 28),
       ('2023-07-01 08:30:00', 26.2, 28),
       ('2023-07-01 08:45:00', 26.5, 28),
       ('2023-07-01 09:00:00', 26.8, 28);

-- Inserting sensor readings for Humidity Sensor 10
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 29),
       ('2023-07-01 08:15:00', 61.0, 29),
       ('2023-07-01 08:30:00', 61.5, 29),
       ('2023-07-01 08:45:00', 61.8, 29),
       ('2023-07-01 09:00:00', 62.1, 29);

-- Inserting sensor readings for Pressure Sensor 10 
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 30),
       ('2023-07-01 08:15:00', 1012.9, 30),
       ('2023-07-01 08:30:00', 1012.7, 30),
       ('2023-07-01 08:45:00', 1012.5, 30),
       ('2023-07-01 09:00:00', 1012.3, 30);


-- Inserting sensor readings for Temperature Sensor 11
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 31),
       ('2023-07-01 08:15:00', 26.0, 31),
       ('2023-07-01 08:30:00', 26.2, 31),
       ('2023-07-01 08:45:00', 26.5, 31),
       ('2023-07-01 09:00:00', 26.8, 31);

-- Inserting sensor readings for Humidity Sensor 11
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 32),
       ('2023-07-01 08:15:00', 61.0, 32),
       ('2023-07-01 08:30:00', 61.5, 32),
       ('2023-07-01 08:45:00', 61.8, 32),
       ('2023-07-01 09:00:00', 62.1, 32);

-- Inserting sensor readings for Pressure Sensor 11 
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 33),
       ('2023-07-01 08:15:00', 1012.9, 33),
       ('2023-07-01 08:30:00', 1012.7, 33),
       ('2023-07-01 08:45:00', 1012.5, 33),
       ('2023-07-01 09:00:00', 1012.3, 33);

-- Inserting sensor readings for Temperature Sensor 12
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 34),
       ('2023-07-01 08:15:00', 26.0, 34),
       ('2023-07-01 08:30:00', 26.2, 34),
       ('2023-07-01 08:45:00', 26.5, 34),
       ('2023-07-01 09:00:00', 26.8, 34);

-- Inserting sensor readings for Humidity Sensor 12
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 35),
       ('2023-07-01 08:15:00', 61.0, 35),
       ('2023-07-01 08:30:00', 61.5, 35),
       ('2023-07-01 08:45:00', 61.8, 35),
       ('2023-07-01 09:00:00', 62.1, 35);

-- Inserting sensor readings for Pressure Sensor 12
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 36),
       ('2023-07-01 08:15:00', 1012.9, 36),
       ('2023-07-01 08:30:00', 1012.7, 36),
       ('2023-07-01 08:45:00', 1012.5, 36),
       ('2023-07-01 09:00:00', 1012.3, 36);

-- Inserting sensor readings for Temperature Sensor 13
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 37),
       ('2023-07-01 08:15:00', 26.0, 37),
       ('2023-07-01 08:30:00', 26.2, 37),
       ('2023-07-01 08:45:00', 26.5, 37),
       ('2023-07-01 09:00:00', 26.8, 37);

-- Inserting sensor readings for Humidity Sensor 13
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 38),
       ('2023-07-01 08:15:00', 61.0, 38),
       ('2023-07-01 08:30:00', 61.5, 38),
       ('2023-07-01 08:45:00', 61.8, 38),
       ('2023-07-01 09:00:00', 62.1, 38);

-- Inserting sensor readings for Pressure Sensor 13
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 39),
       ('2023-07-01 08:15:00', 1012.9, 39),
       ('2023-07-01 08:30:00', 1012.7, 39),
       ('2023-07-01 08:45:00', 1012.5, 39),
       ('2023-07-01 09:00:00', 1012.3, 39);

-- Inserting sensor readings for Temperature Sensor 14
INSERT INTO TEMPERATURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 25.5, 40),
       ('2023-07-01 08:15:00', 26.0, 40),
       ('2023-07-01 08:30:00', 26.2, 40),
       ('2023-07-01 08:45:00', 26.5, 40),
       ('2023-07-01 09:00:00', 26.8, 40);

-- Inserting sensor readings for Humidity Sensor 14
INSERT INTO HUMIDITY (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 60.2, 41),
       ('2023-07-01 08:15:00', 61.0, 41),
       ('2023-07-01 08:30:00', 61.5, 41),
       ('2023-07-01 08:45:00', 61.8, 41),
       ('2023-07-01 09:00:00', 62.1, 41);

-- Inserting sensor readings for Pressure Sensor 14
INSERT INTO PRESSURE (timestamp, value, sensor_id)
VALUES ('2023-07-01 08:00:00', 1013.2, 42),
       ('2023-07-01 08:15:00', 1012.9, 42),
       ('2023-07-01 08:30:00', 1012.7, 42),
       ('2023-07-01 08:45:00', 1012.5, 42),
       ('2023-07-01 09:00:00', 1012.3, 42);


