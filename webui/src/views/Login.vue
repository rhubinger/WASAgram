<script>
import axios from 'axios';
import store from '../services/store.js';

export default {
	methods: {
		async login() {
			try {
				uid = document.getElementById("uid").value;
				let response = await this.$axios.post("/login", {
					userId: uid
				});
				store.identifier = response.data.identifier;
				store.userId = uid;
				document.getElementById("loginForm").reset();
				this.$router.push('/profile/' + uid);
			} catch (e) {
				console.log(e.toString());
			}
		},
	}
}
</script>

<template>
	<div>
        <h2>WASAgram Login</h2>
		<form id="loginForm" onsubmit="return false">
			<label for="User ID">User ID:</label><br>
			<input type="text" id="uid" name="UID" value="@Alan_Turing"><br>
			<input type="submit" value="login" @click="login">
		</form>
	</div>
</template>

<style>
</style>