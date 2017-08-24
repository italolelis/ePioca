<template>
  <section>
    <h1>404</h1>
    <h2>The page {{ url }} was not found...</h2>
  </section>
</template>

<script>
export default {
    data() {
        return {
            url: null,
        }
    },

    created() {
        this.url = this.$route.fullPath
    }
}
</script>
