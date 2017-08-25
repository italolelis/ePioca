<template>
    <b-form @submit.prevent="placeBid">
        <h5>Bidding for {{ thresholdQty }} units</h5>

        <b-form-group>
            <b-input-group left="$">
                <b-form-input v-model="form.bid"
                    type="number"
                    step="0.01"></b-form-input>

                <b-input-group-addon>/ unit</b-input-group-addon>

                <b-input-group-button>
                    <b-btn variant="success">Place Bid</b-btn>
                </b-input-group-button>
            </b-input-group>
        </b-form-group>
    </b-form>
</template>

<script>
export default {
    props: {
        qty: {
            type: Number,
            required: true
        },
        threshold: {
            type: Number,
            required: true
        }
    },

    data() {
        return {
            form: {
                bid: 0,
            }
        }
    },

    computed: {
        thresholdQty() {
            return this.qty / 100 * this.threshold
        },
    }
}
</script>

<style>

</style>
