<template>
	<div class="container">
		<v-search></v-search>
		<div class="tree">
			<!-- <v-tree></v-tree> -->
			<el-tree  :data="data" :props="defaultProps" @node-click="handleNodeClick"></el-tree>
		</div>
	</div>
</template>

<script>
import vSearch from './search.vue'
import vTree from './tree.vue'
export default {
	 data() {
      return {
        data: [],
        defaultProps: {
          children: 'Nodes',
          label: 'Name'
        }
      };
    },
		components:{
			vSearch,
			vTree
		},
		mounted(){
			this.axios.patch('/api').then((res)=>{
				this.data=res.data.Nodes
				console.log(res)
			})
		},
    methods: {
      handleNodeClick(data) {
				console.log()
				 //this.axios.get('/api/'+data.Path)
				this.$emit('updataPath',data.Path)
      }
    }
}
</script>

<style lang="stylus" scoped>
	.container
		display:flex
		flex-direction:column
		width:30%
		height:100%
		//background-color:#ccc
		max-width:224px
		border-right:1px solid #ccc
</style>
