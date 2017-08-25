<template>
    <div class="container">
        <div class="alert alert-success" role="alert">
            <h4 class="alert-heading">Auction for {{ auction.ingredient }}</h4>
            <p> {{ auction.qty }} units </p>
        </div>
        <div class="row">
            <div v-for="threshold in auction.thresholds" class="col">
                <p> {{ threshold }} % </p>
                <p><strong> Lowest bids are </strong></p>
                <ul>
                    <li v-for="bid in lowestBids">
                        {{ bid.toFixed(2) }} $
                    </li>
                    <li> <strong> Your bid:  {{ userBid.value.toFixed(2) }} $ </strong></li>
                </ul>
                <p style="margin-bottom:0px;"><strong> Time Remaining </strong></p>
                <p> {{ auction.end_date }} </p>
                <input name="amount" class="form-control" v-model="amount" placeholder="Amount" required autofocus>
                <button type="submit" class="btn btn-primary btn-block" style="margin-top:5px;">Place bid</button>
            </div>
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
            getAuctionId() {
                return this.$route.params.id;
            },

            getUserId() {
                return localStorage.getItem('user_role');
            },

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

            getAuction() {
                getAuctionFromId(this.getAuctionId()).then(response => {this.auction = response.data[0] })
                                                     .catch(error => { console.error(error) })
            },

            getAuctionBids() {
                getBidsFromId(this.getAuctionId()).then(response => { this.lowestBids = this.getBidsSorted(response) })
                                                  .catch(error => { console.error(error) })
            },

            getUserBid(id) {
                getBidsFromId(this.getUserId()).then(response => { this.userBid = this.filterUserFrom(id, response) })
                                               .catch(error => { console.error(error) })
            }
        }
    }
</script>

<style scoped>
    .col{
        background-color: #f5f5f5;
        color: #000000;
        padding: 20px;
        margin:10px;
        -webkit-border-radius: 5px;
        -moz-border-radius: 5px;
        border-radius: 5px;
    }
    .col ul {
        list-style-type: none;
        margin-left:-30px;
    }
</style>
