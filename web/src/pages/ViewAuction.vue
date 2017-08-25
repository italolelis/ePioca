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
        </div>

        <low-bid-list :auction-id="auctionId"></low-bid-list>

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
                <div v-if="auction.status === 'scheduled'">
                    Scheduled for {{ scheduledDate }}
                </div>

                <div v-else>
                    <epioca-countdown :end-date="endDate"></epioca-countdown>
                </div>
            </div>

        </div>
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
                return moment(this.auction.startDate).format('dddd, MMMM Do YYYY, h:mm:ss a')
            },

            endDate() {
                return moment(this.auction.startDate).add(this.auction.duration, 's').format()
            },

            auctionId() {
                return this.$route.params.id
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
