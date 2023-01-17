<script>
import axios from 'axios';
import store from '../services/store.js';
import dateTime from '../services/datetime.js';

export default {
	data() {
		return {
			likeCount: 0,
			likes: [],
		}
	},

	methods: {
	},

	async created() {
		try {
			let response = await this.$axios.get("/posts/" + this.$route.params.pid + "/likes", { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.likeCount = response.data.length;
			this.likes = response.data.users;
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
			<h1 v-if="this.likeCount != 1" class="h2">{{this.likeCount}} Likes</h1>
			<h1 v-else class="h2">1 like</h1>
		</div>
		<div>
			<div class="container" v-for="user in likes">
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