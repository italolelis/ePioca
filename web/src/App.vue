<template>
    <div id="app">
        <app-header></app-header>

        <main class="container">
            <router-view></router-view>
        </main>
    </div>
</template>

<script>
import AppHeader from '@/components/AppHeader'

export default {
    components: {
        AppHeader
    },

    name: 'ePioca'
}
</script>

<style>
</style>
