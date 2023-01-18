<script>
import axios from 'axios';
import store from '../services/store.js';

export default {
	data() {
		return {
			loggedIn: false,
		}
	},

	methods: {
		async login() {
			uid = document.getElementById("uid").value;
			let pattern = /@[a-zA-z0-9_.]{3,16}/;
			if(!pattern.test(uid)){
				alert("The userId must follow the pattern: " + pattern + "!");
				return;
			}
			try {
				let response = await this.$axios.post("/login", {
					userId: uid
				});
				store.identifier = response.data.identifier;
				store.userId = uid;
				this.loggedIn = true;
				document.getElementById("loginForm").reset();
				this.$router.push('/stream/' + uid);
			} catch (e) {
				console.log(e.toString());
			}
		},

		goToLogin(){
			this.loggedIn = false;
		}
	}
}
</script>

<template>
	<div class="login">
        <h2>WASAgram Login</h2>
		<form id="loginForm" onsubmit="return false">
			<label for="User ID">Please enter your user ID:</label><br>
			<input type="text" id="uid" name="UID" value="@Alan_Turing"><br>
			<input type="submit" class="btn btn-sm btn-primary" value="login" @click="login">
		</form>
	</div>
</template>

<style>
.login{
	padding-top: 100px;
	text-align: center;
}
</style>