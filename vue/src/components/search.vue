<template>
		<div class="search">
			<input class="inpt" v-model="searchContent" placeholder="请输入搜索内容" />
			<i class="el-icon-search iconSearch"></i>
		</div>
</template>
<script>
	export default{
		name:"vSearch",
		data(){
			return {
				searchContent:''
			};
		},
		computed:{
			
		},
		methods:{
			
		},
		mounted(){
		
		}
	}
</script>

<style lang="stylus" scoped>
	.search
		position:relative
		.inpt
			width:100%
			text-indent:30px
			display:block
			height:30px
			border-bottom:1px solid #ccc
			border-top:1px solid #ccc
		.iconSearch
			position:absolute
			left:5%
			top:25%
</style>
