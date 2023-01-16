CREATE TABLE "Menus"(
    "Id" TEXT NOT NULL,
    "Name" TEXT NOT NULL,
    "Price" INTEGER NOT NULL,
    "Stock" INTEGER NOT NULL,
    "Url" TEXT NOT NULL,
    "CreatedAt" TIMESTAMP(0) WITH
        TIME zone NOT NULL,
        "UpdatedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL,
        "DeletedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL
);
ALTER TABLE
    "Menus" ADD PRIMARY KEY("Id");
CREATE TABLE "Users"(
    "Id" TEXT NOT NULL,
    "Name" TEXT NOT NULL,
    "Username" TEXT NOT NULL,
    "Password" TEXT NOT NULL,
    "CreatedAt" TIMESTAMP(0) WITH
        TIME zone NOT NULL,
        "UpdatedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL,
        "DeletedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL
);
ALTER TABLE
    "Users" ADD PRIMARY KEY("Id");
CREATE TABLE "Tables"(
    "Id" TEXT NOT NULL,
    "Number" INTEGER NOT NULL,
    "CreatedAt" TIMESTAMP(0) WITH
        TIME zone NOT NULL,
        "UpdatedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL,
        "DeletedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL
);
ALTER TABLE
    "Tables" ADD PRIMARY KEY("Id");
CREATE TABLE "Orders"(
    "Id" TEXT NOT NULL,
    "UserId" TEXT NOT NULL,
    "TableId" TEXT NOT NULL,
    "Qty" INTEGER NOT NULL,
    "TotalPrice" INTEGER NOT NULL,
    "CreatedAt" TIMESTAMP(0) WITH
        TIME zone NOT NULL,
        "UpdatedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL,
        "DeletedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL
);
ALTER TABLE
    "Orders" ADD PRIMARY KEY("Id");
CREATE TABLE "OrderDetails"(
    "Id" TEXT NOT NULL,
    "OrderId" TEXT NOT NULL,
    "MenuId" TEXT NOT NULL,
    "Qty" INTEGER NOT NULL,
    "TotalPrice" INTEGER NOT NULL,
    "CreatedAt" TIMESTAMP(0) WITH
        TIME zone NOT NULL,
        "UpdatedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL,
        "DeletedAt" TIMESTAMP(0)
    WITH
        TIME zone NOT NULL
);
ALTER TABLE
    "OrderDetails" ADD PRIMARY KEY("Id");
ALTER TABLE
    "OrderDetails" ADD CONSTRAINT "orderdetails_menuid_foreign" FOREIGN KEY("MenuId") REFERENCES "Menus"("Id");
ALTER TABLE
    "OrderDetails" ADD CONSTRAINT "orderdetails_orderid_foreign" FOREIGN KEY("OrderId") REFERENCES "Orders"("Id");
ALTER TABLE
    "Orders" ADD CONSTRAINT "orders_tableid_foreign" FOREIGN KEY("TableId") REFERENCES "Tables"("Id");