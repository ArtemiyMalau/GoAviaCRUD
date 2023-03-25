-- provider seeder
INSERT INTO provider (provider.code, provider.name) VALUES ("AA", "AmericanAir");
SET @AA_provider_id = LAST_INSERT_ID();
INSERT INTO provider (provider.code, provider.name) VALUES ("IF", "InternationFlights");
SET @IF_provider_id = LAST_INSERT_ID();
INSERT INTO provider (provider.code, provider.name) VALUES ("RS", "RedStar");
SET @RS_provider_id = LAST_INSERT_ID();

-- airline seeder
INSERT INTO airline (airline.code, airline.name) VALUES ("SU", "Аэрофлот");
SET @SU_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("S7", "S7");
SET @S7_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("KV", "КрасАвиа");
SET @KV_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("U6", "Уральские авиалинии");
SET @U6_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("UT", "ЮТэйр");
SET @UT_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("FZ", "Flydubai)");
SET @FZ_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("JB", "JetBlue)");
SET @JB_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("SJ", "SuperJet");
SET @SJ_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("WZ", "Wizz Air");
SET @WZ_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("N4", "Nordwind Airline");
SET @N4_airline_id = LAST_INSERT_ID();
INSERT INTO airline (airline.code, airline.name) VALUES ("5N", "SmartAvia");
SET @5N_airline_id = LAST_INSERT_ID();


-- airline_provider seeder
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@FZ_airline_id, @AA_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@JB_airline_id, @AA_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@SJ_airline_id, @AA_provider_id);

INSERT INTO airline_provider (airline_id, provider_id) VALUES (@SU_airline_id, @IF_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@S7_airline_id, @IF_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@FZ_airline_id, @IF_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@N4_airline_id, @IF_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@JB_airline_id, @IF_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@WZ_airline_id, @IF_provider_id);

INSERT INTO airline_provider (airline_id, provider_id) VALUES (@SU_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@S7_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@KV_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@U6_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@UT_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@N4_airline_id, @RS_provider_id);
INSERT INTO airline_provider (airline_id, provider_id) VALUES (@5N_airline_id, @RS_provider_id);

-- scheme seeder
INSERT INTO scheme (name) VALUES ("Основная");
SET @MAIN_scheme_id = LAST_INSERT_ID();
INSERT INTO scheme (name) VALUES ("Тестовая");
SET @TEST_scheme_id = LAST_INSERT_ID();

-- scheme_provider seeder
INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (@MAIN_scheme_id, @AA_provider_id);
INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (@MAIN_scheme_id, @IF_provider_id);
INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (@MAIN_scheme_id, @RS_provider_id);

INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (@TEST_scheme_id, @IF_provider_id);
INSERT INTO scheme_provider (scheme_id, provider_id) VALUES (@TEST_scheme_id, @RS_provider_id);

-- account seeder
INSERT INTO account (scheme_id) VALUES (@TEST_scheme_id);
INSERT INTO account (scheme_id) VALUES (@TEST_scheme_id);
INSERT INTO account (scheme_id) VALUES (@MAIN_scheme_id);