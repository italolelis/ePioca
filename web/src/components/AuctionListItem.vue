<template>
    <div class="item-description">
        <span v-if="type == 'running'">{{ ingredient }}</span>
        <span v-if="type == 'running'">{{ qty }}</span>
        <span v-if="type == 'running'">01:12</span>
    </div>
</template>

<script>
    export default {
        props: {
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
            }
        }
    }
</script>

<style scoped></style>
