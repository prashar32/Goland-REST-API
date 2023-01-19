CREATE TABLE "user"(
"UserId" INT PRIMARY KEY ,
"Name" VARCHAR(50),
"UserEmail" VARCHAR(50),
"UserContact" VARCHAR(50),
"ShopVisited" VARCHAR(50)
);

CREATE TABLE "shop"(
    "ShopId" INT PRIMARY KEY ,
    "ShopNo" INT ,
    "ShopName" VARCHAR(50),
    "FloorNo" INT,
    "ShopCategory" VARCHAR(50),
    "OpeningTime" VARCHAR(50),
    "ClosingTime" VARCHAR(50)
);