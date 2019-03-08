<template>
    <div id="users">
        <navbar></navbar>
        <h1>Users</h1>
        <ul>
            <div v-for="user in users" :key="user.id">{{ user.id}} {{ user.description }}</div>
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
                        console.log(this.service);
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
