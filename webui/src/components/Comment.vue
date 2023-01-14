<script>
export default {
    props: ['cid', 'pid', 'uid', 'username', 'datetime', 'comment'],

	data: function() {
		return {
		}
	},
}
</script>

<template>
	<div>
        <div>
            <div> {{ this.username }} </div>
            <div> {{ this.uid }} </div>
        </div>
        <div>
            <div> {{ this.comment }}</div>
            <div> {{ this.datetime }}</div>
        </div>
    </div>
</template>

<style></style>