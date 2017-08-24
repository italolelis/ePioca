<template>
    <b-list-group>
        <auction-list-item
            v-for="auction in auctions"
            :key="auction.id"
            :id="auction.id"
            :ingredient="auction.ingredient"
            :qty="auction.qty"
            :startTime="auction.start_date"
            :duration="auction.duration"
            :type="filter"
            ></auction-list-item>
    </b-list-group>
</template>

<script>
    import AuctionListItem from '@/components/AuctionListItem'
    import axios from 'axios';
    import config from '@/config'

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
            this.getResponseFrom(`auctions?s=${this.filter}`)
        },

        methods: {
            getResponseFrom(url) {
                axios.get(`${config.api.uri}/${url}`)
                     .then(response => {
                         this.auctions = response.data
                     })
            }
        }
    }
</script>

<style scoped></style>
