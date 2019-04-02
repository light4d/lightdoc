<template>
	<div class="container">
		<v-search></v-search>
		<div class="tree">
			<el-tree :data="data" :props="defaultProps" @node-click="handleNodeClick"></el-tree>
		</div>
	</div>
</template>

<script>
import vSearch from './search.vue'
export default {
	 data() {
      return {
        data: [],
        defaultProps: {
          children: 'Folder',
          label: 'Name'
        }
      };
    },
		components:{
			vSearch
		},
		mounted(){
			this.axios.patch('/api').then((res)=>{
				this.data=res.data.Folder
				console.log(res)
			})
		},
    methods: {
      handleNodeClick(data) {
        
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
