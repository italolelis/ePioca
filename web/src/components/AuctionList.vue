<template>
    <ul>
        <li v-for="auction in auctions">
            <auction-list-item
                :ingredient="auction.ingredient"
                :qty="auction.qty"
                :type="filter"
            >
            </auction-list-item>
        </li>
    </ul>
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
