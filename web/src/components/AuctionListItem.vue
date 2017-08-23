<template>
    <div class="item">
        <div>
            <img :src="image" alt="blog-img">
        </div>
        <div class="item-description">
            <h4><a href="#">{{ title }}</a></h4>
            <p>{{ description }}.</p>

        </div>
        <span class="item-name">{{ name }}</span>
    </div>
</template>

<script>
export default {
    props: {
        name: {
            type: String,
            required: true
        },
        title: {
            type: String,
            required: true
        },
        description : {
            type: String,
            required: true
        },
        image : {
            type: String,
            required: true
        }
    }
}
</script>

<style scoped>
    img {
        width: 100%;
        height: auto;
    }
    .item {
        border: 1px solid #dfdede;
        box-shadow: 2px 5px 10px #dfdede;
        margin: 15px auto;
        padding: 5px;
        position: relative;
    }
    .item-description {
        padding: 15px;
    }
    .item-description h4 {
        font-size: 16px;
        font-weight: 500;
        margin-bottom: 10px;
        text-transform: uppercase;
    }
    .item-description h4 a{
        color:#777;
    }
    .item-description p{
        color: #999;
        font-size: 14px;
        font-weight: 300;
        line-height: 1.3333;
    }
    .item-name {
        position: absolute;
        background: #337ab7;
        top: 35px;
        left: 5px;
        color: #fff;
        border-radius: 0 25px 25px 0;
        padding: 5px 15px;
        font-weight: 700;
    }
</style>
