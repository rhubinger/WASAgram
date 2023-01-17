<script>
import axios from 'axios';
import store from '../services/store.js';
import dateTime from '../services/datetime.js';

export default {
	data() {
		return {
			searchResult: []
		}
	},

	methods: {
		async search(){
			try {
				let searchstring = document.getElementById("searchString").value;
				let response = await this.$axios.get("/search?searchString=" + searchstring + "&uid=" + store.userId, { headers: {
					'Authorization': `Bearer ${store.identifier}` ,
					},
				});
				this.searchResult = response.data;
			} catch (e) {
				console.log(e.toString());
			}
		}
	},
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search</h1>
		</div>
        <form id="loginForm" onsubmit="return false">
			<label for="Search">Search:</label>
			<input type="text" id="searchString" name="searchstring" value="Alan">
			<input type="submit" value="search" @click="search()">
		</form>
		<div>
			<div class="container" v-for="user in searchResult">
				<User v-bind:id="user.userId" class="item" :uid="user.userId" :username="user.name" :posts="user.posts" :followers="user.followers" :followed="user.followed" />
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