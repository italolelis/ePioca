<template>
    <div>
        Create auction

        <form @submit.prevent="saveAuction">

            <div class="form-row">
                <div class="col">
                    <label>Week</label>
                    <div class="input-group">
                        <input type="text" class="form-control" v-model="week">
                        <span class="input-group-btn">
                            <button class="btn btn-secondary" type="button" @click="fetchIngredients">Go</button>
                        </span>
                    </div>
                </div>

                <div class="col">
                    <label>Ingredient</label>
                    <select v-model="selectedIngredient">
                        <option
                            v-for="ingredient in ingredients"
                            :value="ingredient.productSku">
                            {{ ingredient.productName }}
                        </option>
                    </select>
                </div>
            </div>

            <div class="form-group">
                <label>Start Date</label>
                <input type="datetime" v-model="startDate">
            </div>

            <div class="form-group">
                <label>End Date</label>
                <input type="datetime" v-model="endDate">
            </div>

            <div class="form-group">
                <label>Quantity</label>
                <input type="number" v-model="qtyNeeded">
            </div>

            <div class="form-group">
                <label>Threshold</label>
                <input type="number" v-model="threshold">
            </div>

            <div class="form-group">
                <label>Start Price</label>
                <input type="number" v-model="startPrice">
            </div>

            <button type="submit" class="btn btn-primary">Save auction</button>

        </form>
    </div>
</template>

<script>
import { createAuction } from '@/api/auction'
import { getIngredientsForWeekAndDc } from '@/api/ingredient'

export default {
    data() {
        return {
            ingredients: [],
            selectedIngredient: {},
            week: '2017-W40',
            startDate: '2017-08-23T10:00:00Z',
            endDate: '2017-08-23T12:00:00Z',
            qtyNeeded: 1000,
            threshold: 40,
            startPrice: 5.00
        }
    },

    methods: {
        saveAuction() {
            const auction = {
                week: this.week,
                start_date: this.startDate,
                end_date: this.endDate,
                ingredient: this.selectedIngredient,
                qty: this.qtyNeeded,
                threshold: [ this.threshold ], // TODO
                max_price: this.startPrice,
                country: 'US',
                dc: 'TX', // TODO
                duration: 0, // TODO
            }

            createAuction(auction)
                .then(res => console.info(res))
                .catch(err => console.error(err))
        },

        fetchIngredients() {
            getIngredientsForWeekAndDc(this.week, this.dc)
                .then(res => this.ingredients = res.data)
                .catch(err => console.error(err))
        }
    }
}
</script>

<style>

</style>
