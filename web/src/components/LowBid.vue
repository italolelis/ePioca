<template>
    <b-card
        :sub-title="bidYield"
        :title="bid"
    >
        <p class="card-text" v-if="isBuyer">
            By {{ user }} <!-- at (time) -->
        </p>
    </b-card>
</template>

<script>
import { getUserRole } from '@/api'

export default {
    props: {
        user: {
            type: String,
            required: true
        },
        value: {
            type: Number,
            required: true
        },
        threshold: {
            type: Number,
            required: true
        }
    },

    computed: {
        bidYield() {
            return `${this.threshold}% yield`
        },

        isBuyer() {
            return getUserRole() !== 'supplier'
        },

        bid() {
            return '$' + parseFloat(Math.round(this.value * 100) / 100).toFixed(2);
        }
    }
}
</script>
