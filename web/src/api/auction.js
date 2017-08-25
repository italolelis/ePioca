import client from '@/api'

export const createAuction = (auction) =>
    client.post('/auctions', auction)

export const getAuctionById = (auctionId) =>
    client.get(`/auctions/${auctionId}`)

export const getBidsByAuction = (auctionId, lowest = false) => {
    let url = `/auctions/${auctionId}/bids`
    if (lowest) {
        url = `${url}/lowest`
    }
    return client.get(url)
}
