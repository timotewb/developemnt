-- ------------------------------
-- OPTION
-- ------------------------------

OPTION IMPORT;

-- ------------------------------
-- TABLE: account
-- ------------------------------

DEFINE TABLE account SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: article
-- ------------------------------

DEFINE TABLE article SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: author
-- ------------------------------

DEFINE TABLE author SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: company
-- ------------------------------

DEFINE TABLE company SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: openweather
-- ------------------------------

DEFINE TABLE openweather SCHEMALESS PERMISSIONS NONE;

-- ------------------------------
-- TABLE: openweather_city_list
-- ------------------------------

DEFINE TABLE openweather_city_list SCHEMALESS PERMISSIONS NONE;

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
-- TABLE DATA: account
-- ------------------------------

UPDATE account:pqj8z98jv2k00r83xnhs CONTENT { created_at: "2022-09-25T06:08:37.614841900Z", id: account:pqj8z98jv2k00r83xnhs, name: "ACME Inc" };

-- ------------------------------
-- TABLE DATA: article
-- ------------------------------

UPDATE article:s8glp6eyg0ruuznf34de CONTENT { account: account:pqj8z98jv2k00r83xnhs, author: author:john, created_at: "2022-09-25T06:09:44.873383400Z", id: article:s8glp6eyg0ruuznf34de, text: "Donec eleifend, nunc vitae commodo accumsan, mauris est fringilla.", title: "Lorem ipsum dolor" };

-- ------------------------------
-- TABLE DATA: author
-- ------------------------------

UPDATE author:john CONTENT { admin: true, age: 29, id: author:john, name: { first: "John", full: "John Adams", last: "Adams" }, signup_at: "2022-09-25T06:09:12.831989Z" };

-- ------------------------------
-- TABLE DATA: company
-- ------------------------------

UPDATE company:3brnvttmesdvgkrqaxe7 CONTENT { founded: "2021-09-10T00:00:00Z", id: company:3brnvttmesdvgkrqaxe7, name: "SurrealDB", tags: ["big data", "database"] };
UPDATE company:zoht4lg3463mn4fq8iti CONTENT { founded: "2021-09-10T00:00:00Z", id: company:zoht4lg3463mn4fq8iti, name: "SurrealDB", tags: ["big data", "database"] };

-- ------------------------------
-- TABLE DATA: openweather
-- ------------------------------


-- ------------------------------
-- TABLE DATA: openweather_city_list
-- ------------------------------


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

