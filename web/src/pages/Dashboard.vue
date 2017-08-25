<template>
    <section>

        <div class="row" v-if="userRole === 'buyer'">
            <div class="col">
                <router-link :to="{ name: 'Create Auction' }" class="btn btn-success btn-lg btn-block">
                    Create Auction
                </router-link>
            </div>
        </div>

        <b-card-group deck>
            <auction-list :filter="type" :key="type" v-for="type in types"></auction-list>
        </b-card-group>
    </section>
</template>

<script>
import AuctionList from '@/components/AuctionList'
import { getUserRole } from '@/api'

export default {
    components: {
        AuctionList
    },

    computed: {
        userRole() {
            return getUserRole()
        }
    },

    data() {
        return {
            types : [ 'scheduled', 'running', 'completed' ]
        }
    }
}
</script>

<style scoped>
.card-deck {
    margin-top: 2rem;
}
</style>
