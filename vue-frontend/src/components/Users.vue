<template>
    <div id="users">
        <navbar></navbar>
        <h1>Users</h1>
        <div>
            {{ users }}
        </div>
        <ul>
            <li v-for="user in users" v-bind:key="user.id">{{ user }}</li>
        </ul>
    </div>
</template>

<script>
    import navbar from './Nav';
    import {AXIOS} from './js/http-common'

    export default {
        name: 'Users',
        data: () => ({
            users: []
        }),
        components: {
            navbar
        },
        methods: {
            getUsers: function () {
                AXIOS.get("http://localhost:8080/users")
                    .then(response => {
                        this.users = response.data
                        console.log(this.users);
                    }).catch(function (error) {
                    console.log(error);
                });
            }
        },
        beforeMount() {
            this.getUsers()
        },
    }
</script>
