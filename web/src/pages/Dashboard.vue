<template>
    <section>
        <div class="row">
            <div class="col-md-4" v-for="type in types">
                <p class="text-center">
                    <img :src="type.image">
                </p>

                <h3 class="text-center">{{ type.filter | capitalize }} Auctions</h3>

                <auction-list :filter="type.filter"></auction-list>
            </div>
        </div>
    </section>
</template>

<script>
    import AuctionList from '@/components/AuctionList'

    export default {
        components: {
            AuctionList
        },

        filters: {
            capitalize: function (value) {
                if (!value) return ''
                value = value.toString()
                return value.charAt(0).toUpperCase() + value.slice(1)
            }
        },

        data() {
            return {
                types : [
                    {
                        'filter' : 'running',
                        'image': 'https://cdn0.iconfinder.com/data/icons/shopping-and-retail/80/Shopping_retail-06-512.png'
                    },
                    {
                        'filter' : 'scheduled',
                        'image': 'https://cdn4.iconfinder.com/data/icons/ios7-line/512/Timetable.png'
                    },
                    {
                        'filter' : 'completed',
                        'image': 'https://cdn0.iconfinder.com/data/icons/flat-security-icons/512/tick.png'
                    }
                ]
            }
        }
    }
</script>

<style scoped>
img {
    width: 50%;
    margin: 2rem auto 1rem auto;
}
</style>
