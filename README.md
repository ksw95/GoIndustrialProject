# Go Industrial Project
This industrial project is the technical solution to the client brief by Foodpanda.

The task was to create an alert in the cart when added items do not meet with users predefined health conditions and to encourage users to do self pick up.

The solution is develop using Go Language.

Sample data will be added into the application to simulate restaurants and foods.

## Developers
- Koh Shao Wei
- Ahmad Bahrudin
- Xu Ran

## Current Flow

Registration:
```
User Creates Account -> Parameters Sent To API -> Checks Username -> Creates Account
```
Login:
```
User Logs In -> Parameters Sent To API -> Checks User & Pass -> Assign Token & User Condition
-> Store In Client Side
```
Restaurant:
```
Search Restaurant -> Return List Of Restaurants from API/Database -> Choose Restaurants 
-> Retrieve Foods From Restaurant in API/Databse -> Store in Client Side
```
Food:
```
Browse Food From Restaurant -> Select Food -> Add Food Info to Cart from Client Side 
-> Check Total Calories Against User Condition on Client Side -> If Exceed Max Calories Send Alert
```
Checkout:
```
Cart Checkout -> Change Address (If needed) -> Calculate Distance Between Address & Restaurant 
-> Calculate Calories Burn From Walking (Maybe Cycling?) -> Choose Food Collection Method (Self Pick Up/Deliver)
```

## Database Tables

Account:

    - Username    string (Primary key , unique identifier for the account)
    - Password    string
    - Address     string
    - PostalCode  int

Condition:

    - Username    string (Primary key , unique identifier for the condition)
    - MaxCalories int
    - LastLogin   string
    - Diabetic    string
    - Halal       string
    - Vegan       string

Food:

    - ID          int (Primary key , unique identifier for the food)
    - Name        string
    - ShopID      int
    - Description string
    - Calories    int
    - Sugary      string
    - Halal       string
    - Vegan       string

Restaurant:

    - ID          int (Primary key , unique identifier for the restaurant)
    - Name        string
    - Description string
    - Address     string
    - PostalCode  int
