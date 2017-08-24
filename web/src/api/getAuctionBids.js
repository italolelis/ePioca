export const getAuctionBids = (id) => {
    return Promise.resolve({
        "data": [
            {
                "auction_id" : '3c89d209-09cc-4be0-86b6-da94145440f6',
                "user_id" : 1,
                "threshold" : 40,
                "value" : 230.50
            },
            {
                "auction_id" : '3c89d209-09cc-4be0-86b6-da94145440f6',
                "user_id" : 2,
                "threshold" : 50,
                "value" : 240
            },
            {
                "auction_id" : '3c89d209-09cc-4be0-86b6-da94145440f6',
                "user_id" : 3,
                "threshold" : 50,
                "value" : 300
            },
            {
                "auction_id" : '3c89d209-09cc-4be0-86b6-da94145440f6',
                "user_id" : 4,
                "threshold" : 50,
                "value" : 100
            }
        ],
        "status": 200
    })
};
