import client from '@/api'
import { getUserName, getUserId } from '@/api'

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

export const postBid = (auctionId, threshold, value) =>
    client.post(`/auctions/${auctionId}/bids`, {
        user: {
            id: getUserId(),
            name: getUserName()
        },
        threshold,
        value,
    })
