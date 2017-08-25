<template>
    <section>
        <div class="row">
            <auction-header class="col"
                :ingredient="auction.ingredient.name"
                :sku="auction.ingredient.sku"
                status="Completed"
                :week="auction.week"
                :qty="auction.qty"
            ></auction-header>
        </div>

        <article v-if="isSupplier">
            <div class="row">
                <low-bid-list class="col" :auction-id="auctionId"></low-bid-list>
            </div>

            <div class="row">
                <auction-bid-list
                    class="auction-bid-list col"
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :auction-id="auctionId"
                    :user-id="userId"></auction-bid-list>
            </div>
        </article>

        <article v-else>
            <div class="row">
                <winning-bid-list
                    class="auction-bid-list col"
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :auction-id="auctionId"></winning-bid-list>
            </div>
        </article>

        <router-link class="btn btn-block" :to="{ name: 'Dashboard' }">
            &larr; Back home
        </router-link>
    </section>
</template>

<script>
import moment from 'moment'
import { getUserId, getUserName, getUserRole } from '@/api/'
import { getAuctionById, getLowestBidsByAuction, getWinners } from '@/api/auction'
import AuctionBidList from '@/components/AuctionBidList'
import AuctionHeader from '@/components/AuctionHeader'
import LowBidList from '@/components/LowBidList'
import WinningBidList from '@/components/WinningBidList'

export default {
    data() {
        return {
            bids: [],
            auction: {
                ingredient: {
                    name: '',
                    sku: ''
                },
                week: '',
                qty: 0,
                threshold: []
            },
            myBids: []
        }
    },

    components : {
        AuctionBidList,
        AuctionHeader,
        WinningBidList,
        LowBidList,
    },

    computed: {
        isSupplier() {
            return getUserRole() === 'supplier'
        },

        auctionId() {
            return this.$route.params.id
        },

        userId() {
            return getUserId()
        }
    },

    created() {
        getAuctionById(this.auctionId)
            .then(({ data }) => this.auction = data)

        if (!this.isSupplier) {
            getWinners(this.auctionId)
                .then(({ data }) => this.bids = data)
        }
    }
}
</script>

<style scoped>
.auction-bid-list {
    margin-top: 2rem;
}
</style>
