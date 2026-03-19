CREATE TABLE Users (
    user_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Categories (
    category_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Locations (
    location_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100)
);

CREATE TABLE Keywords (
    keyword_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    word VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE Cases (
    case_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    summary TEXT,
    occurred_on DATE,
    status VARCHAR(20) CHECK (status IN ('Open', 'Closed', 'Cold', 'Ongoing')),
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    category_id INT REFERENCES Categories(category_id) ON DELETE SET NULL,
    location_id INT REFERENCES Locations(location_id) ON DELETE SET NULL
);

CREATE TABLE Case_Keywords (
    case_id INT REFERENCES Cases(case_id) ON DELETE CASCADE,
    keyword_id INT REFERENCES Keywords(keyword_id) ON DELETE CASCADE,
    PRIMARY KEY (case_id, keyword_id)
);

CREATE TABLE User_Bookmarks (
    user_id INT REFERENCES Users(user_id) ON DELETE CASCADE,
    case_id INT REFERENCES Cases(case_id) ON DELETE CASCADE,
    saved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, case_id)
);