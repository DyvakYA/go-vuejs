<template>
    <div>
        <nav>
            <ul id="nav">
                <li>
                    <div>
                        <router-link to="/">Home</router-link>
                    </div>
                </li>
                <li>
                    <div>
                        <router-link to="/users">Users</router-link>
                    </div>
                </li>
                <li>
                    <div>
                        <router-link to="/">Products</router-link>
                    </div>
                </li>
                <li>
                    <div>
                        <router-link to="/">Orders</router-link>
                    </div>
                </li>
            </ul>
        </nav>
        <hr>

    </div>
</template>

<script>
    export default {
        name: "nav-menu"
    }
</script>

<style scoped>
    nav ul {
        list-style-type: none;
        margin: 0;
        padding: 0;
        overflow: hidden;

    }

    nav ul li {
        float: left;
    }

    nav ul li div {
        display: block;
        color: #E9EBEF;
        text-align: center;
        padding: 14px 16px;
    }

    nav ul li div:hover {
        background-color: #111;
    }

    nav ul li div a{
        color: #E9EBEF;
    }
</style>