<template>
    <b-form @submit.prevent="placeBid">
        <h5>
            Bidding for contract to supply {{ thresholdQty }} units
            <small>({{ threshold }}% yield)</small>
        </h5>

        <!-- TODO: Extract into alerts component -->
        <div>
            <b-alert :show="response.successCountDown"
                dismissible
                variant="success"
                @dismissed="response.successCountdown=0"
                @dismiss-count-down="successCountDownChanged">
                    Bid placed successfully!
                <b-progress variant="success"
                    :max="response.dismissSecs"
                    :value="response.successCountDown"
                    height="4px">
                </b-progress>
            </b-alert>

            <b-alert :show="response.errorCountDown"
                dismissible
                variant="danger"
                @dismissed="response.errorCountdown=0"
                @dismiss-count-down="errorCountDownChanged">
                    {{ response.error }}
                <b-progress variant="danger"
                    :max="response.dismissSecs"
                    :value="response.errorCountDown"
                    height="4px">
                </b-progress>
            </b-alert>
        </div>

        <b-form-group>
            <b-input-group left="$">
                <b-form-input v-model="form.bid"
                    type="number"
                    step="0.01"></b-form-input>

                <b-input-group-addon>/ unit</b-input-group-addon>

                <b-input-group-button>
                    <b-btn type="submit" variant="success">Place Bid</b-btn>
                </b-input-group-button>
            </b-input-group>
        </b-form-group>
    </b-form>
</template>

<script>
import { postBid } from '@/api/auction'

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
            },
            response: {
                dismissSecs: 4,
                errorCountDown: 0,
                successCountDown: 0,
                error: {}
            }
        }
    },

    computed: {
        thresholdQty() {
            return parseInt(this.qty / 100 * this.threshold, 10)
        },
    },

    methods: {
        placeBid() {
            postBid(this.$route.params.id, this.threshold, parseFloat(this.form.bid))
                .then(res => this.response.successCountDown = this.response.dismissSecs)
                .catch(err => {
                    this.response.error = err.data ? err.data : err
                    this.response.errorCountDown = this.response.dismissSecs
                })
        },

        errorCountDownChanged(errorCountDown) {
            this.response.errorCountDown = errorCountDown
        },

        successCountDownChanged(errorCountDown) {
            this.response.successCountDown = errorCountDown
        },
    }
}
</script>
