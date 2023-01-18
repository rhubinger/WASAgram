<script>
import axios from 'axios';
import store from '../services/store.js';

export default {
	data() {
		return {
            user: null,
		}
	},

	methods: {
		async changeUsername() {
            let newUsername = document.getElementById("newName").value;
			let pattern = /[a-zA-z0-9-. ]{1,30}/;
			if(!pattern.test(newUsername)){
				alert("The username must follow the pattern: " + pattern + "!");
				return;
			}
			try {
				let post = await this.$axios.put("/users/" + store.userId + "/username", {
                    name: newUsername,
                }, { headers: {
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			document.getElementById("changeNameForm").reset(); 
			} catch (e) {
				console.log(e.toString());
			}
			this.$router.push('/profile/' + store.userId);
		}
	},

    async created(){
        try {
			let response = await this.$axios.get("/users/" + this.$route.params.uid, { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.user = response.data;
		} catch (e) {
			console.log(e.toString());
		}
    }
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Change Username</h1>
		</div>
	</div>

    <div class="info">
        Current Name: <br>
		<b>{{ this.user.name }}</b> <br>
    </div>
	<form id="changeNameForm" onsubmit="return false">
		<label for="newName">New Username:</label><br>
        <input type="text" id="newName" name="newName" value="new user"><br>
        <input type="submit" class="btn btn-sm btn-primary" value="Post" @click="changeUsername">
    </form>
</template>

<style>
.info{
	padding-bottom: 10px;
}
</style>