<template>
    <section>
        <div class="row">
            <div class="col">
                <h2>{{ auction.ingredient.name }}</h2>
                <b-badge>{{ auction.status }}</b-badge>

                <p class="lead">
                    Auction to deliver {{ auction.qty }} units of {{ auction.ingredient.sku }} in {{ auction.week }}
                </p>
            </div>
        </div>

        <!-- TODO: Should be a component -->
        <section>
            <div class="row">
                <div class="col">
                    <h3>Lowest bids</h3>
                </div>
            </div>

            <div class="row">
                <div class="col" v-for="bid in lowestBids">
                    <low-bid
                        :user="bid.user.name"
                        :threshold="bid.threshold"
                        :value="bid.value"
                        ></low-bid>
                </div>
            </div>
        </section>

        <div class="row">

            <div class="col-md-8" v-if="!isBuyer">
                <bid-form
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :qty="auction.qty"></bid-form>
            </div>

            <div class="col-md-8" v-if="isBuyer">
                <auction-bid-list
                    class="auction-bid-list"
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :auction-id="auctionId"></auction-bid-list>
            </div>

            <div class="col">
                <div v-if="auction.status === 'running'">
                    Time remaining...
                </div>

                <div v-if="auction.status === 'scheduled'">
                    Scheduled for...
                </div>

                <div v-if="auction.status === 'completed'">
                    Finished at...
                </div>
            </div>

        </div>
    </section>
</template>

<script>
    import axios from 'axios';
    import config from '@/config'
    import { getUserRole } from '@/api'
    import { getAuctionById, getBidsByAuction } from '@/api/auction'
    import AuctionBidList from '@/components/AuctionBidList'
    import BidForm from '@/components/BidForm'
    import LowBid from '@/components/LowBid'

    export default {
        components: {
            AuctionBidList,
            BidForm,
            LowBid,
        },

        data() {
            return {
                auction: {
                    ingredient: {}
                },
                lowestBids: [],
            }
        },

        computed: {
            isBuyer() {
                return getUserRole() == 'supplier'
            },

            auctionId() {
                return this.$route.params.id
            }
        },

        created() {
            getAuctionById(this.auctionId)
                .then(res => this.auction = res.data)
                .catch(err => console.error(err))

            // Get low bids to show
            // TODO: should really be in its own component
            getBidsByAuction(this.auctionId, true)
                .then(res => this.lowestBids = res.data)
                .catch(err => console.error(err))
        },
    }
</script>

<style scoped>
.auction-bid-list {
    margin-top: 2rem;
}
</style>
