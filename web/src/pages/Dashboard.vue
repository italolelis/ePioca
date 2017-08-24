<template>
    <div class ="container">
        <div class ="row">
            <div v-for="type in types" class="col-md-4">
                <div class="item">
                    <div><img :src="type.image"></div>
                    <div class="item-description">
                        <auction-list v-bind:filter="type.filter"></auction-list>
                    </div>
                    <span class="item-name">{{ type.filter }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import AuctionList from '@/components/AuctionList'

    export default {
        components: {
            AuctionList
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
                        'filter' : 'complete',
                        'image': 'https://cdn0.iconfinder.com/data/icons/flat-security-icons/512/tick.png'
                    }
                ]
            }
        }
    }
</script>

<style scoped>
    img {
        width: 100%;
        height: auto;
    }
    .item {
        border: 1px solid #dfdede;
        box-shadow: 2px 5px 10px #dfdede;
        margin: 15px auto;
        padding: 5px;
        position: relative;
    }
    .item-description {
        padding: 15px;
    }
    .item-description h4 {
        font-size: 16px;
        font-weight: 500;
        margin-bottom: 10px;
        text-transform: uppercase;
    }
    .item-description h4 a{
        color:#777;
    }
    .item-description p{
        color: #999;
        font-size: 14px;
        font-weight: 300;
        line-height: 1.3333;
    }
    .item-name {
        position: absolute;
        background: #337ab7;
        top: 35px;
        left: 5px;
        color: #fff;
        border-radius: 0 25px 25px 0;
        padding: 5px 15px;
        font-weight: 700;
    }
</style>
