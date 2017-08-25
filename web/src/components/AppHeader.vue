<template>
    <b-jumbotron>
        <router-link :to="{ name: 'Dashboard' }">
            <img src="/static/epioca.png" class="epioca-logo">

            <img class="hf-logo pull-right" src="/static/chicken.png">
        </router-link>
    </b-jumbotron>
</template>

<script>
import { getUserRole } from '@/api'

export default {
    computed: {
        image() {
            if (getUserRole() === 'supplier') {
                return 'robert-approved.png'
            }

            return 'chicken.png'
        }
    }
}
</script>

<style scoped>
.hf-logo {
    width: 50px;
}

.epioca-logo {
    width: 150px;
}

@media screen and (min-width: 320px) {
    .hf-logo {
        width: 75px;
    }
}

@media screen and (min-width: 400px) {
    .hf-logo {
        width: 100px;
    }

    .epioca-logo {
        width: 200px;
    }
}

</style>
