<template>
    <div class="container" id="app">
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        name: 'app'
    }
</script>

<style>
    body {
        background: #1C2135;
        color: #CACACC;
        font-family: 'Roboto Mono', monospace;
        font-size: 1em;
        justify-content: center;
        align-content: space-between;
    }

    .container {
        margin: 0 auto;
        width: 80%;
    }

    h1 {
        font-weight: 400;
        margin: 0;
    }
</style>
