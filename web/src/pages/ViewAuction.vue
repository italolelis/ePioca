<template>
    <section>
        <div class="row">
            <div class="col">
                <h2>
                    {{ auction.ingredient.name }}
                    <b-badge>{{ auction.status }}</b-badge>
                </h2>

                <p class="lead">
                    Auction to supply {{ auction.qty }} units of {{ auction.ingredient.sku }} in {{ auction.week }}
                </p>
            </div>

            <div class="col-md-4">
                <b-card v-if="auction.status === 'scheduled'" class="text-center">
                    Scheduled for {{ scheduledDate }}
                </b-card>

                <b-card v-else class="text-center">
                    <epioca-countdown :end-date="endDate"></epioca-countdown>
                </b-card>
            </div>
        </div>

        <article v-if="auction.status !== 'scheduled'">

            <low-bid-list :auction-id="auctionId"></low-bid-list>

            <div class="row">

                <div class="col" v-if="!isBuyer">
                    <bid-form
                        v-for="threshold in auction.threshold"
                        :key="threshold"
                        :threshold="threshold"
                        :qty="auction.qty"></bid-form>
                </div>

                <auction-bid-list
                    v-if="isBuyer"
                    class="auction-bid-list col"
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :auction-id="auctionId"></auction-bid-list>

            </div>
        </article>

        <router-link class="btn btn-block" :to="{ name: 'Dashboard' }">
            &larr; Back home
        </router-link>
    </section>
</template>

<script>
import config from '@/config'
import moment from 'moment'
import { getUserRole } from '@/api'
import { getAuctionById } from '@/api/auction'
import AuctionBidList from '@/components/AuctionBidList'
import BidForm from '@/components/BidForm'
import EpiocaCountdown from '@/components/EpiocaCountdown'
import LowBidList from '@/components/LowBidList'
import { registerFor } from '@/api/ws'

export default {
    components: {
        AuctionBidList,
        BidForm,
        EpiocaCountdown,
        LowBidList,
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
            return getUserRole() !== 'supplier'
        },

        scheduledDate() {
            return moment(this.auction.start_date).format('dddd, MMMM Do YYYY, h:mm:ss a')
        },

        endDate() {
            this.getEndDate();
            registerFor('auction_time_changed', this.getEndDate())
        },

        auctionId() {
            return this.$route.params.id
        }
    },

    methods: {
        getEndDate() {
            return moment(this.auction.start_date).add(this.auction.duration, 's').format()
        }
    },

    created() {
        getAuctionById(this.auctionId)
            .then(res => this.auction = res.data)
            .catch(err => console.error(err))
    },

}
</script>

<style scoped>
.auction-bid-list {
    margin-top: 2rem;
}
</style>
