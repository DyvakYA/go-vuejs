<template>
    <div id="home">
        <navbar/>
        <h1>I am home</h1>
    </div>
</template>

<script>
    import navbar from './Nav';

    export default {
        name: 'Home',
        components: {
            navbar
        }
    }
</script>
