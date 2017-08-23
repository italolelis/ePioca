# ePioca


## Key Features

Features...

## Installation

Instructions...

## Getting Started

Instructions.....

## Documentation
## API specs

### Auction

`GET /auctions/{ID}`
```json
{
    "data": [
        {
            "status" : "completed/scheduled/running",
            "week" : "2016-W10",
            "country" : "US",
            "dc" : "TX",
            "ingredient" : "PHF-10-10008-4",
            "duration" : 3600,
            "start_date" : "2017-08-23 10:00:00",
            "end_date" : "2017-08-25 10:00:00",
            "qty" : 2456.50,
            "thresholds" : [
                30,
                40,
                50,
            ],
            "max_price" : 5000
        }
    ],
    "status": 200 
}
```

`GET /auctions`
```json
{
    "data": [
        {
            "status" : "running",
            "week" : "2016-W10",
            "country" : "US",
            "dc" : "TX",
            "ingredient" : "PHF-10-10008-4",
            "duration" : 3600,
            "start_date" : "2017-08-23 10:00:00",
            "end_date" : "2017-08-25 10:00:00",
            "qty" : 2456.50,
            "thresholds" : [
                30,
                40,
                50,
            ],
            "max_price" : 5000
        },
        {
            "status" : "completed",
            "week" : "2016-W11",
            "country" : "US",
            "dc" : "NJ",
            "ingredient" : "PHF-10-10005-4",
            "duration" : 3600,
            "start_date" : "2017-08-23 10:00:00",
            "end_date" : "2017-08-25 10:00:00",
            "qty" : 2456.50,
            "thresholds" : [
                30,
                40,
                50,
            ],
            "max_price" : 1000
        }
    ],
    "status": 200
}    
```

`POST /auctions`
```json
{
    "status" : "running",
    "week" : "2016-W10",
    "country" : "US",
    "dc" : "TX",
    "ingredient" : "PHF-10-10008-4",
    "duration" : 3600,
    "start_date" : "2017-08-23 10:00:00",
    "end_date" : "2017-08-25 10:00:00",
    "qty" : 2456.50,
    "thresholds" : [
        30,
        40,
        50,
    ],
    "max_price" : 5000
}
```

`PUT /auctions/{ID}`
```json
{
    "status" : "running",
    "week" : "2016-W10",
    "country" : "US",
    "dc" : "TX",
    "ingredient" : "PHF-10-10008-4",
    "duration" : 3600,
    "start_date" : "2017-08-23 10:00:00",
    "end_date" : "2017-08-25 10:00:00",
    "qty" : 2456.50,
    "thresholds" : [
        40,
        50,
    ],
    "max_price" : 5000
}
```

### Bid

`GET /auctions/{ID}/bids`
```json
{
    "data": [
        {
            "auction_id" : 123,
            "user_id" : 456,
            "threshold" : 40,
            "price" : 230.50
        },
        {
            "auction_id" : 234,
            "user_id" : 456,
            "threshold" : 50,
            "price" : 250.10
        }
    ],
    "status": 200 
}
```

`GET /auctions/{ID}/bids/{bidID}`
```json
{
    "data": [
        {
            "auction_id" : 123,
            "user_id" : 456,
            "threshold" : 40,
            "price" : 230.50
        }
    ],
    "status": 200
}
```

`POST /auctions/{ID}/bids`
```json
{
    "user_id" : 456,
    "threshold" : 40,
    "price" : 230.50
} 
```

`PUT /auctions/{ID}/bids/{bidID}`
```json
{
    "user_id" : 456,
    "threshold" : 20,
    "price" : 230.50
} 
```

### Auction positions

`GET /auctions/{ID}/positions`
```json
{
    "data": [
        {
            "user_id" : 456,
            "position" : 1,
            "price" : 230.50
        },
        {
            "user_id" : 456,
            "position" : 2,
            "price" : 333.50
        },
        {
            "user_id" : 456,
            "position" : 3,
            "price" : 555
        },
    ],
    "status": 200 
}
```