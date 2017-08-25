<template>
    <div>
        <div>
            <h4> Ingredient of this auction: {{ auction.ingredient.name }} </h4>
        </div>
        <div v-if="'buyer' === userRole">
            <h2>Suppliers for this Auction</h2>
            <h3> {{ auctionDate }} </h3>
            <b-list-group>
                <auction-winners
                    v-for="bid in bidings"
                    :key="bid.id"
                    :threshold="bid.threshold"
                    :user="bid.user.name"
                    :value="bid.value"></auction-winners>
            </b-list-group>
        </div>
        <div v-if="'supplier' === userRole">
            <h3> {{ auctionDate }} </h3>
            <div v-for="bid in userBids">
                {{ bid.message }}
            </div>
        </div>
    </div>
</template>

<script>
    import moment from 'moment'
    import { getUserId } from '@/api/'
    import { getUserName } from '@/api/'
    import { getUserRole } from '@/api/'
    import { getAuctionById } from '@/api/auction'
    import { getWinners } from '@/api/auction'
    import AuctionWinners from '@/components/AuctionWinners'

    export default {
        data() {
            return {
                bidings: [],
                auction: {},
                userBids: [],
                userRole: getUserRole(),
                userName: getUserName(),
                userId: getUserId(),
                auctionDate: ''
            }
        },

        components : {
            AuctionWinners
        },

        created() {
            getWinners(this.$route.params.id).then( res => {
                this.bidings  = res.data;
                this.userBids = res.data.map( (value) => {
                    console.log(value);
                    if (value.user.id === this.userId) {
                        return {
                            'message' : `You won this threshold ${value.threshold} %, your bid was: ${value.value.toFixed(2)} $`,
                        }
                    }
                    return {
                        'message' : `You did not win this threshold ${value.threshold} %, the lowest bid was ${value.value.toFixed(2)} $`,
                    }
                });
            });
            getAuctionById(this.$route.params.id).then( res => {
                this.auction = res.data;
                this.auctionDate = moment(this.auction.start_date).format('LLL');
            });


        }
    }
</script>

<style scoped></style>
