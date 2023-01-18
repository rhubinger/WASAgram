<script>
import store from '../services/store';

export default {
    props: ['cid', 'pid', 'uid', 'username', 'datetime', 'comment'],

	data: function() {
		return {
            deleted: false,
            allowDelete: false,
		}
	},

    methods: {
        async deleteComment() {
            if(!confirm("Are you sure that you want to delete this comment?")){
                return
            }
			try {
				let response = await this.$axios.delete("/posts/" + this.pid + "/comments/" + this.cid, { headers: {
					'Authorization': `Bearer ${store.identifier}` ,
					},
				});
			} catch (e) {
				console.log(e.toString());
			}
            this.deleted = true;
        },
    },
    
    async created (){
        this.allowDelete = this.uid == store.userId;
    },
}
</script>

<template>
    <div class="comment">
        <div v-if="!this.deleted">
            <div>
                <div> {{ this.username }} </div>
                <div> {{ this.uid }} </div>
            </div>
            <div>
                <div> {{ this.comment }}</div>
                <div> {{ this.datetime }}</div>
            </div>
            <div>
                <button v-if="allowDelete" type="button" v-on:click.stop="deleteComment()">
                    delete
                </button> 
            </div>
        </div>
        <div v-else>
            Comment deleted
        </div>
    </div>
</template>

<style>
.comment{
    background-color: lightgrey;
    border-radius: 8px;
}
</style>