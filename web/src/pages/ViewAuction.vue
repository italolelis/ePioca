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
                <bid-form
                    class="col"
                    v-if="!isBuyer"
                    v-show="auction.status === 'running'"
                    v-for="threshold in auction.threshold"
                    :key="threshold"
                    :threshold="threshold"
                    :qty="auction.qty"></bid-form>

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
import { registerFor } from '@/api/ws'
import AuctionBidList from '@/components/AuctionBidList'
import AuctionHeader from '@/components/AuctionHeader'
import BidForm from '@/components/BidForm'
import EpiocaCountdown from '@/components/EpiocaCountdown'
import LowBidList from '@/components/LowBidList'

export default {
    components: {
        AuctionBidList,
        AuctionHeader,
        BidForm,
        EpiocaCountdown,
        LowBidList,
    },

    data() {
        return {
            auction: {
                ingredient: {
                    name: '',
                    sku: ''
                },
                week: '',
                qty: 0
            },
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
            return moment(this.auction.start_date).add(this.auction.duration, 's').format()
        },

        auctionId() {
            return this.$route.params.id
        }
    },

    methods: {
        timeChanged({ data }) {
            this.auction.duration = data.duration
        }
    },

    created() {
        registerFor('auction_time_changed', this.timeChanged)

        getAuctionById(this.auctionId)
            .then(({ data }) => {
                if (data.status === 'completed') {
                    this.$router.push({ name: 'Winners', params: { id: this.auctionId }})
                }

                this.auction = data
            })
            .catch(err => console.error(err))
    },

}
</script>

<style scoped>
.auction-bid-list {
    margin-top: 2rem;
}
</style>
