<script>
import axios from 'axios';
import store from '../services/store.js';

export default {
	data() {
		return {
		}
	},

	methods: {
		async post() {
			var formdata = new FormData();
			var picture = document.getElementById("picture").files[0];
			if(picture == null){
				alert("You must upload a picture!");
				return;
			}
			let name = document.getElementById("picture").value;
			let extension = name.substr(name.lastIndexOf(".") + 1, name.length).toLowerCase();
			if(extension != "png"){
				alert("All uploaded pictures must be PNG's!");
				return;
			}
			formdata.append("image", picture);

			var caption = document.getElementById("caption").value;
			let pattern = /.{1,140}/;
			if(!pattern.test(caption)){
				alert("All captions must follow the pattern " + pattern + "!");
				return;
			}
			var metadata = {
					postId: null,
					userId: store.userId,
					username: null,
					datetime: null,
					caption: caption,
					pictureId: null,
					likes: 0,
					comments: 0,
				};
			formdata.append("post", JSON.stringify(metadata));
			try {
				let post = await this.$axios.post("/posts", formdata, { headers: {
				'Content-Type': `multipart/form-data`,
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			document.getElementById("postForm").reset(); 
			} catch (e) {
				console.log(e.toString());
			}
			this.$router.push('/profile/' + store.userId);
		}
	},
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Create Post</h1>
		</div>
	</div>

	<form id="postForm" onsubmit="return false">
		<label for="picture">Select a picture:</label>
		<input type="file" id="picture" name="picture" accept="image/png" /><br>
		<label for="caption">Caption:</label>
        <input type="text" id="caption" name="caption" size="150" value="Nice day at the beach."><br>
		<input type="submit" class="btn btn-sm btn-primary" value="Post" @click="post">
    </form>
</template>

<style>
</style>