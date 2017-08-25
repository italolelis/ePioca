<template>
    <section>
        <h4>Bids for {{ threshold }}% yield</h4>
        <b-list-group>
            <auction-bid
                v-for="bid in bids"
                :key="bid.id"
                :user="bid.user.name"
                :value="bid.value"></auction-bid>
        </b-list-group>
    </section>
</template>

<script>
import { getBidsByAuctionAndThreshold } from '@/api/auction'
import { registerFor } from '@/api/ws'
import AuctionBid from '@/components/AuctionBid'

export default {
    props: {
        threshold: {
            type: Number,
            required: true
        },
        auctionId: {
            type: String,
            required: true
        }
    },

    components: {
        AuctionBid
    },

    data() {
        return {
            bids: []
        }
    },

    created() {
        this.loadBids()
        registerFor('bid_created', this.loadBids)
    },

    methods: {
        loadBids() {
            getBidsByAuctionAndThreshold(this.auctionId, this.threshold)
                .then(({ data }) => {
                    this.bids = data.sort((a, b) => {
                        if (a.value < b.value) {
                            return -1
                        }

                        if (a.value > b.value) {
                            return 1
                        }

                        return 0
                    })
                })
        }
    }
}
</script>

<style scoped>
.list-group-item:first-child {
    background-color: rgba(134, 184, 23, 0.4);
}
</style>
