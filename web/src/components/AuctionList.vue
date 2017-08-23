<template>
    <div class ="container">
        <div class ="row">
            <div v-for="item in items" class="col-md-4">
                <auction-list-item
                    :name="item.name"
                    :title="item.title"
                    :description="item.description"
                    :image="item.image"
                ></auction-list-item>
            </div>
        </div>
    </div>
</template>

<script>
    import AuctionListItem from '@/components/AuctionListItem'

    export default {
        components: {
            AuctionListItem
        },

        data() {
            return {
                items: [
                    {
                        'name' : 'Running',
                        'title' : 'The current running',
                        'description' : 'Lorem ipsum blah blah',
                        'image' : 'https://cdn0.iconfinder.com/data/icons/shopping-and-retail/80/Shopping_retail-06-512.png'
                    },
                    {

                        'name'  : 'Scheduled',
                        'title' : 'The current scheduled',
                        'description': 'Lorem ipsum blah blah',
                        'image' : 'https://cdn4.iconfinder.com/data/icons/ios7-line/512/Timetable.png'
                    },
                    {
                        'name' : 'Completed',
                        'title': 'The bids completed',
                        'description': 'Lorem ipsum blah blah blah',
                        'image' : 'https://cdn0.iconfinder.com/data/icons/flat-security-icons/512/tick.png'
                    }
                ]
            }
        },
    }
</script>

<style scoped></style>
