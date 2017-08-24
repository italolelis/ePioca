<template>
    <router-link
        :to="{ name: 'View Auction', params: { id: id }}">
        <b-list-group-item>
            <span>{{ ingredient }}</span>
            <span>{{ qty }}</span>
            <p v-if="type == 'running'">{{ timeRemaining }} remaining</p>
        </b-list-group-item>
    </router-link>
</template>

<script>
import moment from 'moment'

export default {
    props: {
        id: {
            type: String,
            required: true
        },
        ingredient: {
            type: String,
            required: true
        },
        qty: {
            type: Number,
            required: true
        },
        type: {
            type: String
        },
        startTime: {
            type: String,
            required: true
        },
        duration: {
            type: Number,
            required: true
        }
    },

    computed: {
        timeRemaining() {
            const now = moment()
            const end = moment(this.startTime).add(this.duration)

            return moment.duration(now.diff(end)).humanize()
        }
    }
}
</script>

<style scoped></style>
