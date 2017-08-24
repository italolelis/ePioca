export const getAuctionFromId = (id) => {
    return Promise.resolve({
        "data": [
            {
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
    })
};
