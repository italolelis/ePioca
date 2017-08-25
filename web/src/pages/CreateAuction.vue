<template>
    <section>
        <!-- TODO: Refactor these into alerts component -->
        <div class="row">
            <div class="col">
                <b-alert variant="danger"
                    dismissible
                    :show="alert.error.show"
                    @dismissed="alert.error.show=false">
                    {{ alert.error.message }}
                </b-alert>
            </div>
        </div>

        <b-form @submit.prevent="saveAuction">

            <b-form-row>
                <div class="col">
                    <b-form-group label="Week">
                        <b-form-input
                            type="text"
                            v-model="form.week"
                            @change="fetchIngredients"
                            :state="weekState"
                            required></b-form-input>
                    </b-form-group>
                </div>

                <div class="col">
                    <b-form-group label="Ingredient">
                        <b-form-select
                            v-model="form.ingredient"
                            :options="ingredients"
                            value-field="skuName"
                            text-field="productName"
                            required>
                        </b-form-select>
                    </b-form-group>
                </div>
            </b-form-row>

            <b-form-row>

                <div class="col">
                    <b-form-group label="Start Date">
                        <flat-pickr
                            v-model="form.startDate"
                            :config="flatpickr.config"></flat-pickr>
                    </b-form-group>
                </div>

                <div class="col">
                    <b-form-group label="Duration">
                        <b-form-input v-model="form.duration"
                            type="number"></b-form-input>
                        <b-form-text>In minutes</b-form-text>
                    </b-form-group>
                </div>
            </b-form-row>

            <b-form-row>
                <div class="col">
                    <b-form-group label="Quantity">
                        <b-form-input v-model="form.qty"
                            type="number"></b-form-input>
                    </b-form-group>
                </div>

                <div class="col">
                    <b-form-group label="Start price">
                        <b-input-group left="$">
                            <b-form-input v-model="form.startPrice"
                                type="number"
                                step="0.01"></b-form-input>
                        </b-input-group>
                    </b-form-group>
                </div>
            </b-form-row>

            <b-form-row>
                <div class="col">
                    <b-form-group label="Thresholds">
                        <b-input-group right="%">
                            <b-form-input v-model="form.threshold"
                            type="text"></b-form-input>
                        </b-input-group>
                        <b-form-text>Comma-separated!</b-form-text>
                    </b-form-group>
                </div>
            </b-form-row>

            <button type="submit" class="btn btn-primary">Save auction</button>

        </b-form>
    </section>
</template>

<script>
import moment from 'moment'
import { createAuction } from '@/api/auction'
import { getIngredientsForWeekAndDc } from '@/api/ingredient'

export default {
    data() {
        return {
            ingredients: [],
            form: {
                ingredient: {},
                week: '2017-W40',
                startDate: moment().toISOString(),
                duration: 3600,
                qty: 1000,
                threshold: "40,20",
                startPrice: 5.00,
                dc: 'TX'  // TODO: Don't hardcode
            },
            flatpickr: {
                config: {
                    enableTime: true,
                }
            },
            alert: {
                error: {
                    show: false,
                    message: ''
                }
            }
        }
    },

    computed: {
        weekState() {
            return this.form.week.match(/\d{4}-W\d{2}/) ? null : 'invalid'
        }
    },

    mounted() {
        this.fetchIngredients()
    },

    methods: {
        saveAuction() {
            const ingredient = this.form.ingredient.split('|')

            const auction = {
                week: this.form.week,
                start_date: moment(this.form.startDate).toISOString(),
                duration: this.form.duration * 60, // TODO: Make this more flexible
                ingredient: {
                    sku: ingredient[0],
                    name: ingredient[1]
                },
                qty: this.form.qty,
                threshold: this.form.threshold.split(',').map(Number), // TODO: Fix this hack
                max_price: this.form.startPrice,
                country: 'US',
                dc: this.form.dc,
            }

            createAuction(auction)
                .then(res => this.$router.push({ path: res.headers.location }))
                .catch(err => {
                    this.alert.error.message = err.message
                    this.alert.error.show = true
                })
        },

        fetchIngredients() {
            getIngredientsForWeekAndDc(this.form.week, this.form.dc)
                .then(res => {
                    this.ingredients = res.data.data.map(i => {
                        i.skuName = `${i.productSku}|${i.productName}`
                        return i
                    })
                })
                .catch(err => {
                    this.alert.error.message = err.message
                    this.alert.error.show = true
                })
        },
    }
}
</script>

<style>

</style>
