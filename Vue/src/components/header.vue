<template>
	<div class="container">
		<div class="logo">
			<img src="../../build/logo.png" />
		</div>
		<button class="hidden"><i class="el-icon-tickets"></i></button>
	</div>
</template>

<script>
import 'element-ui/lib/theme-chalk/display.css';
export default {
	name:'vHeader',
	components: {
	},
	data () {
		return {}
	}
}
</script>

<style lang="stylus" scoped>
	.container
		padding:0 20px
		height:70px
		background:#CCC
		display:flex
		flex-direction:row
		justify-content:space-between
		align-items:center
		.logo
			width:50px
			height:50px
			img
				width:100%
				height:100%
		@media screen and (min-width: 768px)
			.hidden
				display:none
</style>
