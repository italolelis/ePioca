import client from '@/api'
import { getUserName, getUserId } from '@/api'

export const createAuction = (auction) =>
    client.post('/auctions', auction);

export const getAuctionsByStatus = (status) =>
    client.get(`/auctions?s=${status}`)

export const getAuctionById = (auctionId) =>
    client.get(`/auctions/${auctionId}`);

export const getLowestBidsByAuction = (auctionId) =>
    client.get(`/auctions/${auctionId}/bids/lowest`)

export const getBidsByAuctionAndThreshold = (auctionId, threshold) =>
    client.get(`/auctions/${auctionId}/bids`, { params: { threshold } })

export const getBidsByAuction = (auctionId, lowest = false) => {
    let url = `/auctions/${auctionId}/bids`;
    if (lowest) {
        url = `${url}/lowest`
    }
    return client.get(url)
};

export const getWinners = (auctionId) => {
    return client.get(`/auctions/${auctionId}/winners`)
};

export const postBid = (auctionId, threshold, value) =>
    client.post(`/auctions/${auctionId}/bids`, {
        user: {
            id: getUserId(),
            name: getUserName()
        },
        threshold,
        value,
    });
