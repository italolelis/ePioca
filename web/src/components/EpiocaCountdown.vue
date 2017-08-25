<template>
    <div>
        {{ timeRemaining }}
    </div>
</template>

<script>
    import moment from 'moment'
    import countdown from 'countdown'
    import moment_countdown from 'moment-countdown'

    export default {
        props: {
            endDate : {
                required: true
            }
        },

        data() {
            return {
                timeRemaining : '',
                timeInterval: ''
            }
        },

        mounted() {
            this.getCountDown()
        },

        methods: {
            getCountDown() {
                this.timeInterval = setInterval(function(){
                    let difference    = moment(this.endDate).diff(moment(), 'seconds');
                    let timeRemaining = moment().countdown(this.endDate).toString();
                    if ('' === timeRemaining || difference < 0) {
                        this.timeRemaining = 'Finished';
                        clearInterval(this.timeInterval);
                        return;
                    }
                    this.timeRemaining = `${timeRemaining} remaining`;
                }.bind(this), this.interval);
            }
        }
    }
</script>

<style scoped></style>
