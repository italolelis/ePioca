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
                            @change="weekChanged"
                            :state="weekState"
                            placeholder="2017-W34"
                            :required="true"></b-form-input>
                    </b-form-group>
                </div>

                <div class="col">
                    <b-form-group label="Ingredient">
                        <b-form-select
                            v-model="form.ingredient"
                            :options="ingredients"
                            value-field="skuName"
                            text-field="productName"
                            @input="ingredientChanged"
                            :required="true">
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
                        <b-input-group left="$" right="/ unit">
                            <b-form-input v-model="form.startPrice"
                                type="number"
                                step="0.01"></b-form-input>
                        </b-input-group>
                    </b-form-group>
                </div>
            </b-form-row>

            <b-form-row>
                <div class="col">
                    <b-form-group label="Thresholds" class="thresholds">
                        <b-input-group right="%" v-for="(threshold, index) in form.threshold" :key='threshold'>
                            <b-form-input v-model="form.threshold[index]" type="text" :state="thresholdState"></b-form-input>
                            <b-btn class="btn-square" variant="danger" v-if="index !== 0" @click="removeThreshold(index)">-</b-btn>
                        </b-input-group>
                    </b-form-group>
                </div>
                <div class="col">
                    <b-form-group label="&nbsp;">
                        <b-btn class="btn" variant="info" @click="addThreshold()">Add Threshold</b-btn>
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
            ingredients: [{
                skuName: null,
                productName: 'Please select...',
                disabled: true
            }],
            form: {
                ingredient: null,
                week: '',
                startDate: moment().toISOString(),
                duration: 3600,
                qty: null,
                threshold: [0],
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
        },

        thresholdState() {
            return this.getTotalThreshold <= 100 ? null : 'invalid';
        },

        getTotalThreshold() {
            return this.form.threshold.reduce((a, b) => parseInt(a, 10) + parseInt(b, 10))
        }
    },

    methods: {
        saveAuction() {
            const ingredient = this.form.ingredient.split('|')

            const auction = {
                week: this.form.week,
                start_date: moment(this.form.startDate).toISOString(),
                duration: parseInt(this.form.duration, 10) * 60, // TODO: Make this more flexible
                ingredient: {
                    sku: ingredient[0],
                    name: ingredient[1]
                },
                qty: parseFloat(this.form.qty),
                threshold: this.form.threshold.map(t => parseInt(t, 10)),
                max_price: parseFloat(this.form.startPrice),
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

        weekChanged() {
            if (this.weekState === 'invalid') return
            this.fetchIngredients()
        },

        ingredientChanged(value) {
            const ingredient = value.split('|')
            this.form.qty = ingredient[2]
        },

        fetchIngredients() {
            getIngredientsForWeekAndDc(this.form.week, this.form.dc)
                .then(({ data }) => {
                    if (data.data.length === 0) {
                        return
                    }

                    this.ingredients = data.data.map(i => {
                        i.skuName = `${i.productSku}|${i.productName}|${i.qtyNeeded}`
                        return i
                    })

                    this.ingredient = null
                    this.ingredients.unshift({
                        skuName: null,
                        productName: 'Please select...',
                        disabled: true
                    })
                })
                .catch(err => {
                    this.alert.error.message = err.message
                    this.alert.error.show = true
                })
        },

        removeThreshold(index) {
            this.form.threshold.splice(index, 1)
        },

        addThreshold() {
             if (this.getTotalThreshold < 100) {
                this.form.threshold.push(0)
             }
        },
    }
}
</script>

<style scoped>
.btn.btn-square {
    border-radius: 0;
}

.thresholds .input-group {
    margin-bottom: 1rem;
}
</style>
