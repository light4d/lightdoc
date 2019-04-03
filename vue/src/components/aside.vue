<template>
	<div class="aside-container">
		<v-header></v-header>
		<v-search></v-search>
		<div class="tree">
			<!-- <v-tree></v-tree> -->
			<el-tree  :data="data" :props="defaultProps" @node-click="handleNodeClick">
			</el-tree>
		</div>
	</div>
</template>

<script>
import vSearch from './search.vue'
import vHeader from './header.vue'
import marked from 'marked'
export default {
	 data() {
      return {
        data: [],
        defaultProps: {
          children: 'Nodes',
          label: 'Name'
				},
				readmeContent:{}
      };
    },
		components:{
			vSearch,
			vHeader
		},
		mounted(){
			
			this.axios.patch('/api').then((res)=>{
				this.data=res.data.Nodes
			})
		},
    methods: {
      handleNodeClick(data) {
			
				 //this.axios.get('/api/'+data.Path)
				//this.$emit('updataPath',data.Path)
				this.$store.commit('changeValue',data)
      }
    }
}
</script>

<style lang="stylus" scoped>
	.aside-container >>> .el-tree-node__content
		border-left:2px solid #00a0e9
		margin-left:20px
		padding-left:20px !important
	.aside-container >>> .el-tree-node>.el-tree-node__children
		padding-left:10px
	.aside-container >>> .el-tree-node__expand-icon
		position : absolute
		left:75%
	// .container >>> .el-tree-node:focus > .el-tree-node__content
	// 	background:#c7ceb2
	.aside-container >>> .el-tree-node:focus > .el-tree-node__content
		border-left:6px solid #00a0e9
	.aside-container
		display:flex
		flex-direction:column
		overflow :auto
		width:30%
		min-width:200px
		height:105%
		//background-color:#ccc
		max-width:224px
		border-right:1px solid #ccc
</style>
