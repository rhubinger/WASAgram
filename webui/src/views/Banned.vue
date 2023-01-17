<script>
import axios from 'axios';
import store from '../services/store.js';
import dateTime from '../services/datetime.js';

export default {
	data() {
		return {
			bannedCount: 0,
			banned: [],
		}
	},

	methods: {
	},

	async created() {
		try {
			let response = await this.$axios.get("/users/" + this.$route.params.uid + "/banned", { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.bannedCount = response.data.length;
			this.banned = response.data.users;
		} catch (e) {
			console.log(e.toString());
		}
	},
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{this.bannedCount}} Banned</h1>
		</div>
		<div>
			<div class="container" v-for="user in banned">
				<User class="item" :uid="user.userId" :username="user.name" :posts="user.posts" :followers="user.followers" :followed="user.followed" />
			</div>
		</div>
	</div>
</template>

<style>
	.container {
	padding: 10px;
	}
	.item {
	padding: 10px;
	}
</style>