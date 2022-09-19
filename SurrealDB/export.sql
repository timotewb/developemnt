-- ------------------------------
-- OPTION
-- ------------------------------

OPTION IMPORT;

-- ------------------------------
-- TABLE: company
-- ------------------------------

DEFINE TABLE company SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: openweather
-- ------------------------------

DEFINE TABLE openweather SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: temperature
-- ------------------------------

DEFINE TABLE temperature SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: test02
-- ------------------------------

DEFINE TABLE test02 SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: tmp_openweather
-- ------------------------------

DEFINE TABLE tmp_openweather SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TRANSACTION
-- ------------------------------

BEGIN TRANSACTION;

-- ------------------------------
-- TABLE DATA: company
-- ------------------------------

UPDATE company:3brnvttmesdvgkrqaxe7 CONTENT { founded: "2021-09-10T00:00:00Z", id: company:3brnvttmesdvgkrqaxe7, name: "SurrealDB", tags: ["big data", "database"] };
UPDATE company:zoht4lg3463mn4fq8iti CONTENT { founded: "2021-09-10T00:00:00Z", id: company:zoht4lg3463mn4fq8iti, name: "SurrealDB", tags: ["big data", "database"] };

-- ------------------------------
-- TABLE DATA: openweather
-- ------------------------------

UPDATE openweather:2179537 CONTENT { base: "stations", clouds: { all: 40 }, cod: 200, coord: { lat: -41.2866, lon: 174.7756 }, dt: 1663467753, id: openweather:2179537, main: { feels_like: 15.15, humidity: 63, pressure: 1022, temp: 15.86, temp_max: 15.87, temp_min: 14.81 }, name: "Wellington", sys: { country: "NZ", id: 2007945, sunrise: 1663438695, sunset: 1663481549, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "scattered clouds", icon: "03d", id: 802, main: "Clouds" }], wind: { deg: 360, speed: 8.75 } };
UPDATE openweather:2182560 CONTENT { base: "stations", clouds: { all: 12 }, cod: 200, coord: { lat: -38, lon: 175.5 }, dt: 1663467760, id: openweather:2182560, main: { feels_like: 14.93, grnd_level: 1013, humidity: 55, pressure: 1025, sea_level: 1025, temp: 15.85, temp_max: 17.62, temp_min: 15.76 }, name: "Bay of Plenty Region", sys: { country: "NZ", id: 2038613, sunrise: 1663438486, sunset: 1663481411, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "few clouds", icon: "02d", id: 801, main: "Clouds" }], wind: { deg: 341, gust: 2.1, speed: 1.97 } };
UPDATE openweather:2183411 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -45.9, lon: 168.75 }, dt: 1663467231, id: openweather:2183411, main: { feels_like: 14.33, grnd_level: 999, humidity: 47, pressure: 1014, sea_level: 1014, temp: 15.5, temp_max: 15.5, temp_min: 15.5 }, name: "Riversdale", sys: { country: "NZ", sunrise: 1663440195, sunset: 1663482941 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04d", id: 804, main: "Clouds" }], wind: { deg: 105, gust: 2.1, speed: 0.55 } };
UPDATE openweather:1663488815_2183411 CONTENT { base: "stations", clouds: { all: 98 }, cod: 200, coord: { lat: -45.9, lon: 168.75 }, dt: 1663488815, id: openweather:1663488815_2183411, location_id: 2183411, main: { feels_like: 8.86, grnd_level: 997, humidity: 82, pressure: 1013, sea_level: 1013, temp: 9.25, temp_max: 9.25, temp_min: 9.25 }, name: "Riversdale", sys: { country: "NZ", sunrise: 1663440195, sunset: 1663482941 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 23, gust: 2.29, speed: 1.44 } };
UPDATE openweather:1663489040_2179103 CONTENT { base: "stations", clouds: { all: 97 }, cod: 200, coord: { lat: -36.8333, lon: 175.7 }, dt: 1663489040, id: openweather:1663489040_2179103, location_id: 2179103, main: { feels_like: 10.53, grnd_level: 1027, humidity: 79, pressure: 1028, sea_level: 1028, temp: 11.28, temp_max: 11.28, temp_min: 11.28 }, name: "Whitianga", sys: { country: "NZ", sunrise: 1663438426, sunset: 1663481376 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 342, gust: 4.17, speed: 2.58 } };
UPDATE openweather:1663489169_2178896 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -42.95, lon: 171.5667 }, dt: 1663489169, id: openweather:1663489169_2178896, location_id: 2178896, main: { feels_like: 1.5, grnd_level: 932, humidity: 92, pressure: 1023, sea_level: 1023, temp: 3.86, temp_max: 3.86, temp_min: 3.86 }, name: "Arthur's Pass", rain: { 1h: 0.22 }, sys: { country: "NZ", sunrise: 1663439484, sunset: 1663482301 }, timezone: 43200, visibility: 10000, weather: [{ description: "light rain", icon: "10n", id: 500, main: "Rain" }], wind: { deg: 320, gust: 3.14, speed: 2.56 } };
UPDATE openweather:1663489170_2178933 CONTENT { base: "stations", clouds: { all: 99 }, cod: 200, coord: { lat: -46.3333, lon: 168.85 }, dt: 1663489170, id: openweather:1663489170_2178933, location_id: 2178933, main: { feels_like: 7.14, grnd_level: 1009, humidity: 81, pressure: 1012, sea_level: 1012, temp: 8.2, temp_max: 8.2, temp_min: 8.2 }, name: "Wyndham", sys: { country: "NZ", id: 2008959, sunrise: 1663440177, sunset: 1663482912, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 60, gust: 6.79, speed: 1.95 } };
UPDATE openweather:1663489217_2182560 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -38, lon: 175.5 }, dt: 1663489217, id: openweather:1663489217_2182560, location_id: 2182560, main: { feels_like: 8.46, grnd_level: 1014, humidity: 93, pressure: 1027, sea_level: 1027, temp: 9.29, temp_max: 10.64, temp_min: 9.29 }, name: "Bay of Plenty Region", sys: { country: "NZ", id: 2007514, sunrise: 1663438486, sunset: 1663481411, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 29, gust: 1.83, speed: 1.89 } };
UPDATE openweather:1663490000_2178896 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -42.95, lon: 171.5667 }, dt: 1663490000, id: openweather:1663490000_2178896, location_id: 2178896, main: { feels_like: 1.18, grnd_level: 932, humidity: 92, pressure: 1023, sea_level: 1023, temp: 3.73, temp_max: 3.73, temp_min: 3.73 }, name: "Arthur's Pass", sys: { country: "NZ", sunrise: 1663439484, sunset: 1663482301 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 322, gust: 3.23, speed: 2.75 } };
UPDATE openweather:1663490002_2178933 CONTENT { base: "stations", clouds: { all: 99 }, cod: 200, coord: { lat: -46.3333, lon: 168.85 }, dt: 1663490002, id: openweather:1663490002_2178933, location_id: 2178933, main: { feels_like: 7.09, grnd_level: 1009, humidity: 86, pressure: 1012, sea_level: 1012, temp: 8.2, temp_max: 8.2, temp_min: 8.2 }, name: "Wyndham", sys: { country: "NZ", id: 2008959, sunrise: 1663440177, sunset: 1663482912, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 56, gust: 7.49, speed: 2.01 } };
UPDATE openweather:1663490003_2179103 CONTENT { base: "stations", clouds: { all: 98 }, cod: 200, coord: { lat: -36.8333, lon: 175.7 }, dt: 1663490003, id: openweather:1663490003_2179103, location_id: 2179103, main: { feels_like: 10.52, grnd_level: 1027, humidity: 79, pressure: 1028, sea_level: 1028, temp: 11.27, temp_max: 11.27, temp_min: 11.27 }, name: "Whitianga", sys: { country: "NZ", sunrise: 1663438426, sunset: 1663481376 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 345, gust: 4.8, speed: 2.88 } };
UPDATE openweather:1663490012_2182560 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -38, lon: 175.5 }, dt: 1663490012, id: openweather:1663490012_2182560, location_id: 2182560, main: { feels_like: 9.73, grnd_level: 1014, humidity: 94, pressure: 1028, sea_level: 1028, temp: 10.2, temp_max: 10.21, temp_min: 9.29 }, name: "Bay of Plenty Region", sys: { country: "NZ", id: 2038613, sunrise: 1663438486, sunset: 1663481411, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 29, gust: 1.6, speed: 1.72 } };
UPDATE openweather:1663576363_2178896 CONTENT { base: "stations", clouds: { all: 89 }, cod: 200, coord: { lat: -42.95, lon: 171.5667 }, dt: 1663576363, id: openweather:1663576363_2178896, location_id: 2178896, main: { feels_like: 3.32, grnd_level: 927, humidity: 95, pressure: 1016, sea_level: 1016, temp: 5.79, temp_max: 5.79, temp_min: 5.79 }, name: "Arthur's Pass", sys: { country: "NZ", sunrise: 1663525775, sunset: 1663568765 }, timezone: 43200, visibility: 7809, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 319, gust: 3.91, speed: 3.19 } };
UPDATE openweather:1663576364_2178933 CONTENT { base: "stations", clouds: { all: 16 }, cod: 200, coord: { lat: -46.3333, lon: 168.85 }, dt: 1663576364, id: openweather:1663576364_2178933, location_id: 2178933, main: { feels_like: 7.11, grnd_level: 1006, humidity: 76, pressure: 1009, sea_level: 1009, temp: 8.2, temp_max: 8.2, temp_min: 8.2 }, name: "Wyndham", sys: { country: "NZ", id: 2008959, sunrise: 1663526457, sunset: 1663569387, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "few clouds", icon: "02n", id: 801, main: "Clouds" }], wind: { deg: 338, gust: 8.09, speed: 1.98 } };
UPDATE openweather:1663576365_2179103 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -36.8333, lon: 175.7 }, dt: 1663576365, id: openweather:1663576365_2179103, location_id: 2179103, main: { feels_like: 14.85, grnd_level: 1021, humidity: 97, pressure: 1022, sea_level: 1022, temp: 14.78, temp_max: 14.78, temp_min: 14.78 }, name: "Whitianga", sys: { country: "NZ", sunrise: 1663524734, sunset: 1663567823 }, timezone: 43200, visibility: 8000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 358, gust: 14.68, speed: 8.26 } };
UPDATE openweather:1663576374_2183411 CONTENT { base: "stations", clouds: { all: 44 }, cod: 200, coord: { lat: -45.9, lon: 168.75 }, dt: 1663576374, id: openweather:1663576374_2183411, location_id: 2183411, main: { feels_like: 6.47, grnd_level: 995, humidity: 89, pressure: 1010, sea_level: 1010, temp: 7.81, temp_max: 7.81, temp_min: 7.81 }, name: "Riversdale", sys: { country: "NZ", sunrise: 1663526477, sunset: 1663569415 }, timezone: 43200, visibility: 10000, weather: [{ description: "scattered clouds", icon: "03n", id: 802, main: "Clouds" }], wind: { deg: 340, gust: 4.51, speed: 2.2 } };
UPDATE openweather:1663576385_2182560 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -38, lon: 175.5 }, dt: 1663576385, id: openweather:1663576385_2182560, location_id: 2182560, main: { feels_like: 14.22, grnd_level: 1009, humidity: 98, pressure: 1021, sea_level: 1021, temp: 14.19, temp_max: 15.84, temp_min: 14.09 }, name: "Bay of Plenty Region", sys: { country: "NZ", id: 2038613, sunrise: 1663524791, sunset: 1663567862, type: 2 }, timezone: 43200, visibility: 7032, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 15, gust: 8.61, speed: 3.01 } };
UPDATE openweather:1663577859_2179103 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -36.8333, lon: 175.7 }, dt: 1663577859, id: openweather:1663577859_2179103, location_id: 2179103, main: { feels_like: 14.85, grnd_level: 1021, humidity: 97, pressure: 1022, sea_level: 1022, temp: 14.78, temp_max: 14.78, temp_min: 14.78 }, name: "Whitianga", sys: { country: "NZ", sunrise: 1663524734, sunset: 1663567823 }, timezone: 43200, visibility: 8000, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 358, gust: 14.68, speed: 8.26 } };
UPDATE openweather:1663578107_2178896 CONTENT { base: "stations", clouds: { all: 89 }, cod: 200, coord: { lat: -42.95, lon: 171.5667 }, dt: 1663578107, id: openweather:1663578107_2178896, location_id: 2178896, main: { feels_like: 3.32, grnd_level: 927, humidity: 95, pressure: 1016, sea_level: 1016, temp: 5.79, temp_max: 5.79, temp_min: 5.79 }, name: "Arthur's Pass", sys: { country: "NZ", sunrise: 1663525775, sunset: 1663568765 }, timezone: 43200, visibility: 7809, weather: [{ description: "overcast clouds", icon: "04n", id: 804, main: "Clouds" }], wind: { deg: 319, gust: 3.91, speed: 3.19 } };
UPDATE openweather:1663578108_2178933 CONTENT { base: "stations", clouds: { all: 16 }, cod: 200, coord: { lat: -46.3333, lon: 168.85 }, dt: 1663578108, id: openweather:1663578108_2178933, location_id: 2178933, main: { feels_like: 7.11, grnd_level: 1006, humidity: 76, pressure: 1009, sea_level: 1009, temp: 8.2, temp_max: 8.2, temp_min: 8.2 }, name: "Wyndham", sys: { country: "NZ", id: 2008959, sunrise: 1663526457, sunset: 1663569387, type: 2 }, timezone: 43200, visibility: 10000, weather: [{ description: "few clouds", icon: "02n", id: 801, main: "Clouds" }], wind: { deg: 338, gust: 8.09, speed: 1.98 } };
UPDATE openweather:1663578119_2182560 CONTENT { base: "stations", clouds: { all: 100 }, cod: 200, coord: { lat: -38, lon: 175.5 }, dt: 1663578119, id: openweather:1663578119_2182560, location_id: 2182560, main: { feels_like: 14.81, grnd_level: 1009, humidity: 98, pressure: 1021, sea_level: 1021, temp: 14.72, temp_max: 15.84, temp_min: 14.65 }, name: "Bay of Plenty Region", rain: { 1h: 0.15 }, sys: { country: "NZ", id: 2038613, sunrise: 1663524791, sunset: 1663567862, type: 2 }, timezone: 43200, visibility: 7032, weather: [{ description: "light rain", icon: "10n", id: 500, main: "Rain" }], wind: { deg: 15, gust: 8.61, speed: 3.01 } };

-- ------------------------------
-- TABLE DATA: temperature
-- ------------------------------

UPDATE temperature:{ date: $now, location: "London" } CONTENT { date: "2022-09-18T07:56:37.350028200Z", id: temperature:{ date: $now, location: "London" }, location: "London", temperature: 23.7 };
UPDATE temperature:["London", $now] CONTENT { date: "2022-09-18T07:56:37.350028200Z", id: temperature:["London", $now], location: "London", temperature: 23.7 };

-- ------------------------------
-- TABLE DATA: test02
-- ------------------------------

UPDATE test02:1002 CONTENT { id: test02:1002, name: "tim", val: 1.023 };

-- ------------------------------
-- TABLE DATA: tmp_openweather
-- ------------------------------


-- ------------------------------
-- TRANSACTION
-- ------------------------------

COMMIT TRANSACTION;

