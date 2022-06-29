# Cars Rent
__________________________

## Intro
#### i had create three container docker :
    - mysql container
    - phpmyadmin container
    - cars rent backend app
### For building & running the container:

    -type $ make build 

#### For running the container :

    -type $ make run

#### For testing the golang app :

    - make sure you had golang (v1.17)
    - import the databes using the sql file (./data/Cars.sql)
    - run the command $ make test

#### Server Port : 8888

#### API:
    
    -GetCars : /api/cars  (GET)
    -Add Car: /api/cars   (POST)
    -Update Rentals: /api/cars/:registration/rentals || /api/cars/:registration/returns (POST)
