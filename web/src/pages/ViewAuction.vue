<template>
    <div class="container">
        <div>
            Auction for {{ auction.qty }} units of {{ auction.ingredient }}
        </div>
        <div>
            <p> Lowest bids are </p>
            <ul>
                <li v-for="bid in lowestBids">
                    {{ bid }}
                </li>
            </ul>
        </div>
        <div>
            My Bid: {{ userBid.value }}
        </div>
        <div>
            time remaining
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import config from '@/config'
    import { getAuctionFromId } from '@/api/getAuctionFromId'
    import { getBidsFromId } from '@/api/getBidsFromId'

    export default {
        data() {
            return {
                auction : {},
                lowestBids: {},
                userBid: {}
            }
        },

        mounted() {
            this.getAuction('theId');
            this.getAuctionBids('theId');
            this.getUserBid(2)
        },

        methods: {
            sortBids(bidA, bidB) {
                return bidA - bidB;
            },

            filterUserFrom(id, response) {
                return response.data.filter( (object) => {
                    return id === object.user_id;
                })[0];
            },

            getBidsSorted(response) {
                return response.data.map( (object) => { return object.value; })
                                    .sort(this.sortBids)
                                    .slice(0, 3);
            },

            getAuction(id) {
                getAuctionFromId(id).then(response => {this.auction = response.data[0] })
                              .catch(error => { console.error(error) })
            },

            getAuctionBids(id) {
                getBidsFromId(id).then(response => { this.lowestBids = this.getBidsSorted(response) })
                                  .catch(error => { console.error(error) })
            },

            getUserBid(id) {
                getBidsFromId(id).then(response => { this.userBid = this.filterUserFrom(id, response) })
                                       .catch(error => { console.error(error) })
            }
        }
    }
</script>

<style scoped></style>
