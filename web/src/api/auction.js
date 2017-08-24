import client from '@/api'

export const createAuction = (auction) =>
    client.post('/auctions', auction)
