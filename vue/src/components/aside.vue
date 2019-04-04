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
			
			this.axios.patch('http://localhost:8000').then((res)=>{
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
		border-left:1px solid #00a0e9
		margin-left:20px
		padding-left:20px !important
		padding-top 8px
		padding-bottom:8px
	.aside-container >>> .el-tree-node>.el-tree-node__children
		padding-left:10px
	.aside-container >>> .el-tree-node__expand-icon
		position : absolute
		left:75%
	// .container >>> .el-tree-node:focus > .el-tree-node__content
	// 	background:#c7ceb2
	.aside-container >>> .el-tree-node:focus > .el-tree-node__content
		border-left:3px solid #00a0e9
		color :#00a0e9
		font-weight:800
	.aside-container >>>  .el-tree-node__content:hover
		background:rgba(0,160,233,.1)
		color:#00a0e9
		font-weight :800
	.aside-container
		display:flex
		flex-direction:column
		overflow :auto
		width:25%
		min-width:200px
		height:100%
		//background-color:#ccc
		//max-width:240px
		border-right:1px solid #ccc
</style>
