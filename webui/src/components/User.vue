<script>
export default {
    props: ['uid', 'username', 'posts', 'followed', 'followers'],

	data: function() {
		return {
		}
	},
	methods: {

	}

}
</script>

<template>
	<div>
        <div>
            {{ this.username }} - {{ this.uid }}
        </div>
        <div>
            Posts: {{ this.posts }} | Followers: {{ this.followed }} | Followed: {{ this.followers }}
        </div>
    </div>
</template>

<style></style>