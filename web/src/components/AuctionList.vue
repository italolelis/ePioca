<template>
    <b-list-group>
        <auction-list-item
            v-for="auction in auctions"
            :key="auction.id"
            :id="auction.id"
            :ingredient="auction.ingredient.name"
            :qty="auction.qty"
            :startTime="auction.start_date"
            :duration="auction.duration"
            :type="filter"
            ></auction-list-item>
    </b-list-group>
</template>

<script>
import AuctionListItem from '@/components/AuctionListItem'
import { getAuctionsByStatus } from '@/api/auction';
import config from '@/config'
import { registerFor } from '@/api/ws'

export default {
    components: {
        AuctionListItem
    },

    props: {
        filter: {
            type: String
        }
    },

    data() {
        return {
            auctions: []
        }
    },

    mounted() {
        this.loadAuctions()
        registerFor('auction_created', this.loadAuctions)
    },

    methods: {
        loadAuctions() {
            getAuctionsByStatus(this.filter)
                .then(({ data }) => this.auctions = data)
                .catch(err => console.error(err))
        }
    }
}
</script>

<style scoped></style>
