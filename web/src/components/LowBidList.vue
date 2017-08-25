<template>
    <section>
        <div class="row" v-if="lowestBids">
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
</template>

<script>
import { getLowestBidsByAuction } from '@/api/auction'
import { registerFor } from '@/api/ws'
import LowBid from '@/components/LowBid'

export default {
    components: {
        LowBid
    },

    data() {
        return {
            lowestBids: []
        }
    },

    props: {
        auctionId: {
            type: String,
            required: true
        }
    },

    created() {
        this.loadLowestBids()
        registerFor('bid_created', this.loadLowestBids)
    },

    methods: {
        loadLowestBids() {
            getLowestBidsByAuction(this.auctionId)
                .then(res => this.lowestBids = res.data)
                .catch(err => console.error(err))
        }
    }
}
</script>

<style>

</style>
