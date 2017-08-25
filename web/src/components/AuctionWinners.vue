<template>
    <div>
        <h3>Threshold {{ threshold }} %</h3>
        <b-list-group-item>
            <strong>{{ user }}</strong>
            <span class="pull-right">{{ bid }}</span>
        </b-list-group-item>
    </div>
</template>

<script>
export default {
    props: {
        threshold: {
            type: Number,
            required: true,
        },
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
