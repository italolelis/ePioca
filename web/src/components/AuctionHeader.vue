<template>
    <header>
        <h2>
            {{ ingredient }}
            <b-badge>{{ status }}</b-badge>
        </h2>

        <p class="lead">
            Auction to supply {{ qty }} units of {{ sku }} in {{ week }}
        </p>
    </header>
</template>

<script>
export default {
    props: {
        ingredient: {
            type: String,
            required: true
        },
        sku: {
            type: String,
            required: true
        },
        qty: {
            type: Number,
            required: true
        },
        status: {
            type: String,
            required: true
        },
        week: {
            type: String,
            required: true
        }
    }
}
</script>
