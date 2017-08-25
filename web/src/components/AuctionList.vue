<template>
    <b-card :header="cardHeader"
        :border-variant="borderVariant">
        <b-list-group>
            <auction-list-item
                v-for="auction in auctions"
                :key="auction.id"
                :id="auction.id"
                :ingredient="auction.ingredient.name"
                :qty="auction.qty"
                :startTime="auction.start_date"
                :duration="auction.duration"
                :type="filter"></auction-list-item>

            <b-list-group-item v-if="auctions.length === 0" class="text-center">
                No {{ filter }} auctions
            </b-list-group-item>
        </b-list-group>
    </b-card>
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

    computed: {
        borderVariant() {
            switch (this.filter) {
                case 'running':
                    return 'success'
                default:
                    return 'secondary'
            }
        },

        icon() {
            switch (this.filter) {
                case 'running':
                    return 'fa-fighter-jet'
                case 'scheduled':
                    return 'fa-clock-o'
                case 'completed':
                    return 'fa-check'
            }
        },

        cardHeader() {
            const title = this.filter.charAt(0).toUpperCase() + this.filter.slice(1)

            return `<h5>${title} <i class="fa ${this.icon} pull-right"></i></h5>`
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

<style scoped>
/*.card {
    margin: 2rem 1rem;
}

.card:first-of-type {
    margin-left: 0;
    padding-left: 0;
}

.card:last-of-type {
    margin-right: 0;
    padding-right: 0;
}*/
</style>
