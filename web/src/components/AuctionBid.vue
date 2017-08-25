<template>
    <b-list-group-item>
        <strong>{{ user }}</strong>
        <span class="pull-right">{{ bid }}</span>
    </b-list-group-item>
</template>

<script>
export default {
    props: {
        user: {
            type: String,
            required: true,
        },
        value: {
            type: Number,
            required: true
        }
    },

    computed: {
        bid() {
            return '$' + parseFloat(Math.round(this.value * 100) / 100).toFixed(2);
        }
    }
}
</script>
