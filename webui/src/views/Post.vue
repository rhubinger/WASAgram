<script>
import store from '../services/store.js'

export default {
	data() {
		return {
			postId: null,
			comments: [],
		}
	},
	methods: {
		async postComment() {
			let comment = document.getElementById("commentInput").value;
			let pattern = /.{1,140}/;
			if(!pattern.test(comment)){
				alert("The comment must follow the pattern: " + pattern + "!");
				return;
			}

			try {
				let post = await this.$axios.post("/posts/" + this.$route.params.pid + "/comments", {
					commentId: null,
					postId: this.postId,
					userId: store.userId,
					username: null,
					datetime: null,
					comment: comment,
					}, { headers: {
					'Authorization': `Bearer ${store.identifier}`,
					},
				});
			} catch (e) {
				console.log(e.toString());
			}

			this.comments = [];
			try {
				let response = await this.$axios.get("/posts/" + this.$route.params.pid + "/comments", { headers: {
					'Authorization': `Bearer ${store.identifier}`,
					},
				});
				this.comments = response.data.comments;
			} catch (e) {
				console.log(e.toString());
			}

			document.getElementById("commentForm").reset(); 
		},
	},

	async created() {
		try {
			let response = await this.$axios.get("/posts/" + this.$route.params.pid + "/comments", { headers: {
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			this.comments = response.data.comments;
		} catch (e) {
			console.log(e.toString());
		}
    },
}
</script>

<template>
	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Post</h1>
	</div>

	<div class="postView">
		<div>
			<Post allign="center" justify="center" class="post" id="post" :pid="$route.params.pid" />
		</div>
		<div>
			<div class="container" v-for="comment in this.comments">
				<Comment v-bind:id="comment.cid" class="item" :cid="comment.commentId" :pid="comment.postId" :uid="comment.userId" :username="comment.username" 
				:datetime="comment['date-time']" :comment="comment.comment"/>
			</div>
		</div>
		<div>
			<form id="commentForm" onsubmit="return false">
				<label for="commentInput">Comment here:</label><br>
				<input type="text" id="commentInput" name="commentInput" value="Nice Picture!"><br>
				<input type="submit" value="comment" class="btn btn-sm btn-primary" @click="postComment()">
			</form>
		</div>
	</div>
</template>

<style>
.postView {
	text-align: center;
}
.post {
	text-align: center;
	max-width: 100%;
}
.container {
	padding: 10px;
}
.item {
	padding: 10px;
}
</style>