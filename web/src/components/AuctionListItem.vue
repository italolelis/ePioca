<template>
    <router-link
        :to="{ name: 'View Auction', params: { id: id }}">
        <b-list-group-item>
            <strong>{{ ingredient }}</strong><span class="pull-right">&times;{{ qty }}</span>

            <p v-if="type === 'running'">{{ timeRemaining }} remaining</p>
            <p v-if="type === 'scheduled'">Starts in {{ timeToStart }}</p>
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
            const end = moment(this.startTime).add(this.duration, 's')

            return moment.duration(now.diff(end)).humanize()
        },

        timeToStart() {
            const now = moment()
            const start = moment(this.startTime)

            return moment.duration(now.diff(start)).humanize()
        }
    }
}
</script>

<style scoped>
p {
    margin: 1rem 0 0 0;
}
</style>
