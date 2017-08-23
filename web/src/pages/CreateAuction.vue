<template>
    <div>
        Create auction

        <form @submit.prevent="saveAuction">

            <div class="form-group">
                <label for="exampleInputPassword1">Password</label>
                <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
            </div>


        </form>
    </div>
</template>

<script>
    export default {

        methods: {
            saveAuction(auction) {
                console.info(auction)
            }
        }
    }
</script>

<style>

</style>
