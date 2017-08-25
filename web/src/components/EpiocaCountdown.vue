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
                interval: 1000,
                timeRemaining : '',
                timeInterval: '',
            }
        },

        mounted() {
            this.setCountdown()
        },

        methods: {
            getTimeRemaining() {
                return moment().countdown(this.endDate).toString();
            },

            getTimeDifference() {
                return moment(this.endDate).diff(moment(), 'seconds');
            },

            timeIsUp() {
                return '' === this.getTimeRemaining();
            },

            timeHasPassed() {
                return 0 > this.getTimeDifference()
            },

            clearCurrentInterval() {
                this.setTimeRemaining('Finished');
                clearInterval(this.timeInterval);
            },

            getTimeInterval() {
                return setInterval(() => {
                    this.timeIsUp() || this.timeHasPassed()
                        ? this.clearCurrentInterval()
                        : this.setTimeRemaining(`${this.getTimeRemaining()} remaining`);
                }, this.interval);
            },

            setTimeRemaining(value) {
                this.timeRemaining = value;
            },

            setCountdown() {
                this.timeInterval = this.getTimeInterval();
            }
        }
    }
</script>

<style scoped></style>
